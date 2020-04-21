package routers

import (
	"fresh_every_day/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Post")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/order", &controllers.OrderController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/cart", &controllers.CartController{})

}
