package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"order/pkg"
	"order/service"
)

type OrderController struct {
	s service.OrderService
}

func NewOrderController() OrderController {
	return OrderController{
		s: service.NewOrderService(),
	}
}

func (u OrderController) Oss(g *gin.Context) {

	var (
		err error
	)

	file, err := g.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}

	fileHandle, err := file.Open() //打开上传文件
	if err != nil {
		fmt.Println(err)
	}
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		if err != nil {

		}
	}(fileHandle)
	fileByte, err := ioutil.ReadAll(fileHandle) //获取上传文件字节流
	if err != nil {
		fmt.Println(err)
	}
	url, err := pkg.Upload(file.Filename, fileByte)
	fmt.Println(err)
	g.JSON(http.StatusOK, gin.H{
		"error":    "",
		"errno":    "0",
		"dataType": "OBJECT",
		"data": gin.H{
			"url": url,
		},
	})

}
