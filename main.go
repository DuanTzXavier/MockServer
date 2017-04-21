package main

import (
	_ "MockServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

