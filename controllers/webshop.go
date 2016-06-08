package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "WebShopOnBeeGO/models"
    "strconv"
    "fmt"
)

type AdminController struct {
    BaseController
}

func (c *AdminController) Get() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    if c.GetSession("authorized") == nil || c.GetSession("userright").(int) != 0 {
        c.Redirect("/webshop", 302)
        return
    }
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    o := orm.NewOrm()
    var purchases []models.Purchase
    o.Raw(`SELECT p.id, p.sum, p.count, p.date FROM purchases p`).QueryRows(&purchases)
    var goods []models.Good
    num, err := o.Raw(`SELECT g.purchase_id, g.count, g.name, g.price, g.equip_id FROM goods g`).QueryRows(&goods)
    fmt.Println(num, err)
    c.Data["Purchases"] = purchases
    c.Data["Goods"] = goods
    c.TplName = "admin.tpl"
}

type WebShopController struct {
    BaseController
}

func (c *WebShopController) HomePage() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    var equipments []models.EquipHomePage
    o := orm.NewOrm()
    o.Raw(`SELECT e.equip_id, e.name, e.price, e.small_image FROM equipments e`).QueryRows(&equipments)
    extEquipments := make([]models.EquipHomePageIsCount, len(equipments))
    P := c.GetSession("Purchases")
    var purchases map[string]int
    if P != nil {
        purchases = P.(map[string]int)
    } else {
        purchases = map[string]int{}
    }
    for key, value := range equipments {
        extEquipments[key].EquipHomePage = value
        extEquipments[key].IsCount = purchases[strconv.Itoa(value.Equip_id)] > 0
    }
    c.Data["Equipment"] = extEquipments
    c.TplName = "main.tpl"
}

func (c *WebShopController) Catalog() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    c.Data["Fields"] = &models.Fields {
        Id:     "Код товара",
        Name:   "Название",
        Price:  "Стоимость",
        Nation: "Нация",
    }
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    var equipments []models.EquipInTable
    query := `SELECT e.equip_id, e.name, e.price, e.small_image, n.name_i18n as "nation" FROM equipments e `
    var query_WHERE string
    equip_name := c.Ctx.Input.Param(":equip")
    equip_type := c.Ctx.Input.Param(":type")
    equip_level := c.Ctx.Input.Param(":level")
    if equip_type != "" {
        query += `INNER JOIN types t ON e.type = t.name `
        query_WHERE += fmt.Sprintf(`WHERE t.name_catalog = "%s" and e.equip_type = "%s" `,
            equip_type, equip_name)
    } else if equip_name != "" {
        query_WHERE += fmt.Sprintf(`WHERE e.equip_type = "%s" `, equip_name)
    }
    o := orm.NewOrm()
    if equip_level != "" {
        var levels []struct { Level int }
        o.Raw(fmt.Sprintf(`SELECT l.level FROM levels l WHERE l.value = "%s"`, equip_level)).QueryRows(&levels)
        add_query_WHERE := ""
        for _, l := range levels {
            if add_query_WHERE != "" {
                add_query_WHERE += "OR "
            }
            add_query_WHERE += fmt.Sprintf("e.level = %d ", l.Level)
        }
        if query_WHERE == "" {
            query_WHERE = "WHERE "
        } else {
            query_WHERE += "AND "
        }
        query_WHERE += "(" + add_query_WHERE + ")"
    }
    query += "INNER JOIN nations n ON n.name = e.nation "
    o.Raw(query + query_WHERE).QueryRows(&equipments)
    fmt.Println(len(equipments))
    extEquipments := make([]models.EquipInTableIsCount, len(equipments))
    P := c.GetSession("Purchases")
    var purchases map[string]int
    if P != nil {
        purchases = P.(map[string]int)
    } else {
        purchases = map[string]int{}
    }
    i := 0
    for _, value := range equipments {
        extEquipments[i] = models.EquipInTableIsCount{
            IsCount: purchases[strconv.Itoa(value.Equip_id)] != 0,
            EquipInTable: value,
        }
        i++
    }
    c.Data["Equipment"] = extEquipments
    c.TplName = "grid.tpl"
}

type MainController struct {
    beego.Controller
}

func (c *MainController) Get() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"] = "astaxie@gmail.com"
    c.TplName = "index.tpl"
}
