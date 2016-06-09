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
    beego.Router("/webshop/purchases/change/:equip_id([0-9]+)/:add([0-1])", &ctls.PurchaseController{}, "get,post:Change")
    beego.Router("/webshop/purchases/delete/:equip_id([0-9]+)", &ctls.PurchaseController{}, "get,post:Delete")
    beego.Router("/webshop/goodsinfo/:equip_id([0-9]+)", &ctls.GoodsInfoController{})
    beego.Router("/webshop/purchases/checkout", &ctls.CheckoutController{})
    beego.Router("/webshop/admin", &ctls.AdminController{}, "get:AdminHome")
    beego.Router("/webshop/admin/purchases", &ctls.AdminController{})
    beego.Router("/webshop/admin/editcard/:tableName([a-z]+)/:equip_id([0-9]+)", &ctls.AdminController{}, "post:SaveGoodsInfo")
    beego.Router("/webshop/admin/edittable/:tableName([a-z]+)", &ctls.AdminController{}, "get:EditTable")
    beego.Router("/webshop/admin/edittable/:tableName([a-z]+)", &ctls.AdminController{}, "post:ApplyChanges")
    beego.Router("/webshop/user", &ctls.UserController{})
}
