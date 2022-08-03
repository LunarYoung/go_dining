package service

import (
	"user/model"
	"user/model/req"
	"user/repo"
)

type UserService interface {
	Reg(req req.Org)
}

type userService struct {
	repo repo.CategoryRepository
}

func (u userService) Reg(req req.Org) {
	var i = model.Org{
		Id:       req.Id,
		PassWord: req.PassWord,
	}
	u.repo.Create(i)
}

func NewUserService() UserService {
	return &userService{}
}
