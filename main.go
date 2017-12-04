package main

import (
	"github.com/astaxie/beego"
	_ "wr_web/routers"
	"wr_web/services"
)

func main() {
	go services.Task()
	beego.Run()
}
