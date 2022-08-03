package req

type Org struct {
	Id            int64  `binding:"required" json:"id"`
	Name          string ` json:"name"`
	BusinessPhoto string ` json:"business_photo"`
	Sex           bool   ` json:"sex"`
	Phone         int    ` json:"phone"`
	PassWord      string `binding:"required" json:"password"`
	Status        int8   `json:"status"`
}
