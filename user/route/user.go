package route

import (
	"github.com/gin-gonic/gin"
	"user/controller"
	"user/pkg"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」
	router.Use(pkg.Cors())

	//pc端口
	v1 := router.Group("dining")
	{
		c := controller.NewUserController()
		v1.POST("/reg", c.Create)
		v1.POST("/login", c.Create)

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
