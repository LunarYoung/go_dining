package websocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"user/model/rep"
)

func Send(g *gin.Context) {

	var req MsgInfoReq
	if err := g.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		return
	}
	msg, _ := json.Marshal(req.MsgContent)
	var t = Manager.Clients[req.SendTo]
	t.SendChan <- msg
	SaveMongo(req)
	var r = rep.BaseRep{Code: 200}
	g.JSON(200, r)
}

func Query(g *gin.Context) {
	var req MsgFirstReq
	if err := g.ShouldBind(&req); err != nil {
		log.Println(err.Error())
		return
	}

	var r = rep.NewNoPageRep()
	r.Date = QueryMongo(req)
	g.JSON(200, r)
}
