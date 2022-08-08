package route

import (
	"github.com/gin-gonic/gin"
	"order/controller"
)

func InitRouter() {
	router := gin.Default()

	v1 := router.Group("order")
	{
		c := controller.NewOrderController()
		v1.POST("/oss", c.Oss)
		v1.POST("/save/order", c.SaveOrder)
		v1.POST("/search/order", c.SearchOrder)
		v1.POST("/change/order", c.ChangeOrder)

	}

	v2 := router.Group("menu")
	{
		c := controller.NewMenuController()

		v2.POST("/save/menu", c.SaveMenu)
		v2.POST("/search/menu", c.SearchMenu)
		v2.POST("/change/menu", c.ChangeMenu)

	}

	err := router.Run(":80")
	if err != nil {
		return
	}

}
