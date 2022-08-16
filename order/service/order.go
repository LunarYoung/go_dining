package service

import (
	"order/middleware"
	"order/model"
	"order/model/req"
	"order/pkg"
)

type OrderService interface {
	SaveOrder(m model.Order, i string)
	SearchOrder(r req.OrderSearchReq, i string) (re []model.Order, count int64)
	ChangeOrder(r req.OrderSearchReq, i string)
}

type orderService struct {
}

func (u orderService) ChangeOrder(r req.OrderSearchReq, i string) {
	middleware.Update(r, i)
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (u orderService) SaveOrder(m model.Order, i string) {
	for {
		id := pkg.Uuid()
		if middleware.GetStr(id) == "" {
			m.PickUp = id
			err := middleware.SetStr(id, "pickup num")
			if err != nil {
				return
			}
			break
		}
	}

	middleware.Create(m, i, m.OrderId)
}

func (u orderService) SearchOrder(r req.OrderSearchReq, i string) (re []model.Order, count int64) {
	return middleware.Query(r, i)
}
