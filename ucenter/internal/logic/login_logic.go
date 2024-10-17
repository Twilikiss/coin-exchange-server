package logic

import (
	"common/tools"
	"context"
	"errors"
	"grpc-common/ucenter/types/login"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/model"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {

	logx.Info("接收到RPC远程调用请求，处理中")

	// 1.我们在进行注册操作时要先校验我们的人机校验是否通过
	isVerify := l.CaptchaDomain.Verify(
		in.GetCaptcha().GetServer(),
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.SecretKey,
		in.GetCaptcha().GetToken(),
		2,
		in.Ip)
	if !isVerify {
		return nil, errors.New("人机校验未通过")
	}
	logx.Info("人机校验通过")

	// 2.校验用户密码
	ctx1, cancelFunc1 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc1()

	member, err := l.MemberDomain.FindByUserName(ctx1, in.Username)
	if err != nil {
		logx.Error("数据库异常，error=", err)
		return nil, errors.New("数据库异常")
	}
	if member == nil {
		return nil, errors.New("此用户未注册")
	}

	pwd := member.Password
	salt := member.Salt

	isVerify = tools.Verify(in.Password, salt, pwd, nil)

	if !isVerify {
		return nil, errors.New("密码错误")
	}

	logx.Info("密码校验通过")

	// 3.登录成功，生成token，前端传递token，后续校验token即可
	// 借助jwt的技术，JWT由3部分组成：标头(Header)、有效载荷(Payload)和签名(Signature)
	secret := l.svcCtx.Config.JWT.AccessSecret
	expire := l.svcCtx.Config.JWT.AccessExpire

	token, err := tools.GetJwtToken(secret, time.Now().Unix(), expire, member.Id)
	if err != nil {
		logx.Error("jwt生成异常，error=", err)
		return nil, errors.New("token生成错误")
	}

	// 对于登录次数而言，我们并不需要特别关心它实时的值
	go func(member *model.Member) {
		ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
		defer cancelFunc()
		err := l.MemberDomain.UpdateLoginCountById(ctx, member.Id, 1)
		if err != nil {
			logx.Error("更新loginCount数值异常，error=", err)
		}
	}(member)

	loginCount := member.LoginCount + 1
	// 4.最后我们将返回我们的登录信息
	return &login.LoginRes{
		Username:      member.Username,
		Token:         token,
		MemberLevel:   member.MemberLevelStr(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		Id:            member.Id,
		LoginCount:    int32(loginCount),
		SuperPartner:  member.SuperPartner,
		MemberRate:    member.MemberRate(),
	}, nil
}
