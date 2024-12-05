package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market/internal/domain"
	"market/internal/svc"
	"time"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeCoinDomain *domain.ExchangeCoinDomain
	marketDomain       *domain.MarketDomain
	coinDomain         *domain.CoinDomain
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeCoinDomain: domain.NewExchangeCoinDomain(svcCtx.Db),
		marketDomain:       domain.NewMarketDomain(svcCtx.MongoClient),
		coinDomain:         domain.NewCoinDomain(svcCtx.Db),
	}
}

// FindSymbolThumbTrend 函数用于处理市场请求，根据给定的市场请求参数 req，查询并返回指定符号的价格趋势信息。
//
// 参数：
//
//	req *market.MarketReq: 包含市场请求参数的指针，用于指定查询的条件。
//
// 返回值：
//
//	*market.SymbolThumbRes: 指向市场符号缩略图响应的指针，包含查询结果。
//	error: 如果在查询过程中发生错误，则返回相应的错误信息；否则返回 nil。
//
// 说明：
//
//	该函数首先记录接收到 RPC 远程调用请求的信息。
//	然后，使用两个上下文和取消函数来设置查询操作的超时时间（5秒）。
//	第一个查询操作是调用 exchangeCoinDomain 的 FindVisible 方法，获取可见的交易所硬币信息。
//	第二个查询操作是根据请求中的周期（Period）参数，从市场域（marketDomain）中查询指定周期内的符号价格趋势信息。
//	如果请求中未指定周期，则默认使用 "1H"（1小时）作为查询周期。
//	查询结果通过 copier.Copy 函数复制到 market.SymbolThumbRes 结构的 List 字段中。
//	如果在复制过程中发生错误，将记录错误信息并返回该错误。
//	最后，返回包含查询结果的数据指针和 nil 错误值（表示成功）。
func (l *MarketLogic) FindSymbolThumbTrend(req *market.MarketReq) (*market.SymbolThumbRes, error) {
	logx.Info("接收到RPC远程调用请求，处理中")

	ctx1, cancel1 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel1()
	exchangeCoins := l.exchangeCoinDomain.FindVisible(ctx1)

	//查询mongo中相应的数据
	//查询1H间隔的 可以根据时间来进行查询 当天的价格变化趋势
	ctx2, cancel2 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel2()

	// 将原项目修改自定义获取时间间隔 period
	period := req.GetPeriod()
	if period == "" {
		period = "1H"
	}
	coinThumbs := l.marketDomain.SymbolThumbTrend(ctx2, period, exchangeCoins)
	data := &market.SymbolThumbRes{}
	list := &data.List
	err := copier.Copy(list, coinThumbs)
	if err != nil {
		logx.Error("转换出错，error=", err)
		return nil, err
	}
	//coinThumbs := make([]*market.CoinThumb, len(exchangeCoins))
	//for i, v := range exchangeCoins {
	//	trend := make([]float64, 0)
	//	for p := 0; p <= 100; p++ {
	//		trend = append(trend, rand.Float64())
	//	}
	//	ct := &market.CoinThumb{}
	//	ct.Symbol = v.Symbol
	//	ct.Trend = trend
	//	coinThumbs[i] = ct
	//}

	return data, nil
}

func (l *MarketLogic) FindSymbolInfo(req *market.MarketReq) (*market.ExchangeCoin, error) {
	logx.Info("接收到RPC远程调用请求，处理中")

	ctx1, cancel1 := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel1()
	exchangeCoins, err := l.exchangeCoinDomain.FindBySymbol(req, ctx1)
	if err != nil {
		return nil, err
	}
	data := &market.ExchangeCoin{}
	err = copier.Copy(data, exchangeCoins)
	if err != nil {
		logx.Error("转换出错，error=", err)
		return nil, err
	}
	return data, nil
}

func (l *MarketLogic) FindCoinInfo(req *market.MarketReq) (*market.Coin, error) {
	coin, err := l.coinDomain.FindCoinInfo(l.ctx, req.Unit)
	if err != nil {
		return nil, err
	}
	mc := &market.Coin{}
	if err := copier.Copy(mc, coin); err != nil {
		logx.Error("转换出错，error=", err)
		return nil, err
	}
	return mc, nil
}

func (l *MarketLogic) HistoryKline(req *market.MarketReq) (*market.HistoryRes, error) {
	// 具体到MongoDB中去查询
	// 按照时间范围查询，并排序（按照时间升序排序）
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()

	var period string
	if req.Resolution == "60" {
		period = "1H"
	} else if req.Resolution == "30" {
		period = "30m"
	} else if req.Resolution == "15" {
		period = "15m"
	} else if req.Resolution == "5" {
		period = "5m"
	} else if req.Resolution == "1" {
		period = "1m"
	} else if req.Resolution == "1D" {
		period = "1D"
	} else if req.Resolution == "1W" {
		period = "1W"
	} else if req.Resolution == "1M" {
		period = "1M"
	} else {
		// 如果都不匹配就默认查询1H的数据
		period = "1H"
	}
	histories, err := l.marketDomain.HistoryKline(ctx, req.Symbol, req.From, req.To, period)
	if err != nil {
		logx.Error("查询KlineHistory失败，err=", err)
		return nil, err
	}
	list := make([]*market.History, len(histories))
	err = copier.Copy(&list, histories)
	if err != nil {
		logx.Error("转换出错，error=", err)
		return nil, err
	}
	return &market.HistoryRes{
		List: list,
	}, nil
}

func (l *MarketLogic) FindExchangeCoinVisible(req *market.MarketReq) (*market.ExchangeCoinRes, error) {
	exchangeCoins := l.exchangeCoinDomain.FindVisible(l.ctx)
	var list []*market.ExchangeCoin
	err := copier.Copy(&list, exchangeCoins)
	if err != nil {
		logx.Error("copier.Copy操作失败，err=", err)
		return nil, err
	}
	return &market.ExchangeCoinRes{
		List: list,
	}, nil
}

func (l *MarketLogic) FindAllCoin(req *market.MarketReq) (*market.CoinList, error) {
	coinList, err := l.coinDomain.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var list []*market.Coin
	err = copier.Copy(&list, coinList)
	if err != nil {
		logx.Error("格式转换出错， err=", err)
	}
	return &market.CoinList{
		List: list,
	}, nil
}

func (l *MarketLogic) FindById(req *market.MarketReq) (*market.Coin, error) {
	coin, err := l.coinDomain.FindCoinById(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	mc := &market.Coin{}
	if err := copier.Copy(mc, coin); err != nil {
		logx.Error("转换出错，error=", err)
		return nil, err
	}
	return mc, nil
}
