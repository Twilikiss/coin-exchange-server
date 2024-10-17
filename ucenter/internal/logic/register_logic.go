package logic

import (
	"common/tools"
	"context"
	"errors"
	"fmt"
	"grpc-common/ucenter/types/register"
	"strconv"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER::"
const MinWaitingTime = 60

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {

	logx.Info("接收到RPC远程调用请求，处理中")

	ctx1, cancelFunc1 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc1()

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

	// 2.查验短信验证码
	country := in.GetCountry()
	phone := in.GetPhone()

	key := RegisterCacheKey + country + "::" + phone

	// 判断是否存在该验证码
	if isExist, _ := l.svcCtx.Cache.ExistsCtx(ctx1, key); !isExist {
		logx.Errorf("所查询的key=%s，验证码不存在", key)
		return nil, errors.New("验证失败")
	}

	// 设置context，设置查询redis的超时时间为5s
	ctx2, cancelFunc2 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc2()

	code, err := l.svcCtx.Cache.GetCtx(ctx2, key)
	if err != nil {
		logx.Error("redis操作异常，error=", err)
		return nil, errors.New("系统异常，请联系管理员")
	}
	if in.GetCode() != code {
		logx.Errorf("验证码错误，input=%s, code=%s\n", in.GetCode(), code)
		return nil, errors.New("验证失败")
	}

	// 验证码验证过后就要马上销毁防止验证码重复验证
	ctx3, cancelFunc3 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc3()
	_, err = l.svcCtx.Cache.DelCtx(ctx3, key)
	if err != nil {
		logx.Error("redis操作异常，error=", err)
		return nil, errors.New("系统异常，请联系管理员")
	}

	// 3.验证码通过，进行注册，我们要先判断手机号是否被注册过了
	mem, err := l.MemberDomain.FindByPhone(l.ctx, phone)

	if err != nil {
		return nil, errors.New("数据库读取异常，请联系管理员")
	}

	if mem != nil {
		return nil, errors.New("此用户已被注册")
	}

	logx.Info("存入数据库中")

	// 4.存入数据库

	ctx4, cancelFunc4 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc4()

	err = l.MemberDomain.Register(
		ctx4,
		in.GetUsername(),
		in.GetPhone(),    // 手机号码
		in.GetPassword(), // 密码
		in.GetCountry(),  // 国家
		in.GetSuperPartner(),
		in.GetPromotion(),
	)

	if err != nil {
		return nil, errors.New("数据库异常，注册失败")
	}

	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(in *register.CodeReq) (*register.NoRes, error) {

	logx.Info("接收到RPC远程调用请求:SendCode，处理中")

	ctx1, cancelFunc1 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc1()

	//* 收到手机号和国家标识
	phone := in.GetPhone()
	country := in.GetCountry()

	key := RegisterCacheKey + country + "::" + phone
	// 先检查redis该手机号是否存在有未失效的验证码
	temp, _ := l.svcCtx.Cache.GetCtx(ctx1, key)
	if temp != "" {
		ttl, err := l.svcCtx.Cache.Ttl(key)

		if err != nil {
			return nil, err
		}

		t := MinWaitingTime - (5*60 - ttl)

		if (5*60 - ttl) < 60 {
			return nil, errors.New(fmt.Sprintf("请求过于频繁，请等待%d秒后重试", t))
		}
	}
	//* 生成验证码
	code := tools.Gen4Num()

	// 根据对应的国家和手机号调用对应的短信平台发送验证码
	go func(code int) {
		tools.SendMessage(phone, strconv.Itoa(code))
	}(code)

	//* 将验证码存入redis，过期时间5分钟
	// 设置context，设置存入redis的超时时间为5s
	ctx2, cancelFunc2 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc2()

	err := l.svcCtx.Cache.SetexCtx(ctx2, key, strconv.Itoa(code), 5*60)
	if err != nil {
		return nil, errors.New("验证码存入Cache失败")
	}
	//* 返回成功
	return &register.NoRes{}, nil
}
