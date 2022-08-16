package main

import (
	"user/middleware"
	"user/route"
)

func main() {

	middleware.Nacos()
	middleware.Mysql()
	middleware.InitRedis()
	middleware.ConnectToDB()
	route.InitRouter()
	//service.KafkaSend()
	//client.C()

}
