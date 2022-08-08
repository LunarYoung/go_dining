package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"order/model"
	"order/model/rep"
	"order/model/req"
	"order/pkg"
	"order/service"
)

//es的index
const order = "order"
const menu = "menu"

//status 含义
//	1 下单
//  2 商家在处理，无法取消
//  3 商家取消
//  4 用户取消
//  5 商家完成，待取货
//  6 订单完成

type OrderController struct {
	s service.OrderService
}

func NewOrderController() OrderController {
	return OrderController{
		s: service.NewOrderService(),
	}
}

// Oss
// @description: 上传图片，转url
// @param g
// @2022-08-06 09:52:14
func (u OrderController) Oss(g *gin.Context) {
	file, err := g.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileHandle, err := file.Open() //打开上传文件
	if err != nil {
		panic(err)
	}
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		if err != nil {
			panic(err)
		}
	}(fileHandle)
	fileByte, err := ioutil.ReadAll(fileHandle) //获取上传文件字节流
	if err != nil {
		fmt.Println(err)
	}
	id := uuid.NewV4()
	ids := id.String()
	url, err := pkg.Upload(ids+".png", fileByte)
	fmt.Println(err)
	g.JSON(http.StatusOK, gin.H{
		"dataType": "Image",
		"data": gin.H{
			"url": url,
		},
	})
}

// SaveOrder
// @description: 保存订单
// @param g
// @2022-08-06 10:41:47
func (u OrderController) SaveOrder(g *gin.Context) {
	var r req.OrderReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		g.JSON(200, rep.NewBSEJRep())
	}
	var m model.Order
	err := copier.Copy(&m, &r)
	if err != nil {
		return
	}
	m.OrderId = uuid.NewV4().String()

	u.s.SaveOrder(m, order)

	g.JSON(200, rep.NewBSSRep())

}
func (u OrderController) SearchOrder(g *gin.Context) {
	var r req.OrderSearchReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		g.JSON(200, rep.NewBSEJRep())
	}
	//var rep rep.Rep
	var count int64
	Rep := rep.NewRep()
	Rep.Date.Content, count = u.s.SearchOrder(r, order)
	Rep.Date.PageRep = pkg.ToPage(r.Page.PageSize, r.Page.PageIndex, count)
	g.JSON(200, Rep)

}

func (u OrderController) ChangeOrder(g *gin.Context) {
	var r req.OrderSearchReq
	if err := g.ShouldBind(&r); err != nil {
		log.Println(err.Error())
		g.JSON(200, rep.NewBSEJRep())
	}
	u.s.ChangeOrder(r, order)
	g.JSON(200, rep.NewBSSRep())

}
