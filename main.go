package main

import (
	_ "fresh_every_day/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/images", "static/img")
	beego.Run()
}
