package controllers

import (
    "github.com/astaxie/beego"
    "github.com/BabichMikhail/WebShopOnBeeGO/catalogs"
    "github.com/astaxie/beego/orm"
    "strconv"
    "fmt"
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
        c.Data["IsAdmin"] = false
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

func (c *BaseController) RedirectOnLastPage() {
    url := c.GetSession("LastUrl")
    if url != nil {
        c.Redirect(url.(string), 302)
        return
    }
    c.Redirect("/webshop", 302)
}

func (c *BaseController) UpdatePurchases() {
    P := c.GetSession("Purchases")
    if P == nil {
        c.SetSession("Sum", 0)
        c.SetSession("PurchaseCount", 0)
    }
    purchases := P.(map[string]int)
    Sum := 0
    Count := 0
    o := orm.NewOrm()
    var eqs[]struct{Equip_id int; Price int}
    query := ""
    for key, _ := range purchases {
        if query != "" {
            query += ","
        }
        query += key
    }
    o.Raw(fmt.Sprintf("SELECT e.equip_id, e.price FROM equipments e WHERE e.equip_id IN (%s)", query)).QueryRows(&eqs)
    for _, value := range eqs {
        Sum += value.Price*purchases[strconv.Itoa(value.Equip_id)]
        Count += purchases[strconv.Itoa(value.Equip_id)]
    }
    c.SetSession("Sum", Sum)
    c.SetSession("PurchaseCount", Count)
}
