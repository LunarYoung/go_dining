package service

import (
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
	redis pkg.RedisUtil
}

func (u orderService) ChangeOrder(r req.OrderSearchReq, i string) {
	pkg.Update(r, i)
}

func NewOrderService() OrderService {
	return &orderService{
		redis: pkg.NewRedisClient(),
	}
}

func (u orderService) SaveOrder(m model.Order, i string) {
	for {
		id := pkg.Uuid()
		if u.redis.GetStr(id) == "" {
			m.PickUp = id
			err := u.redis.SetStr(id, "pickup num")
			if err != nil {
				return
			}
			break
		}
	}

	pkg.Create(m, i, m.OrderId)
}

func (u orderService) SearchOrder(r req.OrderSearchReq, i string) (re []model.Order, count int64) {
	return pkg.Query(r, i)
}
