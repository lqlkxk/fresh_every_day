package controllers

import (
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

func (c *OrderController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user_center_order.html"
}
