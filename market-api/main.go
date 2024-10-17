package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/chain"
	"market-api/internal/config"
	"market-api/internal/handler"
	"market-api/internal/svc"
	"market-api/internal/ws"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	// 为了调试方便，我们先将log日志设置一下
	logx.MustSetup(logx.LogConf{
		Stat:     false,
		Encoding: "plain",
	})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	wsServer := ws.NewWebsocketServer("/socket.io")
	server := rest.MustNewServer(
		c.RestConf,
		rest.WithChain(chain.New(wsServer.ServerHandler)),
		rest.WithCustomCors(func(header http.Header) {
			header.Set(
				"Access-Control-Allow-Headers",
				"DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token,x-auth-token")
		}, nil, "http://localhost:8080"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c, wsServer)

	router := handler.NewRouters(server)

	// 这里我们将server套多一层Routers
	handler.ExchangeRateHandlers(router, ctx)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(wsServer) // 如果我们希望添加我们自定义的server就需要包含Start()和Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	group.Start()
}
