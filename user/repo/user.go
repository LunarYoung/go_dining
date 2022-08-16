package repo

import (
	uuid "github.com/satori/go.uuid"
	_ "gorm.io/gorm"
	"time"
	"user/middleware"
	"user/model"
)

type UserRepository interface {
	Create(req model.Org)
	AppCreate(req model.AppUser)
	Login(req model.Org) string
}

type userRepository struct {
}

func (c userRepository) AppCreate(req model.AppUser) {
	//第一次注册，有微信id就create，否则update
	if req.WxId != "" {
		req.SocketId = uuid.NewV4().String()
		req.RegTime = time.Now().Format("2006-01-02 15:04:05")
		middleware.Db.Create(&req)
	} else {
		middleware.Db.Model(&model.AppUser{}).Where("id", req.Id).Updates(req)
	}
}

func NewUserService() UserRepository {
	return &userRepository{}
}
func (c userRepository) Login(req model.Org) string {
	var rep model.Org
	middleware.Db.Where("phone =?", req.Phone).Find(&rep)
	return rep.PassWord
}

func (c userRepository) Create(req model.Org) {
	middleware.Db.Create(&req)
}
