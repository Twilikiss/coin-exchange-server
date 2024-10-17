// Package ws
// @Author twilikiss 2024/5/9 15:34:34
package ws

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

const ROOM = "market"

type WebsocketServer struct {
	path   string
	server *socketio.Server
}

func (ws *WebsocketServer) Start() {
	_ = ws.server.Serve()
}

func (ws *WebsocketServer) Stop() {
	_ = ws.server.Close()
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func NewWebsocketServer(path string) *WebsocketServer {
	//解决跨域
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		logx.Info("connected:", s.ID())
		s.Join(ROOM) // 为前端的连接配置加入的room
		return nil
	})

	return &WebsocketServer{
		path:   path,
		server: server,
	}
}

// BroadcastToNamespace "/" 前端通过event="/topic/market/thumb"
func (ws *WebsocketServer) BroadcastToNamespace(path string, event string, data any) {
	go func() {
		ws.server.BroadcastToRoom(path, ROOM, event, data)
	}()
}

func (ws *WebsocketServer) ServerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 所有的请求都会来到这里
		path := request.URL.Path

		logx.Info("消息传递到了ServerHandler，path=", path)

		if strings.HasPrefix(path, ws.path) {
			// 关于websocket交由我们自己处理
			logx.Info("消息传递到了websocket处理，path=", path)
			ws.server.ServeHTTP(writer, request)
		} else {
			logx.Info("消息传递到了go-zero框架处理，path=", path)
			// 其余的交给go-zero自带的handler处理
			next.ServeHTTP(writer, request)
		}
	})
}
