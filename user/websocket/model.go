package websocket

type MsgInfoReq struct {
	OrgId      string `json:"org_id" binding:"required"`
	ConType    int    `json:"con_type" binding:"required"`
	SendFrom   string `json:"send_from" binding:"required"` //
	SendTo     string `json:"send_to"`                      //
	MsgContent string `json:"msg_content"`                  // 内容
}

type MsgFirstReq struct {
	OrgId    string `json:"org_id" binding:"required"`
	SendFrom string `json:"send_from" binding:"required"` //
	SendTo   string `json:"send_to" binding:"required"`   //
	Size     int64  `json:"size"`
}

type ContractChatMsgInfo struct {
	OrgId      string `json:"org_id"`
	ConType    int    `json:"con_type"`
	SendFrom   string `json:"send_from" `   //
	SendTo     string `json:"send_to"`      //
	MsgContent string `json:"msg_content" ` // 内容
	MsgDate    string `json:"msg_date"`     // 创建时间
}
