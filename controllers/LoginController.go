package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}
func (c *LoginController) Post() { //Form
	username := c.GetString("username")
	password := c.GetString("password")
	logs.Info("[username]:", username)
	logs.Info("[password]:", password)
	if username == "" || password == "" {
		c.Data["err"] = "賬號或密碼錯誤"
		c.TplName = "login.html"
		return
	}
	c.TplName = "user_center_info.html"
}

func (c *LoginController) PostJson() { // json接收
	username := c.GetString("username")
	password := c.GetString("password")
	logs.Info("[username]:", username)
	logs.Info("[password]:", password)
	if username == "" || password == "" {
		c.Data["err"] = "賬號或密碼錯誤"
		c.TplName = "login.html"
		return
	}
	c.TplName = "user_center_info.html"
}
