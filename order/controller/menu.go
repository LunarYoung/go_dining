package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"log"
	"order/model"
	"order/model/rep"
	"order/model/req"
	"order/pkg"
	"order/service"
)

//es的index
const Menu = "menu"

//status 含义
//	1 下单
//  2 商家在处理，无法取消
//  3 商家取消
//  4 用户取消
//  5 商家完成，待取货
//  6 订单完成

type MenuController struct {
	s service.MenuService
}

func NewMenuController() MenuController {
	return MenuController{
		s: service.NewMenuService(),
	}
}

// Oss
// @description: 上传图片，转url
// @param g
// @2022-08-06 09:52:14

// SaveMenu
// @description: 保存订单
// @param g
// @2022-08-06 10:41:47
func (u MenuController) SaveMenu(g *gin.Context) {
	var r req.MenuReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		g.JSON(200, rep.NewBSEJRep())
		return
	}
	var m model.Menu
	err := copier.Copy(&m, &r)
	if err != nil {
		return
	}
	m.Uuid = uuid.NewV4().String()

	u.s.SaveMenu(m, Menu)

	g.JSON(200, rep.NewBSSRep())

}

func (u MenuController) SearchMenu(g *gin.Context) {
	var r req.MenuSearchReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		g.JSON(500, rep.NewBSEJRep())
		return
	}

	var count int64
	Rep := rep.NewRep()
	Rep.Date.Content, count = u.s.SearchMenu(r, Menu)
	Rep.Date.PageRep = pkg.ToPage(r.Page.PageSize, r.Page.PageIndex, count)
	g.JSON(200, Rep)

}

func (u MenuController) ChangeMenu(g *gin.Context) {
	var r req.MenuChangeReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		g.JSON(200, rep.NewBSEJRep())
		return
	}
	u.s.ChangeMenu(r, Menu)
	g.JSON(200, rep.NewBSSRep())

}
