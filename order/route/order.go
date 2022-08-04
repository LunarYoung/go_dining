package route

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func InitRouter() {
	router := gin.Default()

	v2 := router.Group("order")
	{
		c := controller.NewOrderController()
		v2.POST("/oss", c.Oss)

	}

	err := router.Run(":80")
	if err != nil {
		return
	}

}
