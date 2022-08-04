package main

import (
	"order/pkg"
	"order/route"
)

func main() {

	pkg.ViperInit()
	pkg.NewOss()
	route.InitRouter()

	//pkg.Create()
	//delete()
	//pkg.Update()
	//gets()
	//pkg.Query()
	//pkg.List(2, 1)
}
