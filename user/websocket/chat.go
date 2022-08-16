package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"user/middleware"
)

const OnlineMap = "onlineMap"

type Client struct {
	Id             string
	LastOnlineTime time.Time
	Socket         *websocket.Conn
	SendChan       chan []byte
}

type ClientManager struct {
	Clients map[string]*Client
}

var Manager = ClientManager{
	Clients: make(map[string]*Client),
}

func WsHandler(g *gin.Context) {
	connectId := g.Query("connect_id")
	err := middleware.SetMap(OnlineMap, connectId, "true")
	if err != nil {
		return
	}

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
		SendChan:       make(chan []byte, 10),
	}

	Manager.Clients[connectId] = client
	go client.Read(client.SendChan)

}
func (c *Client) Read(ch chan []byte) {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		err := c.Socket.WriteMessage(websocket.TextMessage, <-ch)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.Socket.PongHandler()
	}

}
