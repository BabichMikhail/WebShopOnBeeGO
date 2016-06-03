package controllers

import (
    "github.com/astaxie/beego"
    "WebShopOnBeeGO/catalogs"
    "github.com/astaxie/beego/orm"
)

type BaseController struct {
    beego.Controller
}

func (c *BaseController) SetAuthorized() {
    authorized := c.GetSession("authorized")
    if authorized != nil && authorized.(bool) {
        c.Data["Authorized"] = true
        c.Data["Username"] = c.GetSession("username")
        c.Data["IsAdmin"] = c.GetSession("userright").(int) == 0
    } else {
        c.Data["Authorized"] = false
    }
}

func (c *BaseController) SetCatalog() {
    is_catalog := c.GetSession("is_catalog")
    if is_catalog == nil || !is_catalog.(bool) {
        o := orm.NewOrm()
        var cats []catalogs.Catalog
        o.Raw(
            "SELECT c.cid, c.ancestor, c.descendant, ctp.name, ctp.name_i18n" +
            "    FROM catalogs c inner join catalogstreepath ctp" +
            "    WHERE c.cid = ctp.ctpid").QueryRows(&cats)
        c.SetSession("Catalog", catalogs.GetCatalog(cats))
    }
    c.Data["Catalog"] = c.GetSession("Catalog")
}


func (c *BaseController) SetPurchases() {
    Sum := c.GetSession("Sum")
    Count := c.GetSession("PurchaseCount")
    if Sum == nil {
        c.Data["Sum"] = 0
        c.Data["PurchaseCount"] = 0
    } else {
        c.Data["Sum"] = Sum
        c.Data["PurchaseCount"] = Count
    }
}
