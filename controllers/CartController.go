package controllers

import (
	"github.com/astaxie/beego"
)

type CartController struct {
	beego.Controller
}

func (c *CartController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "cart.html"
}
