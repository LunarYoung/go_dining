package main

import (
	"user/pkg"
	"user/route"
)

func main() {
	//pkg.Mysql()
	//pkg.ViperInit()
	pkg.Nacos()
	pkg.ConnectToDB()
	route.InitRouter()
	//service.KafkaSend()
	//client.C()

}
