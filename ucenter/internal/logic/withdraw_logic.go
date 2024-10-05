// Package logic
// @Author twilikiss 2024/8/14 0:16:16
package logic

import (
	"common/dbutils"
	"common/dbutils/tran"
	"common/op"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/withdraw"
	"strconv"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
	"ucenter/internal/model"
	"ucenter/internal/svc"
)

type WithdrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberTAddressDomain *domain.MemberTAddressDomain
	memberDomain         *domain.MemberDomain
	memberWalletDomain   *domain.MemberWalletDomain
	transaction          tran.Transaction
	withdrawDomain       *domain.WithdrawDomain
}

func (l *WithdrawLogic) FindAddressByCoinId(req *withdraw.WithdrawReq) (*withdraw.AddressSimpleList, error) {
	list, err := l.memberTAddressDomain.FindAddressList(l.ctx, req.UserId, req.CoinId)
	if err != nil {
		return nil, err
	}
	var addressList []*withdraw.AddressSimple
	err = copier.Copy(&addressList, list)
	if err != nil {
		logx.Error("copier.Copy is error, err=", err)
	}
	return &withdraw.AddressSimpleList{
		List: addressList,
	}, nil
}
func (l *WithdrawLogic) SendCode(req *withdraw.WithdrawReq) (*withdraw.NoRes, error) {
	//假设发送了一条短信 验证码是123456
	code := 123456
	//* 将验证码存入redis，过期时间5分钟
	// 设置context，设置存入redis的超时时间为5s
	ctx2, cancelFunc2 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc2()
	err := l.svcCtx.Cache.SetexCtx(ctx2, "WITHDRAW::"+req.Phone, strconv.Itoa(code), 5*60)
	return &withdraw.NoRes{}, err
}

func (l *WithdrawLogic) WithdrawCode(req *withdraw.WithdrawReq) (*withdraw.NoRes, error) {
	// 1.校验验证码
	member, err := l.memberDomain.FindMemberById(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var redisCode string
	redisCode, err = l.svcCtx.Cache.GetCtx(l.ctx, "WITHDRAW::"+member.MobilePhone)
	if err != nil {
		return nil, err
	}
	if req.Code != redisCode {
		logx.Error("验证码错误")
		return nil, errors.New("验证码不正确")
	}
	// 2.校验交易密码是否正确
	if member.JyPassword != req.JyPassword {
		logx.Error("交易密码错误")
		return nil, errors.New("交易密码不正确")
	}
	// 3.根据用户id 和 unit 查询用户的钱包 判断余额是否足够
	memberWallet, err := l.memberWalletDomain.FindWalletByMemIdAndCoin(l.ctx, req.UserId, req.Unit)
	if err != nil {
		return nil, err
	}
	if memberWallet == nil {
		logx.Error("无法找到用户对应钱包")
		return nil, errors.New("钱包不存在")
	}
	if memberWallet.Balance < req.Amount {
		logx.Error("钱包余额不足")
		return nil, errors.New("余额不足")
	}
	err = l.transaction.Action(func(conn dbutils.DbConn) error {
		// 4. 冻结用户的钱 提现币 原因：经过比特币网络 需要时间
		// 注意冻结用户金额属于敏感操作，需要保证其原子性，这里我们加入事务处理
		err2 := l.memberWalletDomain.Freeze(l.ctx, conn, req.UserId, req.Amount, req.Unit)
		if err2 != nil {
			return err2
		}
		// 5.构建&model.WithdrawRecord{}， 保存用户的提现记录
		wr := &model.WithdrawRecord{}
		wr.CoinId = memberWallet.CoinId
		wr.Address = req.Address
		wr.Fee = req.Fee
		wr.TotalAmount = req.Amount
		wr.ArrivedAmount = op.SubFloor(req.Amount, req.Fee, 10)
		wr.Remark = ""
		wr.CanAutoWithdraw = 0
		wr.IsAuto = 0
		wr.Status = 0 //审核中
		wr.CreateTime = time.Now().UnixMilli()
		wr.DealTime = 0
		wr.MemberId = req.UserId
		wr.TransactionNumber = "" //目前还没有交易编号
		var err error
		err = l.withdrawDomain.SaveRecord(l.ctx, wr)
		if err != nil {
			return err
		}
		// 6. 发送用户的提现事件到Kafka当中 MQ消费者去处理提现（交由Market模块：①创建交易 ②广播到比特币的网络）
		kafkaData, _ := json.Marshal(wr)
		data := database.KafkaData{
			Topic: "withdraw",
			Data:  kafkaData,
			Key:   []byte(fmt.Sprintf("%d", req.UserId)),
		}
		// 如果发送出现错误，先行重试3次，尽可能将我们的数据发送到kafka中， 如果三次仍未成功就触发事务机制，进行回滚操作
		for i := 0; i < 3; i++ {
			err = l.svcCtx.KafkaClient.SendSync(data)
			if err != nil {
				// 防止短时间重试连续失败
				time.Sleep(500 * time.Millisecond)
				continue
			}
			//发送成功 跳出循环
			break
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &withdraw.NoRes{}, nil
}

func (l *WithdrawLogic) WithdrawRecord(req *withdraw.WithdrawReq) (*withdraw.RecordList, error) {
	list, total, err := l.withdrawDomain.RecordList(l.ctx, req.UserId, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	var rList []*withdraw.WithdrawRecord
	copier.Copy(&rList, list)
	return &withdraw.RecordList{
		List:  rList,
		Total: total,
	}, nil
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		ctx:                  ctx,
		svcCtx:               svcCtx,
		Logger:               logx.WithContext(ctx),
		transaction:          tran.NewTransaction(svcCtx.Db.Conn),
		memberTAddressDomain: domain.NewMemberTAddressDomain(svcCtx.Db),
		memberDomain:         domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain:   domain.NewMemberWalletDomain(svcCtx.Db),
		withdrawDomain:       domain.NewWithdrawDomain(svcCtx.Db, svcCtx.MarketRPC, svcCtx.Config.Bitcoin.Address),
	}
}
