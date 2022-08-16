package route

import (
	"github.com/gin-gonic/gin"
	"user/controller"
	"user/pkg"
	"user/websocket"
)

func InitRouter() {
	router := gin.Default()

	router.Use(pkg.Cors())
	//router.Use(pkg.Logger())

	//pc端口
	v1 := router.Group("base")
	{
		c := controller.NewUserController()
		v1.POST("/reg", c.Create)
		v1.POST("/app/reg", c.AppCreate)
		v1.POST("/login", c.Login)
		v1.GET("/ws", websocket.WsHandler)
		v1.POST("/ws/query", websocket.Query)
		v1.POST("/ws/send", websocket.Send)

	}

	router.Use(pkg.JwtToken())

	v2 := router.Group("dining")
	{
		c := controller.NewUserController()
		v2.GET("/test", c.Test)

	}

	err := router.Run(":8089")
	if err != nil {
		return
	}

}
