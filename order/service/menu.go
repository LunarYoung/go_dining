package service

import (
	"order/middleware"
	"order/model"
	"order/model/req"
)

type MenuService interface {
	SaveMenu(m model.Menu, i string)
	SearchMenu(r req.MenuSearchReq, i string) (re []model.Menu, count int64)
	ChangeMenu(r req.MenuChangeReq, i string)
}

type menuService struct {
	//redis middleware.RedisUtil
}

func (u menuService) ChangeMenu(r req.MenuChangeReq, i string) {
	//TODO implement me
	middleware.UpdateMenu(r, i)
}

func NewMenuService() MenuService {
	return &menuService{
		//redis: middleware.NewRedisClient(),
	}
}

func (u menuService) SaveMenu(m model.Menu, i string) {

	middleware.Create(m, i, m.Uuid)
}

func (u menuService) SearchMenu(r req.MenuSearchReq, i string) (re []model.Menu, count int64) {
	return middleware.QueryMenu(r, i)
}
