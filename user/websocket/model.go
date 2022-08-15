package websocket

import "time"

type MsgInfoReq struct {
	ConType    int    `json:"con_type" binding:"required"`
	SendFrom   string `json:"send_from" binding:"required"` //
	SendTo     string `json:"send_to"`                      //
	MsgContent string `json:"msg_content"`                  // 内容
}

type ContractChatMsgInfo struct {
	ConFromTo  string    `json:"con_from_to" binding:"required"` // 通话记录
	MsgContent string    `json:"msg_content" binding:"required"` // 内容
	MsgDate    time.Time `json:"msg_date"`                       // 创建时间
}
