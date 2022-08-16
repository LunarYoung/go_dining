package main

import (
	"user/pkg"
	"user/route"
)

func main() {

	pkg.Nacos()
	pkg.Mysql()
	pkg.InitRedis()
	pkg.ConnectToDB()
	route.InitRouter()
	//service.KafkaSend()
	//client.C()

}
