package model

type AppUser struct {
	Id            int64  `gorm:"column:id" db:"column:id" json:"id"`
	WxId          string `gorm:"column:wx_id" db:"column:wx_id" json:"wx_id"`
	Name          string `gorm:"column:name" db:"column:name" json:"name"`
	Address       string `gorm:"column:address" db:"column:address" json:"address"`
	Avt           string `gorm:"column:avt" db:"column:avt" json:"avt"`
	Sex           bool   `gorm:"column:sex" db:"column:sex" json:"sex"`
	RegTime       string `gorm:"column:reg_time" db:"column:reg_time" json:"reg_time"`
	LastLoginTime string `gorm:"column:last_login_time" db:"column:last_login_time" json:"last_login_time"`
	Status        int8   `gorm:"column:status" db:"column:status" json:"status"`
	SocketId      string `gorm:"column:socket_id;Index:socket_index" db:"column:socket_id" json:"socket_id"`
}

type OrgUser struct {
	Id            int64  `gorm:"column:id" db:"column:id" json:"id"`
	Name          string `gorm:"column:name" db:"column:name" json:"name"`
	PassWord      string `gorm:"column:password" db:"column:password" json:"password"`
	Avt           string `gorm:"column:avt" db:"column:avt" json:"avt"`
	Sex           bool   `gorm:"column:sex" db:"column:sex" json:"sex"`
	RegTime       string `gorm:"column:reg_time" db:"column:reg_time" json:"reg_time"`
	LastLoginTime string `gorm:"column:last_login_time" db:"column:last_login_time" json:"last_login_time"`
	Status        int8   `gorm:"column:status" db:"column:status" json:"status"`
	SocketId      string `gorm:"column:socket_id;Index:socket_index" db:"column:socket_id" json:"socket_id"`
}

type Org struct {
	Id            int64  `gorm:"column:id" db:"column:id" json:"id"`
	Name          string `gorm:"column:name" db:"column:name" json:"name"`
	BusinessPhoto string `gorm:"column:business_photo" db:"column:business_photo" json:"business_photo"`
	Sex           bool   `gorm:"column:sex" db:"column:sex" json:"sex"`
	Phone         int    `gorm:"column:phone" db:"column:phone" json:"phone"`
	PassWord      string `gorm:"column:password" db:"column:password" json:"password"`
	Status        int8   `gorm:"column:status" db:"column:status" json:"status"`
}
