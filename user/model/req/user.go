package req

type Org struct {
	Id            int64  ` json:"id"`
	Name          string ` json:"name"`
	BusinessPhoto string ` json:"business_photo"`
	Sex           bool   ` json:"sex"`
	Phone         int    `binding:"required" json:"phone"`
	PassWord      string `binding:"required" json:"password"`
	Status        int8   `json:"status"`
}

type AppUserReq struct {
	Id            int64  `json:"id"`
	WxId          string `json:"wx_id"`
	Name          string `json:"name"`
	Address       string ` json:"address"`
	Avt           string ` json:"avt"`
	Sex           bool   ` json:"sex"`
	RegTime       string ` json:"reg_time"`
	LastLoginTime string ` json:"last_login_time"`
	Status        int8   ` json:"status"`
}
