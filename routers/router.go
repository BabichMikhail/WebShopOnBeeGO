package routers

import (
    ctls "WebShopOnBeeGO/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/webshop", &ctls.WebShopController{}, "get,post:HomePage")
    beego.Router("/webshop/catalog", &ctls.WebShopController{}, "get,post:Catalog")
    beego.Router("/webshop/catalog/:equip([a-z]+)", &ctls.WebShopController{}, "get,post:Catalog")
    beego.Router("/webshop/catalog/:equip([a-z]+)/:type([a-z_]+)", &ctls.WebShopController{}, "get,post:Catalog")
    beego.Router("/webshop/catalog/:equip([a-z]+)/:type([a-z_]+)/:level([a-z_]+)", &ctls.WebShopController{}, "get,post:Catalog")
    beego.Router("/", &ctls.MainController{})
    beego.Router("/webshop/login", &ctls.LoginController{}, "get,post:Login")
    beego.Router("/webshop/logout", &ctls.LoginController{}, "get:Logout")
    beego.Router("/webshop/signup", &ctls.LoginController{}, "get,post:Signup")
    beego.Router("/webshop/card/add/:equip_id([0-9]+)", &ctls.AddProductController{})
    beego.Router("/webshop/purchases", &ctls.PurchaseController{})
    beego.Router("/webshop/goodsinfo/:equip_id([0-9]+)", &ctls.GoodsInfoController{})
    beego.Router("/webshop/purchases/checkout", &ctls.CheckoutController{})
    beego.Router("/webshop/admin", &ctls.AdminController{})
}
