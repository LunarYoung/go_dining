package service

import (
	"order/pkg"
)

type OrderService interface {
	UpFood()
}

type orderService struct {
	redis pkg.RedisUtil
}

func NewOrderService() OrderService {
	return &orderService{
		redis: pkg.NewRedisClient(),
	}
}

func (u orderService) UpFood() {

}
