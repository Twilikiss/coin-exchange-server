// Package logic
// @Author twilikiss 2024/5/13 23:09:09
package logic

import (
	"common/dbutils"
	"common/dbutils/tran"
	"context"
	"errors"
	"exchange/internal/domain"
	"exchange/internal/model"
	"exchange/internal/svc"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"grpc-common/ucenter/types/member"
)

type ExchangeOrderLogic struct {
	logx.Logger
	ctx                 context.Context
	svcCtx              *svc.ServiceContext
	exchangeOrderDomain *domain.ExchangeOrderDomain
	transaction         tran.Transaction // 关于事务相关的配置
	kafkaDomain         *domain.KafkaDomain
}

func NewExchangeOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeOrderLogic {
	orderDomain := domain.NewExchangeOrderDomain(svcCtx.Db)
	return &ExchangeOrderLogic{
		ctx:                 ctx,
		svcCtx:              svcCtx,
		Logger:              logx.WithContext(ctx),
		exchangeOrderDomain: orderDomain,
		transaction:         tran.NewTransaction(svcCtx.Db.Conn),
		kafkaDomain:         domain.NewKafkaDomain(svcCtx.KafkaClient, orderDomain),
	}
}

func (l *ExchangeOrderLogic) FindOrderHistory(req *order.OrderReq) (*order.OrderRes, error) {
	exchangeOrders, total, err := l.exchangeOrderDomain.FindOrderHistory(
		l.ctx,
		req.Symbol,
		req.Page,
		req.PageSize,
		req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	err = copier.Copy(&list, exchangeOrders)
	if err != nil {
		logx.Error("转换错误，err=", err)
		return nil, err
	}
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}
func (l *ExchangeOrderLogic) FindOrderCurrent(req *order.OrderReq) (*order.OrderRes, error) {
	exchangeOrders, total, err := l.exchangeOrderDomain.FindOrderCurrent(
		l.ctx,
		req.Symbol,
		req.Page,
		req.PageSize,
		req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*order.ExchangeOrder
	err = copier.Copy(&list, exchangeOrders)
	if err != nil {
		logx.Error("转换错误，err=", err)
		return nil, err
	}
	return &order.OrderRes{
		List:  list,
		Total: total,
	}, nil
}

// Add 增加委托，这部分的代码是比较繁琐和重要的
func (l *ExchangeOrderLogic) Add(req *order.OrderReq) (*order.AddOrderRes, error) {
	// 添加委托订单
	// 检查参数是否合法
	memberRes, err := l.svcCtx.MemberRPC.FindMemberById(l.ctx, &member.MemberReq{
		MemberId: req.UserId,
	})
	if err != nil {
		logx.Error("FindMemberById error, err=", err)
		return nil, errors.New("无法找到用户信息")
	}

	if memberRes.TransactionStatus == 0 {
		return nil, errors.New("此用户已经被禁止交易")
	}

	if req.Type == model.TypeMap[model.LimitPrice] && req.Price <= 0 {
		return nil, errors.New("限价模式下价格不能小于等于0")
	}

	if req.Amount <= 0 {
		return nil, errors.New("数量不能小于等于0")
	}

	exchangeCoin, err := l.svcCtx.MarketRPC.FindSymbolInfo(l.ctx, &market.MarketReq{
		Symbol: req.Symbol,
	})
	if err != nil {
		logx.Error("MarketRPC.FindSymbolInfo error, err=", err)
		return nil, errors.New("nonsupport coin")
	}
	if exchangeCoin.Exchangeable != 1 && exchangeCoin.Enable != 1 {
		return nil, errors.New("coin forbidden")
	}

	// ====================================根据数据库中比特币的相关参数来校验请求参数===========================================
	if req.Type == model.TypeMap[model.MarketPrice] && req.Direction == model.DirectionMap[model.BUY] {
		if exchangeCoin.GetMinTurnover() > 0 && req.Amount < float64(exchangeCoin.GetMinTurnover()) {
			return nil, errors.New("成交额至少是" + fmt.Sprintf("%d", exchangeCoin.GetMinTurnover()))
		}
	} else {
		if exchangeCoin.GetMaxVolume() > 0 && exchangeCoin.GetMaxVolume() < req.Amount {
			return nil, errors.New("数量超出" + fmt.Sprintf("%f", exchangeCoin.GetMaxVolume()))
		}
		if exchangeCoin.GetMinVolume() > 0 && exchangeCoin.GetMinVolume() > req.Amount {
			return nil, errors.New("数量不能低于" + fmt.Sprintf("%f", exchangeCoin.GetMinVolume()))
		}
	}

	if req.Direction == model.DirectionMap[model.SELL] && exchangeCoin.GetMinSellPrice() > 0 {
		if req.Price < exchangeCoin.GetMinSellPrice() || req.Type == model.TypeMap[model.MarketPrice] {
			return nil, errors.New("不能低于最低限价:" + fmt.Sprintf("%f", exchangeCoin.GetMinSellPrice()))
		}
	}
	if req.Direction == model.DirectionMap[model.BUY] && exchangeCoin.GetMaxBuyPrice() > 0 {
		if req.Price > exchangeCoin.GetMaxBuyPrice() || req.Type == model.TypeMap[model.MarketPrice] {
			return nil, errors.New("不能低于最高限价:" + fmt.Sprintf("%f", exchangeCoin.GetMaxBuyPrice()))
		}
	}
	//是否启用了市价买卖
	if req.Type == model.TypeMap[model.MarketPrice] {
		if req.Direction == model.DirectionMap[model.BUY] && exchangeCoin.EnableMarketBuy == 0 {
			return nil, errors.New("不支持市价购买")
		} else if req.Direction == model.DirectionMap[model.SELL] && exchangeCoin.EnableMarketSell == 0 {
			return nil, errors.New("不支持市价出售")
		}
	}
	// =================================================================================================================

	// BTC/USDT
	//基准币:USDT
	baseSymbol := exchangeCoin.GetBaseSymbol()
	//交易币:BTC
	coinSymbol := exchangeCoin.GetCoinSymbol()

	// 要注意买/卖检查的方向是不一样的
	cc := baseSymbol
	if req.Direction == model.DirectionMap[model.SELL] {
		//根据交易币查询
		cc = coinSymbol
	}

	// 查询对应货币的相关信息
	coin, err := l.svcCtx.MarketRPC.FindCoinInfo(l.ctx, &market.MarketReq{
		Unit: cc,
	})

	if err != nil || coin == nil {
		logx.Error("MarketRPC.FindCoinInfo error, err=", err)
		return nil, errors.New("nonsupport coin")
	}

	// TODO 根据用户的钱包来进一步校验
	//查询用户钱包 BTC/USDT
	baseWallet, err := l.svcCtx.AssetRPC.FindWalletBySymbol(l.ctx, &asset.AssetReq{
		UserId:   req.UserId,
		CoinName: baseSymbol,
	})
	if err != nil {
		return nil, errors.New("no wallet")
	}
	exCoinWallet, err := l.svcCtx.AssetRPC.FindWalletBySymbol(l.ctx, &asset.AssetReq{
		UserId:   req.UserId,
		CoinName: coinSymbol,
	})
	if err != nil {
		return nil, errors.New("no wallet")
	}

	if baseWallet.IsLock == 1 || exCoinWallet.IsLock == 1 {
		return nil, errors.New("wallet locked")
	}
	// 查询用户的该coin的委托数量，检查委托数量限制
	count, err := l.exchangeOrderDomain.FindCurrentTradingCount(l.ctx, req.UserId, req.Symbol, req.Direction)
	if err != nil {
		return nil, err
	}
	if exchangeCoin.GetMaxTradingOrder() > 0 && count >= exchangeCoin.GetMaxTradingOrder() {
		return nil, errors.New("超过最大挂单数量 " + fmt.Sprintf("%d", exchangeCoin.GetMaxTradingOrder()))
	}

	// =======================生成订单========================= //
	exchangeOrder := model.NewOrder()
	exchangeOrder.MemberId = req.UserId
	exchangeOrder.Symbol = req.Symbol
	exchangeOrder.BaseSymbol = baseSymbol
	exchangeOrder.CoinSymbol = coinSymbol
	typeCode := model.TypeMap.Code(req.Type)
	exchangeOrder.Type = typeCode
	directionCode := model.DirectionMap.Code(req.Direction)
	exchangeOrder.Direction = directionCode
	if exchangeOrder.Type == model.MarketPrice {
		// 如果使用的市价委托，价格就取决实际成交时获得的市场价格，price为0
		exchangeOrder.Price = 0
	} else {
		exchangeOrder.Price = req.Price
	}
	exchangeOrder.UseDiscount = "0"
	exchangeOrder.Amount = req.Amount
	// =======================生成订单========================= //

	// =====================保存订单到数据库===================== //
	// TODO 目前一下功能上存在尚存一些bug
	// 保存订单到数据库，发送消息到kafka，ucenter（钱包服务）接收到消息，进行资金的冻结，
	// 如果出现消息发送失败还需要考虑回滚
	// AddOrder 保存订单 计算所需要的钱
	l.transaction.Action(func(conn dbutils.DbConn) error {
		money, err := l.exchangeOrderDomain.AddOrder(l.ctx, conn, exchangeOrder, exchangeCoin, baseWallet, exCoinWallet)
		if err != nil {
			return errors.New("订单提交失败")
		}
		//通过kafka发消息 订单创建成功的消息 对应用户的钱包就要冻结相关的资产
		logx.Info("+++++++++++++++++++通过Kafka发送订单信息start++++++++++++++++++++++++++")
		err = l.kafkaDomain.SendOrderAdd(
			"add-exchange-order", // 设置发送该类消息的topic为add-exchange-order
			req.UserId,
			exchangeOrder.OrderId,
			money,
			req.Symbol,
			exchangeOrder.Direction,
			baseSymbol,
			coinSymbol)
		if err != nil {
			return errors.New("发消息失败")
		}
		logx.Info("+++++++++++++++++++通过Kafka发送订单信息end++++++++++++++++++++++++++")
		return nil
	})
	// 如果出现错误就返回错误
	if err != nil {
		return nil, err
	}
	// 成功就需要返回OrderId
	return &order.AddOrderRes{
		OrderId: exchangeOrder.OrderId,
	}, nil
}

func (l *ExchangeOrderLogic) FindByOrderId(req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	o, err := l.exchangeOrderDomain.FindOrderByOrderId(l.ctx, req.OrderId)
	oo := &order.ExchangeOrderOrigin{}
	err = copier.Copy(oo, o)
	if err != nil {
		logx.Error("copier.Copy转换失败, err", err)
		return nil, errors.New("copier.Copy转换失败")
	}
	return oo, err
}

func (l *ExchangeOrderLogic) CancelOrder(req *order.OrderReq) (*order.CancelOrderRes, error) {
	err := l.exchangeOrderDomain.UpdateStatusCancel(l.ctx, req.OrderId)
	return nil, err
}
