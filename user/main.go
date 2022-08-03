package main

import (
	"user/pkg"
	"user/route"
)

func main() {
	pkg.Mysql()
	route.InitRouter()

}
