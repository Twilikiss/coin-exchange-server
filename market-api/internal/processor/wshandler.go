// Package processor
// @Author twilikiss 2024/5/8 22:23:23
package processor

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"market-api/internal/model"
	"market-api/internal/ws"
)

type WebSocketHandler struct {
	wsServer *ws.WebsocketServer
}

func (w *WebSocketHandler) HandleTradePlate(symbol string, plate *model.TradePlateResult) {
	bytes, _ := json.Marshal(plate)
	logx.Info("====买卖盘通知:", symbol, plate.Direction, ":", fmt.Sprintf("%d", len(plate.Items)))
	w.wsServer.BroadcastToNamespace("/", "/topic/market/trade-plate/"+symbol, string(bytes))
}

func (w *WebSocketHandler) HandleTrade(symbol string, data []byte) {
	// 订单交易完成后 进入这里进行处理 订单就称为K线的一部分 数据量小 无法维持K线 所以我们的K线来源于okx平台
	//TODO implement me
	panic("implement me")
}

func (w *WebSocketHandler) HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*model.CoinThumb) {
	logx.Info("================WebsocketHandler Start=======================")
	logx.Info("symbol:", symbol)
	logx.Info("close:", kline.ClosePrice)
	logx.Info("high:", kline.HighestPrice)
	logx.Info("time:", kline.Time)
	// 在向我们的websocket发送前先处理一下我们的kline数据
	thumb := thumbMap[symbol]
	if thumb == nil {
		thumb = kline.InitCoinThumb(symbol)
	}
	coinThumb := kline.ToCoinThumb(symbol, thumb)
	finalData, err := json.Marshal(coinThumb)
	if err != nil {
		logx.Error("json转换失败，err=", err)
	}
	w.wsServer.BroadcastToNamespace("/", "/topic/market/thumb", string(finalData))

	// 同样借助WebSocket把我们Kline实时变动传递到前端
	bytes, _ := json.Marshal(kline)
	w.wsServer.BroadcastToNamespace("/", "/topic/market/kline/"+symbol, string(bytes))
	logx.Info("================WebsocketHandler End=========================")
}

func NewWebSocketHandler(wsServer *ws.WebsocketServer) *WebSocketHandler {
	return &WebSocketHandler{
		wsServer: wsServer,
	}
}
