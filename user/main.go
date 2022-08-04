package main

import (
	"user/pkg"
	"user/route"
)

func main() {
	pkg.Mysql()
	pkg.ViperInit()
	route.InitRouter()
}
