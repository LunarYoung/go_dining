package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"net/http"
	"time"
)

type Client struct {
	Id             string
	LastOnlineTime time.Time
	Socket         *websocket.Conn
	Send           chan []byte
}

type ClientManager struct {
	Clients map[string]*Client
}

var Manager = ClientManager{
	Clients: make(map[string]*Client),
}

func WsHandler(g *gin.Context) {
	connectId := g.Query("connect_id")
	fmt.Println("qwe")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(g.Writer, g.Request, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(g.Writer, g.Request)
		return
	}

	client := &Client{
		Id:             connectId,
		LastOnlineTime: time.Now(),
		Socket:         conn,
		Send:           make(chan []byte),
	}

	Manager.Clients[connectId] = client

	go client.Read()

}
func (c *Client) Read() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
	}

}
