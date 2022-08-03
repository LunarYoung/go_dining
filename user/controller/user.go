package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"user/model/req"
	"user/service"
)

type UserController struct {
	s service.UserService
}

func NewUserController() UserController {
	return UserController{
		s: service.NewUserService(),
	}
}

func (g UserController) Create(c *gin.Context) {
	var r req.Org
	// 数据验证
	if err := c.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	g.s.Reg(r)
}
