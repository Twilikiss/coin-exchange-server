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
