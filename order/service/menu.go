package service

import (
	"order/model"
	"order/model/req"
	"order/pkg"
)

type MenuService interface {
	SaveMenu(m model.Menu, i string)
	SearchMenu(r req.MenuSearchReq, i string) (re []model.Menu, count int64)
	ChangeMenu(r req.MenuChangeReq, i string)
}

type menuService struct {
	//redis pkg.RedisUtil
}

func (u menuService) ChangeMenu(r req.MenuChangeReq, i string) {
	//TODO implement me
	pkg.UpdateMenu(r, i)
}

func NewMenuService() MenuService {
	return &menuService{
		//redis: pkg.NewRedisClient(),
	}
}

func (u menuService) SaveMenu(m model.Menu, i string) {

	pkg.Create(m, i, m.Uuid)
}

func (u menuService) SearchMenu(r req.MenuSearchReq, i string) (re []model.Menu, count int64) {
	return pkg.QueryMenu(r, i)
}
