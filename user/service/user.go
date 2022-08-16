package service

import (
	"github.com/jinzhu/copier"
	"time"
	"user/model"
	"user/model/req"
	"user/pkg"
	"user/repo"
)

type UserService interface {
	Reg(req req.Org)
	AppCreate(req req.AppUserReq)
	Login(req req.Org) (flag bool)
}

type userService struct {
	repo repo.UserRepository
}

func (u userService) AppCreate(req req.AppUserReq) {
	var m model.AppUser
	err := copier.Copy(&m, &req)
	if err != nil {
		panic(err)
	}
	m.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")
	u.repo.AppCreate(m)
}

func NewUserService() UserService {
	return &userService{
		repo: repo.NewUserService(),
	}
}

// Login
// @description: 登陆
// @param req
// @return flag
// @2022-08-06 09:35:28
func (u userService) Login(req req.Org) (flag bool) {
	var i = model.Org{
		Phone:    req.Phone,
		PassWord: pkg.HashEncode(req.PassWord),
	}
	//登陆查询数据库
	var p = u.repo.Login(i)
	if p == "" {
		return false
	}
	if pkg.ComparePasswords(p, req.PassWord) {
		return true
	}
	return false

}

// Reg
// @description: 注册用户
// @param req
// @2022-08-06 09:35:44
func (u userService) Reg(req req.Org) {
	var i = model.Org{
		Phone:    req.Phone,
		PassWord: pkg.HashEncode(req.PassWord),
	}
	u.repo.Create(i)
}
