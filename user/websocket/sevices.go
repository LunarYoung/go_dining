package websocket

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"user/middleware"
)

func SaveMongo(req MsgInfoReq) {
	var rep ContractChatMsgInfo
	rep.OrgId = req.OrgId
	rep.ConType = req.ConType
	rep.MsgContent = req.MsgContent
	rep.MsgDate = time.Now().Format("2006-01-02 15:04:05")
	rep.SendTo = req.SendTo
	rep.SendFrom = req.SendFrom
	middleware.AddOne(rep, "msg")
}

func QueryMongo(req MsgFirstReq) (rep []middleware.ContractChatMsgInfo) {
	fmt.Println(req)
	m := bson.M{"$or": []bson.M{{"sendfrom": req.SendFrom, "sendto": req.SendTo, "org_id": req.OrgId}, {"sendto": req.SendFrom, "sendfrom": req.SendTo, "org_id": req.OrgId}}}
	rep = middleware.GetList(m, req.Size)
	return rep
}
