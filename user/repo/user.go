package repo

import (
	_ "gorm.io/gorm"
	"user/model"
	"user/pkg"
)

type UserRepository interface {
	Create(req model.Org)
	Login(req model.Org) string
}

type userRepository struct {
}

func NewUserService() UserRepository {
	return &userRepository{}
}
func (c userRepository) Login(req model.Org) string {
	var rep model.Org
	pkg.Db.Where("phone =?", req.Phone).Find(rep)
	return rep.PassWord
}

func (c userRepository) Create(req model.Org) {
	pkg.Db.Create(&req)
}
