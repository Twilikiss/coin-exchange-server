// Package processor
// @Author twilikiss 2024/7/1 16:41:41
package processor

import (
	"common/dbutils"
	"common/op"
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/domain"
	"exchange/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"sort"
	"sync"
	"time"
)

// CoinTradeFactory 工厂 专门生产对应symbol的交易引擎
type CoinTradeFactory struct {
	tradeMap map[string]*CoinTrade // 对应的交易对的引擎
	mux      sync.RWMutex          // 读写锁
}

func NewCoinTradeFactory() *CoinTradeFactory {
	return &CoinTradeFactory{
		tradeMap: make(map[string]*CoinTrade),
	}
}

func (c *CoinTradeFactory) Init(marketRPC mclient.Market, cli *database.KafkaClient, db *dbutils.ElysiaDB) {
	// 初始化操作
	// 查询所有的exchange_coin的内容，循环创建我们的交易引擎
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exchangeCoinRes, err := marketRPC.FindExchangeCoinVisible(ctx, &market.MarketReq{})
	if err != nil {
		logx.Error("调用FindExchangeCoinVisible()失败， err=", err)
		return
	}
	for _, v := range exchangeCoinRes.List {
		c.AddCoinTrade(v.Symbol, NewCoinTrade(v.Symbol, cli, db))
	}
}
func (c *CoinTradeFactory) AddCoinTrade(symbol string, ct *CoinTrade) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.tradeMap[symbol] = ct
}

func (c *CoinTradeFactory) GetCoinTrade(symbol string) *CoinTrade {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.tradeMap[symbol]
}

// TradeTimeQueue 交易时间排序队列，按照时间升序排序
type TradeTimeQueue []*model.ExchangeOrder

func (t TradeTimeQueue) Len() int {
	return len(t)
}
func (t TradeTimeQueue) Less(i, j int) bool {
	//升序
	return t[i].Time < t[j].Time
}
func (t TradeTimeQueue) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type LimitPriceQueue struct {
	mux  sync.RWMutex
	list TradeQueue
}
type LimitPriceMap struct {
	price float64
	list  []*model.ExchangeOrder
}

// TradeQueue 限价交易队列
type TradeQueue []*LimitPriceMap

func (t TradeQueue) Len() int {
	return len(t)
}
func (t TradeQueue) Less(i, j int) bool {
	//降序
	return t[i].price > t[j].price
}
func (t TradeQueue) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// CoinTrade 撮合交易引擎 每一个交易对都有各自的一个引擎
type CoinTrade struct {
	symbol          string
	buyMarketQueue  TradeTimeQueue // 买的市价队列, 按照时间由低到高排序
	bmMux           sync.RWMutex
	sellMarketQueue TradeTimeQueue // 卖的市价队列, 按照时间由低到高排序
	smMux           sync.RWMutex
	buyLimitQueue   *LimitPriceQueue // 买的限价队列，从卖家角度来说，希望该队列从高到低
	sellLimitQueue  *LimitPriceQueue // 卖的限价队列，从买方角度来说，希望该队列从低到高
	buyTradePlate   *TradePlate      // 买盘
	sellTradePlate  *TradePlate      // 卖盘
	kafkaClient     *database.KafkaClient
	db              *dbutils.ElysiaDB // 用于数据库查询交易中的订单
}

// TradePlate 盘口信息
type TradePlate struct {
	Items     []*TradePlateItem `json:"items"`
	Symbol    string
	direction int
	maxDepth  int
	mux       sync.RWMutex
}

type TradePlateItem struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

func (p *TradePlate) Add(order *model.ExchangeOrder) {
	if order.Direction != p.direction {
		logx.Error("内部错误")
		return
	}
	p.mux.Lock()
	defer p.mux.Unlock()
	if order.Type == model.MarketPrice {
		//市价不进入买卖盘
		//买卖盘 委托订单的形式下产生的一个金融说辞
		//委托 基本上可以认定是成交的，一旦发生委托  那么就意味着 买方和卖方市场 形成了 100 100  0 200
		//买卖盘 委托 已成交 未成交 全部 这些都可以叫买卖盘
		//看我们的应用 给用户展示哪方面的数据 我们选择展示 未成交的买单（20000） 和 卖单
		return
	}
	size := len(p.Items)
	if size > 0 {
		for _, v := range p.Items {
			if v.Price == order.Price {
				//order.Amount= 20  order.TradedAmount=10 10 Amount:买入或卖出量 TradeAmount:成交额
				v.Amount = op.FloorFloat(v.Amount+(order.Amount-order.TradedAmount), 8)
				return
			}
		}
	}

	if size < p.maxDepth {
		tpi := &TradePlateItem{
			Amount: op.FloorFloat(order.Amount-order.TradedAmount, 8),
			Price:  order.Price,
		}
		p.Items = append(p.Items, tpi)
	}
}

func NewTradePlate(symbol string, direction int) *TradePlate {
	return &TradePlate{
		Symbol:    symbol,
		direction: direction,
		maxDepth:  100,
	}
}

// Trade 撮合交易的核心代码
func (t *CoinTrade) Trade(exchangeOrder *model.ExchangeOrder) {
	//1. 当订单进来之后，确定 buy还是sell
	//2. 确定 市价 还是限价
	//3. buy 和 sell队列进行匹配
	//4. sell 和买的队列进行匹配
	//5. exchangeOrder 还未交易的，放入买卖盘 同时放入 交易引擎的队列中 等待下次匹配
	//6. 订单就会更新 订单的状态要变 冻结的金额 扣除等等
	//if exchangeOrder.Direction == model.BUY {
	//	// 放入买盘和卖盘，然后发送到前端进行展示
	//	t.buyTradePlate.Add(exchangeOrder)
	//
	//	// 将买卖盘数据发送到kafka中
	//	t.sendTradePlateMsg(t.buyTradePlate)
	//} else {
	//	t.sellTradePlate.Add(exchangeOrder)
	//	// 将买卖盘数据发送到kafka中
	//	t.sendTradePlateMsg(t.sellTradePlate)
	//}

	//exchangeOrder 买 和卖的队列进行匹配 还是卖 和买的队列进行匹配
	var limitPriceList *LimitPriceQueue
	var marketPriceList TradeTimeQueue
	if exchangeOrder.Direction == model.BUY {
		limitPriceList = t.sellLimitQueue
		marketPriceList = t.sellMarketQueue
	} else {
		limitPriceList = t.buyLimitQueue
		marketPriceList = t.buyMarketQueue
	}
	if exchangeOrder.Type == model.MarketPrice {
		//先处理市价 市价订单和限价的订单进行匹配
		t.matchMarketPriceWithLP(limitPriceList, exchangeOrder)
	} else {
		//限价单 先于限价单进行成交 如果未成交 继续与市价单进行成交
		t.matchLimitPriceWithLP(limitPriceList, exchangeOrder)
		if exchangeOrder.Status == model.Trading {
			t.matchLimitPriceWithMP(marketPriceList, exchangeOrder)
		}
		if exchangeOrder.Status == model.Trading {
			//证明还未交易完成
			t.addLimitQueue(exchangeOrder)
			// 因为修改了和买卖盘相关的数据，需要把数据发送到前端
			if exchangeOrder.Direction == model.BUY {
				t.sendTradePlateMsg(t.buyTradePlate)
			} else {
				t.sendTradePlateMsg(t.sellTradePlate)
			}
		}
	}
}

func (t *CoinTrade) init() {
	t.buyTradePlate = NewTradePlate(t.symbol, model.BUY)
	t.sellTradePlate = NewTradePlate(t.symbol, model.SELL)
	t.buyLimitQueue = &LimitPriceQueue{}
	t.sellLimitQueue = &LimitPriceQueue{}
	t.initData()
}

type TradePlateResult struct {
	Direction    string            `json:"direction"`
	MaxAmount    float64           `json:"maxAmount"`
	MinAmount    float64           `json:"minAmount"`
	HighestPrice float64           `json:"highestPrice"`
	LowestPrice  float64           `json:"lowestPrice"`
	Symbol       string            `json:"symbol"`
	Items        []*TradePlateItem `json:"items"`
}

func (p *TradePlate) AllResult() *TradePlateResult {
	result := &TradePlateResult{}
	direction := model.DirectionMap.Value(p.direction)
	result.Direction = direction
	result.MaxAmount = p.getMaxAmount()
	result.MinAmount = p.getMinAmount()
	result.HighestPrice = p.getHighestPrice()
	result.LowestPrice = p.getLowestPrice()
	result.Symbol = p.Symbol
	result.Items = p.Items
	return result
}

func (p *TradePlate) Result(num int) *TradePlateResult {
	if num > len(p.Items) {
		num = len(p.Items)
	}
	result := &TradePlateResult{}
	direction := model.DirectionMap.Value(p.direction)
	result.Direction = direction
	result.MaxAmount = p.getMaxAmount()
	result.MinAmount = p.getMinAmount()
	result.HighestPrice = p.getHighestPrice()
	result.LowestPrice = p.getLowestPrice()
	result.Symbol = p.Symbol
	result.Items = p.Items[:num]
	return result
}

func (p *TradePlate) getMaxAmount() float64 {
	if len(p.Items) <= 0 {
		return 0
	}
	var amount float64 = 0
	for _, v := range p.Items {
		if v.Amount > amount {
			amount = v.Amount
		}
	}
	return amount
}

func (p *TradePlate) getMinAmount() float64 {
	if len(p.Items) <= 0 {
		return 0
	}
	var amount float64 = p.Items[0].Amount
	for _, v := range p.Items {
		if v.Amount < amount {
			amount = v.Amount
		}
	}
	return amount
}

func (p *TradePlate) getHighestPrice() float64 {
	if len(p.Items) <= 0 {
		return 0
	}
	var price float64 = 0
	for _, v := range p.Items {
		if v.Price > price {
			price = v.Price
		}
	}
	return price
}
func (p *TradePlate) getLowestPrice() float64 {
	if len(p.Items) <= 0 {
		return 0
	}
	var price float64 = p.Items[0].Price
	for _, v := range p.Items {
		if v.Price < price {
			price = v.Price
		}
	}
	return price
}

func (p *TradePlate) Remove(order *model.ExchangeOrder, amount float64) {
	for i, v := range p.Items {
		if v.Price == order.Price {
			v.Amount = op.SubFloor(v.Amount, amount, 8)
			if v.Amount <= 0 {
				p.Items = append(p.Items[:i], p.Items[i+1:]...)
			}
			break
		}
	}
}
func (t *CoinTrade) sendTradePlateMsg(plate *TradePlate) {
	// 构造好需要返回的数据
	result := plate.Result(24)
	marshal, _ := json.Marshal(result)
	data := database.KafkaData{
		Topic: "exchange_order_trade_plate",
		Key:   []byte(plate.Symbol),
		Data:  marshal,
	}
	// 这里如果需要保证买卖盘数据必须发送成功的话可以设置错误重试
	// 这里我们的买卖盘数据可以容忍少量的数据发送失败的情况，所以就没采用
	err := t.kafkaClient.SendSync(data)
	if err != nil {
		logx.Error(err)
	} else {
		logx.Info("======exchange_order_trade_plate发送成功....==========")
	}
}

func (t *CoinTrade) initData() {
	orderDomain := domain.NewExchangeOrderDomain(t.db)
	//应该去查询对应symbol的订单 将其赋值到coinTrade里面的各个队列中，同时加入买卖盘
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	exchangeOrders, err := orderDomain.FindOrderListBySymbol(ctx, t.symbol, model.Trading)
	if err != nil {
		logx.Error("CoinTrade initData() has a error, err=", err)
		return
	}
	for _, v := range exchangeOrders {
		if v.Type == model.MarketPrice {
			if v.Direction == model.BUY {
				t.bmMux.Lock()
				// 放入我们的市价买入队列
				t.buyMarketQueue = append(t.buyMarketQueue, v)
				t.bmMux.Unlock()
				continue
			}
			if v.Direction == model.SELL {
				t.smMux.Lock()
				// 放入我们的市价卖出队列
				t.sellMarketQueue = append(t.sellMarketQueue, v)
				t.smMux.Unlock()
				continue
			}
			//市价单 不进入买卖盘的
		} else if v.Type == model.LimitPrice {

			// 如果是限价买入
			if v.Direction == model.BUY {
				t.buyLimitQueue.mux.Lock()
				//deal
				isPut := false
				for _, o := range t.buyLimitQueue.list {
					if o.price == v.Price { // 存在同样价位的list可以直接添加到对应list里
						o.list = append(o.list, v)
						isPut = true
						break
					}
				}
				if !isPut {
					lpm := &LimitPriceMap{
						price: v.Price,
						list:  []*model.ExchangeOrder{v},
					}
					t.buyLimitQueue.list = append(t.buyLimitQueue.list, lpm)
				}
				// 放入对应的买卖盘
				t.buyTradePlate.Add(v)
				t.buyLimitQueue.mux.Unlock()
			} else if v.Direction == model.SELL {

				// 如果是限价卖出
				t.sellLimitQueue.mux.Lock()
				//deal
				isPut := false
				for _, o := range t.sellLimitQueue.list {
					if o.price == v.Price {
						o.list = append(o.list, v)
						isPut = true
						break
					}
				}
				if !isPut {
					lpm := &LimitPriceMap{
						price: v.Price,
						list:  []*model.ExchangeOrder{v},
					}
					t.sellLimitQueue.list = append(t.sellLimitQueue.list, lpm)
				}

				// 放入对应的买卖盘
				t.sellTradePlate.Add(v)
				t.sellLimitQueue.mux.Unlock()
			}
		}
	}
	//排序
	sort.Sort(t.buyMarketQueue)
	sort.Sort(t.sellMarketQueue)
	sort.Sort(t.buyLimitQueue.list)                //从高到低
	sort.Sort(sort.Reverse(t.sellLimitQueue.list)) //从低到高
	if len(exchangeOrders) > 0 {
		t.sendTradePlateMsg(t.buyTradePlate)
		t.sendTradePlateMsg(t.sellTradePlate)
	}
}

// matchMarketPriceWithLP focusedOrder 市价单
func (t *CoinTrade) matchMarketPriceWithLP(lpList *LimitPriceQueue, focusedOrder *model.ExchangeOrder) {
	lpList.mux.Lock()
	defer lpList.mux.Unlock()

	var delOrders []string
	buyNotify := false
	sellNotify := false
	var completeOrders []*model.ExchangeOrder

	//如果是买  卖的队列价格是从低到高  如果是卖 买的队列 价格是从高到低
	for _, v := range lpList.list {
		for _, matchOrder := range v.list {
			// 自己的订单就不处理了
			if matchOrder.MemberId == focusedOrder.MemberId {
				continue
			}
			//matchOrder和focusedOrder 是否匹配
			price := matchOrder.Price
			// 计算交易额数量
			matchAmount := op.SubFloor(matchOrder.Amount, matchOrder.TradedAmount, 8)
			if matchAmount <= 0 {
				continue
			}
			focusedAmount := op.SubFloor(focusedOrder.Amount, focusedOrder.TradedAmount, 8)

			//市价单 买的时候 amount是 usdt 金额 这时候我们需要计算数量
			if focusedOrder.Direction == model.BUY {
				// 这里通过计算计算成交的金额总数 / 单价 = 实际交易量
				focusedAmount = op.DivFloor(op.SubFloor(focusedOrder.Amount, focusedOrder.Turnover, 8), price, 8)
			}

			if matchAmount >= focusedAmount {
				//订单直接就交易完成了 能满足
				turnover := op.MulFloor(price, focusedAmount, 8)
				matchOrder.TradedAmount = op.AddFloor(matchOrder.TradedAmount, focusedAmount, 8)
				matchOrder.Turnover = op.AddFloor(matchOrder.Turnover, turnover, 8)
				if op.SubFloor(matchOrder.Amount, matchOrder.TradedAmount, 8) <= 0 {
					matchOrder.Status = model.Completed
					//从队列进行删除
					delOrders = append(delOrders, matchOrder.OrderId)
				}
				focusedOrder.TradedAmount = op.AddFloor(focusedOrder.TradedAmount, focusedAmount, 8)
				focusedOrder.Turnover = op.AddFloor(focusedOrder.Turnover, turnover, 8)
				focusedOrder.Status = model.Completed

				// 交易完成后也需要通知买卖盘，对我们买卖盘的数据进行修改
				if matchOrder.Direction == model.BUY {
					t.buyTradePlate.Remove(matchOrder, focusedAmount)
					buyNotify = true
				} else {
					t.sellTradePlate.Remove(matchOrder, focusedAmount)
					sellNotify = true
				}
				break
			} else {
				// 当前订单不满足交易额，继续进行匹配
				turnover := op.MulFloor(price, matchAmount, 8)
				matchOrder.TradedAmount = op.AddFloor(matchOrder.TradedAmount, matchAmount, 8)
				matchOrder.Turnover = op.AddFloor(matchOrder.Turnover, turnover, 8)
				matchOrder.Status = model.Completed
				completeOrders = append(completeOrders, matchOrder)
				//从队列进行删除
				delOrders = append(delOrders, matchOrder.OrderId)

				focusedOrder.TradedAmount = op.AddFloor(focusedOrder.TradedAmount, matchAmount, 8)
				focusedOrder.Turnover = op.AddFloor(focusedOrder.Turnover, turnover, 8)

				if matchOrder.Direction == model.BUY {
					t.buyTradePlate.Remove(matchOrder, matchAmount)
					buyNotify = true
				} else {
					t.sellTradePlate.Remove(matchOrder, matchAmount)
					sellNotify = true
				}
				continue
			}
		}
	}

	//处理队列中 已经完成的订单进行删除
	for _, orderId := range delOrders {
		for _, v := range lpList.list {
			for index, matchOrder := range v.list {
				if orderId == matchOrder.OrderId {
					v.list = append(v.list[:index], v.list[index+1:]...)
					break
				}
			}
		}
	}

	//判断是否订单完成
	if focusedOrder.Status == model.Trading {
		//未完成 放入队列
		t.addMarketQueue(focusedOrder)
	}
	//通知买or卖盘更新
	if buyNotify {
		t.sendTradePlateMsg(t.buyTradePlate)
	}
	if sellNotify {
		t.sendTradePlateMsg(t.sellTradePlate)
	}
}

func (t *CoinTrade) addMarketQueue(order *model.ExchangeOrder) {
	if order.Type != model.MarketPrice {
		return
	}
	if order.Direction == model.BUY {
		t.buyMarketQueue = append(t.buyMarketQueue, order)
		// 重新按照时间进行排序
		sort.Sort(t.buyMarketQueue)
	} else {
		t.sellMarketQueue = append(t.sellMarketQueue, order)
		// 重新按照时间进行排序
		sort.Sort(t.sellMarketQueue)
	}
}

// matchLimitPriceWithLP focusedOrder 限价单
func (t *CoinTrade) matchLimitPriceWithLP(lpList *LimitPriceQueue, focusedOrder *model.ExchangeOrder) {
	lpList.mux.Lock()
	defer lpList.mux.Unlock()
	var delOrders []string
	buyNotify := false
	sellNotify := false
	var completeOrders []*model.ExchangeOrder
	//如果是买  卖的队列价格是从低到高  如果是卖 买的队列 价格是从高到低
	for _, v := range lpList.list {
		for _, matchOrder := range v.list {
			if matchOrder.MemberId == focusedOrder.MemberId {
				//自己的订单就不处理了
				continue
			}
			//如果是买  卖队列是 价格从低到高  如果买的价格比卖的价格还低 无法成交  100  120
			// 可以这么理解， 卖队列最低价都大于我们的买的价格，暂时没有成交的可能
			if model.BUY == focusedOrder.Direction {
				if focusedOrder.Price < matchOrder.Price {
					break
				}
			}
			//如果是卖 买队列是 价格从高到低   如果卖的价格比买的价格还高 无法成交  100  90 80 70
			if model.SELL == focusedOrder.Direction {
				if focusedOrder.Price > matchOrder.Price {
					break
				}
			}
			//matchOrder和focusedOrder 是否匹配
			price := matchOrder.Price
			//计算可交易的数量
			matchAmount := op.SubFloor(matchOrder.Amount, matchOrder.TradedAmount, 8)
			if matchAmount <= 0 {
				continue
			}
			focusedAmount := op.SubFloor(focusedOrder.Amount, focusedOrder.TradedAmount, 8)
			if matchAmount >= focusedAmount {
				//订单直接就交易完成了 能满足
				turnover := op.MulFloor(price, focusedAmount, 8)
				matchOrder.TradedAmount = op.AddFloor(matchOrder.TradedAmount, focusedAmount, 8)
				matchOrder.Turnover = op.AddFloor(matchOrder.Turnover, turnover, 8)
				if op.SubFloor(matchOrder.Amount, matchOrder.TradedAmount, 8) <= 0 {
					matchOrder.Status = model.Completed
					//从队列进行删除
					delOrders = append(delOrders, matchOrder.OrderId)
					completeOrders = append(completeOrders, matchOrder)
				}
				focusedOrder.TradedAmount = op.AddFloor(focusedOrder.TradedAmount, focusedAmount, 8)
				focusedOrder.Turnover = op.AddFloor(focusedOrder.Turnover, turnover, 8)
				focusedOrder.Status = model.Completed
				// 添加到我们的完成队列中
				completeOrders = append(completeOrders, focusedOrder)
				if matchOrder.Direction == model.BUY {
					t.buyTradePlate.Remove(matchOrder, focusedAmount)
					buyNotify = true
				} else {
					t.sellTradePlate.Remove(matchOrder, focusedAmount)
					sellNotify = true
				}
				break
			} else {
				//当前的订单 不满足交易额 继续进行匹配
				turnover := op.MulFloor(price, matchAmount, 8)
				matchOrder.TradedAmount = op.AddFloor(matchOrder.TradedAmount, matchAmount, 8)
				matchOrder.Turnover = op.AddFloor(matchOrder.Turnover, turnover, 8)
				matchOrder.Status = model.Completed
				// 添加到我们的完成队列中
				completeOrders = append(completeOrders, matchOrder)
				//从队列进行删除
				delOrders = append(delOrders, matchOrder.OrderId)

				focusedOrder.TradedAmount = op.AddFloor(focusedOrder.TradedAmount, matchAmount, 8)
				focusedOrder.Turnover = op.AddFloor(focusedOrder.Turnover, turnover, 8)

				if matchOrder.Direction == model.BUY {
					t.buyTradePlate.Remove(matchOrder, matchAmount)
					buyNotify = true
				} else {
					t.sellTradePlate.Remove(matchOrder, matchAmount)
					sellNotify = true
				}
				continue
			}
		}
	}
	//处理队列中 已经完成的订单进行删除
	for _, orderId := range delOrders {
		for _, v := range lpList.list {
			for index, matchOrder := range v.list {
				if orderId == matchOrder.OrderId {
					v.list = append(v.list[:index], v.list[index+1:]...)
					break
				}
			}
		}
	}
	// 因为还有下一轮交易判断，这里暂时不判断是否完成订单
	// 通知买卖盘更新
	if buyNotify {
		t.sendTradePlateMsg(t.buyTradePlate)
	}
	if sellNotify {
		t.sendTradePlateMsg(t.sellTradePlate)
	}
	for _, v := range completeOrders {
		t.sendCompleteOrder(v)
	}
}

// matchLimitPriceWithMP focusedOrder 限价单
func (t *CoinTrade) matchLimitPriceWithMP(mpList TradeTimeQueue, focusedOrder *model.ExchangeOrder) {
	var delOrders []string
	for _, matchOrder := range mpList {
		if matchOrder.MemberId == focusedOrder.MemberId {
			//自己的订单就不处理了
			continue
		}
		price := focusedOrder.Price
		//计算可交易的数量
		matchAmount := op.SubFloor(matchOrder.Amount, matchOrder.TradedAmount, 8)
		if matchAmount <= 0 {
			continue
		}
		focusedAmount := op.SubFloor(focusedOrder.Amount, focusedOrder.TradedAmount, 8)
		if matchAmount >= focusedAmount {
			//订单直接就交易完成了 能满足
			turnover := op.MulFloor(price, focusedAmount, 8)
			matchOrder.TradedAmount = op.AddFloor(matchOrder.TradedAmount, focusedAmount, 8)
			matchOrder.Turnover = op.AddFloor(matchOrder.Turnover, turnover, 8)
			if op.SubFloor(matchOrder.Amount, matchOrder.TradedAmount, 8) <= 0 {
				matchOrder.Status = model.Completed
				//从队列进行删除
				delOrders = append(delOrders, matchOrder.OrderId)
			}
			focusedOrder.TradedAmount = op.AddFloor(focusedOrder.TradedAmount, focusedAmount, 8)
			focusedOrder.Turnover = op.AddFloor(focusedOrder.Turnover, turnover, 8)
			focusedOrder.Status = model.Completed
			break
		} else {
			//当前的订单 不满足交易额 继续进行匹配
			turnover := op.MulFloor(price, matchAmount, 8)
			matchOrder.TradedAmount = op.AddFloor(matchOrder.TradedAmount, matchAmount, 8)
			matchOrder.Turnover = op.AddFloor(matchOrder.Turnover, turnover, 8)
			matchOrder.Status = model.Completed
			//从队列进行删除
			delOrders = append(delOrders, matchOrder.OrderId)
			focusedOrder.TradedAmount = op.AddFloor(focusedOrder.TradedAmount, matchAmount, 8)
			focusedOrder.Turnover = op.AddFloor(focusedOrder.Turnover, turnover, 8)
			continue
		}
	}
	//处理已经匹配完成的订单 从队列删除
	for _, orderId := range delOrders {
		for index, matchOrder := range mpList {
			if matchOrder.OrderId == orderId {
				mpList = append(mpList[:index], mpList[index+1:]...)
				break
			}
		}
	}
}

func (t *CoinTrade) addLimitQueue(order *model.ExchangeOrder) {
	if order.Type != model.LimitPrice {
		return
	}
	if order.Direction == model.BUY {
		t.buyLimitQueue.mux.Lock()
		//deal
		isPut := false
		for _, o := range t.buyLimitQueue.list {
			if o.price == order.Price {
				o.list = append(o.list, order)
				isPut = true
				break
			}
		}
		if !isPut {
			lpm := &LimitPriceMap{
				price: order.Price,
				list:  []*model.ExchangeOrder{order},
			}
			t.buyLimitQueue.list = append(t.buyLimitQueue.list, lpm)
		}
		t.buyTradePlate.Add(order)
		t.buyLimitQueue.mux.Unlock()
	} else if order.Direction == model.SELL {
		t.sellLimitQueue.mux.Lock()
		//deal
		isPut := false
		for _, o := range t.sellLimitQueue.list {
			if o.price == order.Price {
				o.list = append(o.list, order)
				isPut = true
				break
			}
		}
		if !isPut {
			lpm := &LimitPriceMap{
				price: order.Price,
				list:  []*model.ExchangeOrder{order},
			}
			t.sellLimitQueue.list = append(t.sellLimitQueue.list, lpm)
		}
		t.sellTradePlate.Add(order)
		t.sellLimitQueue.mux.Unlock()
	}
}

// sendCompleteOrder 发送完成订单的相关信息，更改我们的订单状态
func (t *CoinTrade) sendCompleteOrder(order *model.ExchangeOrder) {
	if order.Status != model.Completed {
		return
	}
	marshal, _ := json.Marshal(order)
	kafkaData := database.KafkaData{
		Topic: "exchange_order_complete",
		Key:   []byte(t.symbol),
		Data:  marshal,
	}
	// 我们的订单状态转换是必须执行成功的
	for {
		err := t.kafkaClient.SendSync(kafkaData)
		if err != nil {
			logx.Error(err)
			time.Sleep(250 * time.Millisecond)
			continue
		} else {
			break
		}
	}
}

func NewCoinTrade(symbol string, cli *database.KafkaClient, db *dbutils.ElysiaDB) *CoinTrade {
	c := &CoinTrade{
		symbol:      symbol,
		kafkaClient: cli,
		db:          db,
	}
	c.init()
	return c
}
