package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

func Send(g *gin.Context) {

	var req MsgInfoReq
	if err := g.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		return
	}

	msg, _ := json.Marshal(req.MsgContent)
	var t = Manager.Clients[req.SendTo]
	err := t.Socket.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}
