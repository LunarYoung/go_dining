package main

import (
	"order/pkg"
	"order/route"
)

func main() {

	pkg.ViperInit()
	pkg.InitEs()
	pkg.NewOss()
	route.InitRouter()
	//service.KafkaRe()
	//pkg.ConnectToDB()
	//var a = model.Food{
	//	Name:  "dfsd222f",
	//	Price: "dfds222f",
	//}
	//var b = model.Order{Name: "dd222d", Id: 12}
	//pkg.AddOne(a, "order")
	//pkg.AddOne(b, "food")
	//pkg.Create()
	//delete()
	//pkg.Update()
	//gets()
	//pkg.Query()
	//pkg.List(2, 1)
}
