package routers

import (
    ctls "WebShopOnBeeGO/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/webshop", &ctls.WebShopController{}, "get,post:HomePage")
    beego.Router("/webshop/catalog/*", &ctls.WebShopController{}, "get,post:Catalog")
    beego.Router("/", &ctls.MainController{})
    beego.Router("/webshop/login", &ctls.LoginController{}, "get,post:Login")
    beego.Router("/webshop/logout", &ctls.LoginController{}, "get:Logout")
    beego.Router("/webshop/signup", &ctls.LoginController{}, "get,post:Signup")
}
