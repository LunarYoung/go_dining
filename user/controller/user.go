package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"user/middleware"
	"user/model/rep"
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

func (u UserController) Test(g *gin.Context) {
	g.JSON(200, "测试token")
}

// Create
// @description: 管理端注册
// @param g
// @2022-08-16 15:03:44
func (u UserController) Create(g *gin.Context) {

	var r req.Org
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	u.s.Reg(r)
	g.JSON(200, rep.NewBSSRep())
}

// AppCreate
// @description: app用户注册
// @param g
// @2022-08-16 15:05:03

func (u UserController) AppCreate(g *gin.Context) {
	var r req.AppUserReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	u.s.AppCreate(r)
}
func (u UserController) Login(g *gin.Context) {
	var r req.Org
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	var flag = u.s.Login(r)
	if flag {
		//返回token
		var myClaims = middleware.MyClaims{
			Phone: r.Phone,
		}
		var tokenUtil = middleware.NewJWT()
		var token, _ = tokenUtil.CreateToken(myClaims)
		g.JSON(200, rep.Token{
			Code:  200,
			Token: token,
		})
		//存到redis，自动过期
		err := middleware.SetStr(token, strconv.Itoa(r.Phone))
		if err != nil {
			return
		}
	} else {
		var r rep.BaseRep
		r.Msg = "密码账号错误"
		r.Code = 500
		g.JSON(500, r)
	}
}
