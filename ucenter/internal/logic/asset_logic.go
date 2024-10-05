// Package logic
// @Author twilikiss 2024/5/13 23:09:09
package logic

import (
	"common/bc"
	"common/op"
	"common/tools"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/types/asset"
	"ucenter/internal/domain"
	"ucenter/internal/model"
	"ucenter/internal/svc"
)

type AssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain            *domain.MemberDomain
	memberWalletDomain      *domain.MemberWalletDomain
	memberTransactionDomain *domain.MemberTransactionDomain
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                     ctx,
		svcCtx:                  svcCtx,
		Logger:                  logx.WithContext(ctx),
		memberDomain:            domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain:      domain.NewMemberWalletDomain(svcCtx.Db),
		memberTransactionDomain: domain.NewMemberTransactionDomain(svcCtx.Db),
	}
}

func (l *AssetLogic) FindWalletBySymbol(req *asset.AssetReq) (*asset.MemberWallet, error) {
	//通过market rpc 进行coin表的查询 coin信息
	//通过钱包 查询对应币的钱包信息  coin_id  user_id 查询用户的钱包信息 组装信息
	coinInfo, err := l.svcCtx.MarketRPC.FindCoinInfo(l.ctx, &mclient.MarketReq{
		Unit: req.CoinName,
	})
	if err != nil {
		return nil, err
	}
	memberWalletCoin, err := l.memberWalletDomain.FindWalletBySymbol(l.ctx, req.UserId, req.CoinName, coinInfo)
	if err != nil {
		return nil, err
	}
	resp := &asset.MemberWallet{}
	_ = copier.Copy(resp, memberWalletCoin)
	return resp, nil
}

func (l *AssetLogic) FindWallet(req *asset.AssetReq) (*asset.MemberWalletList, error) {
	memberWallets, err := l.memberWalletDomain.FindWallet(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	//查询cny的汇率
	cnyRateStr, err := l.svcCtx.Cache.Get("USDT::CNY::RATE")
	if err != nil {
		logx.Error("redis获取cny汇率数据失败,使用默认汇率，err=", err)
	}

	var temp []*model.MemberWalletCoin

	var cnyRate float64 = 7
	if cnyRateStr != "" {
		cnyRate = tools.ToFloat64(cnyRateStr)
	}
	//需要查询 币种的详情
	for _, v := range memberWallets {
		coinInfo, err := l.svcCtx.MarketRPC.FindCoinInfo(l.ctx, &mclient.MarketReq{
			Unit: v.CoinName,
		})
		if err != nil {
			return nil, err
		}
		if coinInfo.Unit == "USDT" {
			coinInfo.CnyRate = cnyRate
			coinInfo.UsdRate = 1
		} else {
			var usdtRate float64 = 20000
			usdtRateStr, err := l.svcCtx.Cache.Get(v.CoinName + "::USDT::RATE")

			if err != nil {
				logx.Error("redis获取usdt汇率数据失败,使用默认汇率，err=", err)
			}

			if usdtRateStr != "" {
				usdtRate = tools.ToFloat64(usdtRateStr)
			}
			coinInfo.UsdRate = usdtRate
			coinInfo.CnyRate = op.MulFloor(cnyRate, coinInfo.UsdRate, 10)
		}
		temp = append(temp, v.Copy(coinInfo))
	}

	var list []*asset.MemberWallet
	_ = copier.Copy(&list, temp)
	return &asset.MemberWalletList{
		List: list,
	}, nil

}

//func (l *AssetLogic) getWalletAddress() (string, error) {
//	//{
//	//    "jsonrpc": "1.0",
//	//    "method": "getmininginfo",
//	//    "params":[],
//	//    "id": "elcoin"
//	//}
//	params := make(map[string]any)
//	params["jsonrpc"] = "1.0"
//	params["method"] = "getmininginfo"
//	params["params"] = []int{}
//	params["id"] = "elcoin"
//	headers := make(map[string]string)
//	headers["Authorization"] = "Basic Yml0Y29pbjoxMjM0NTY="
//	bytes, err := tools.PostWithHeader(address, params, headers, "")
//	if err != nil {
//		return 0, err
//	}
//	var result MiningInfoResult
//	_ = json.Unmarshal(bytes, &result)
//	if result.Error != "" {
//		return 0, errors.New(result.Error)
//	}
//	return int64(result.Result.Blocks), nil
//}

func (l *AssetLogic) ResetAddress(req *asset.AssetReq) (*asset.RestAddrResp, error) {
	// 查询用户的钱包 检查address是否为空 如果未空 生成地址 进行更新
	memberWallet, err := l.memberWalletDomain.FindWalletByMemIdAndCoin(l.ctx, req.UserId, req.CoinName)
	if err != nil {
		return nil, err
	}
	if req.CoinName == "BTC" {
		if memberWallet.Address == "" {
			wallet, err := bc.NewWallet()
			if err != nil {
				return nil, err
			}
			address := wallet.GetTestAddress()
			priKey := wallet.GetPriKey()
			memberWallet.AddressPrivateKey = priKey
			memberWallet.Address = string(address)
			err = l.memberWalletDomain.UpdateAddress(l.ctx, memberWallet)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if memberWallet.Address == "" {
			address := "test address"
			priKey := "test private key"
			memberWallet.AddressPrivateKey = priKey
			memberWallet.Address = address
			err = l.memberWalletDomain.UpdateAddress(l.ctx, memberWallet)
			if err != nil {
				return nil, err
			}
		}
	}
	return &asset.RestAddrResp{}, err
}

func (l *AssetLogic) FindAllTransaction(req *asset.AssetReq) (*asset.MemberTransactionList, error) {
	// 查询所有的充值记录
	//查询所有的充值记录 分页查询
	memberTransactionVos, total, err := l.memberTransactionDomain.FindTransaction(
		l.ctx,
		req.UserId,
		req.PageNo,
		req.PageSize,
		req.Symbol,
		req.StartTime,
		req.EndTime,
		req.Type,
	)
	if err != nil {
		return nil, err
	}
	var list []*asset.MemberTransaction
	err = copier.Copy(&list, memberTransactionVos)
	if err != nil {
		logx.Error("copy操作失败，err=", err)
		return nil, err
	}
	return &asset.MemberTransactionList{
		List:  list,
		Total: total,
	}, nil
}

func (l *AssetLogic) GetAddress(req *asset.AssetReq) (*asset.AddressList, error) {
	addressList, err := l.memberWalletDomain.GetAllAddress(l.ctx, req.CoinName)
	if err != nil {
		return nil, err
	}
	return &asset.AddressList{
		List: addressList,
	}, nil
}
