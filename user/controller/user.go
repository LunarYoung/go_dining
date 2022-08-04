package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"user/model/rep"
	"user/model/req"
	"user/pkg"
	"user/service"
)

type UserController struct {
	s     service.UserService
	redis pkg.RedisUtil
}

func NewUserController() UserController {
	return UserController{
		s:     service.NewUserService(),
		redis: pkg.NewRedisClient(),
	}
}

func (u UserController) Test(g *gin.Context) {
	g.JSON(200, "测试token")
}

func (u UserController) Create(g *gin.Context) {

	var r req.Org
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		return
	}
	u.s.Reg(r)
	g.JSON(200, rep.NewBSSRep())
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
		var myClaims = pkg.MyClaims{
			Phone: r.Phone,
		}
		var tokenUtil = pkg.NewJWT()
		var token, _ = tokenUtil.CreateToken(myClaims)
		g.JSON(200, rep.Token{
			Code:  200,
			Token: token,
		})
		//存到redis，自动过期
		err := u.redis.SetStr(token, strconv.Itoa(r.Phone))
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
