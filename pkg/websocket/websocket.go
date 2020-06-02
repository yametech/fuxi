package websocket

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Clients map[*websocket.Conn]struct{}

func NewClients() map[*websocket.Conn]struct{} {
	return make(map[*websocket.Conn]struct{})
}

func (clients Clients) Upgrade(g *gin.Context, rsize, wsize int, deadline time.Time) (*websocket.Conn, error) {
	fn := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  rsize,
		WriteBufferSize: wsize,
	}
	conn, err := fn.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		return nil, err
	}
	// Register our new clientv2
	clients[conn] = struct{}{}

	go func() {
		if err := conn.WriteControl(websocket.PingMessage, nil, deadline); err != nil {
			CloseTheWebsocket(conn)
			delete(clients, conn)
		}
	}()

	return conn, nil
}

func CloseTheWebsocket(conn *websocket.Conn) {
	_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	_ = conn.Close()
}
