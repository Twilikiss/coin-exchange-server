// Package code_gen
// @Author twilikiss 2024/5/11 13:08:08
package code_gen

import "testing"

func TestGenRpc(t *testing.T) {
	rpcCommon := RpcCommon{
		PackageName: "eclient",
		ModuleName:  "exchange",
		ServiceName: "Order",
		GrpcPackage: "order",
	}
	// rpc FindSymbolThumb(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolThumbTrend(MarketReq) returns(SymbolThumbRes);
	//  rpc FindSymbolInfo(MarketReq) returns(ExchangeCoin);
	rpc1 := Rpc{
		FunName: "FindOrderHistory",
		Resp:    "OrderRes",
		Req:     "OrderReq",
	}
	rpc2 := Rpc{
		FunName: "FindOrderCurrent",
		Resp:    "OrderRes",
		Req:     "OrderReq",
	}
	//rpc3 := Rpc{
	//	FunName: "FindSymbolInfo",
	//	Resp:    "ExchangeCoin",
	//	Req:     "MarketReq",
	//}
	var rpcList []Rpc
	rpcList = append(rpcList, rpc1, rpc2)
	result := RpcResult{
		RpcCommon: rpcCommon,
		Rpc:       rpcList,
	}
	GenZeroRpc(result)
}
