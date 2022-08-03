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
	if err := c.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	g.s.Reg(r)
}
func (g UserController) Login(c *gin.Context) {

	var r req.Org
	if err := c.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	var flag = g.s.Login(r)
	if flag {
		//返回token
	}
}
