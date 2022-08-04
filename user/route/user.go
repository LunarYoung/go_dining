package route

import (
	"github.com/gin-gonic/gin"
	"user/controller"
	"user/pkg"
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
		v1.POST("/login", c.Login)

	}

	router.Use(pkg.JwtToken())

	v2 := router.Group("dining")
	{
		c := controller.NewUserController()
		v2.GET("/test", c.Test)

	}
	//小程序端口
	//v2 := router.Group("Applets")
	//{
	//	//v2.GET("/code/image", Common.Code)
	//	v2.POST("/query", Controller.Query)
	//	v2.POST("/myOrder", Controller.MyOrder)
	//	v2.POST("/cfMyOrder", Controller.CfMyOrder)
	//	v2.GET("/showFood", Controller.ShowFood)
	//	v2.POST("/addAdvise", Controller.AddAdvise)
	//	v2.POST("/phone", Controller.Link)
	//	v2.POST("/address", Controller.Address)
	//	v2.POST("/login", Controller.AppLogin)
	//	v2.POST("/handleOrder", Controller.HandleOrder)
	//}

	//设置外部访问静态资源
	//router.StaticFS("/imageAssets", http.Dir("./imageAssets"))

	err := router.Run(":80")
	if err != nil {
		return
	}

}
