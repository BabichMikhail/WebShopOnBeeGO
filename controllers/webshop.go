package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/BabichMikhail/WebShopOnBeeGO/models"
    "strings"
    "strconv"
    "fmt"
)

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
    current_place := "Каталог"
    o := orm.NewOrm()
    var name_i18n struct { Name_i18n string }
    if equip_name != "" {
        o.Raw(`SELECT t.name_i18n FROM translates t WHERE t.name = ?`, equip_name).QueryRow(&name_i18n)
        current_place += fmt.Sprintf(" / %s", name_i18n.Name_i18n)
        if equip_type != "" {
            o.Raw(`SELECT t.name_i18n FROM translates t WHERE t.name = ?`, equip_type).QueryRow(&name_i18n)
            current_place += fmt.Sprintf(" / %s", name_i18n.Name_i18n)
            if equip_level != "" {
                o.Raw(`SELECT t.name_i18n FROM translates t WHERE t.name = ?`, equip_level).QueryRow(&name_i18n)
                current_place += fmt.Sprintf(" / %s", name_i18n.Name_i18n)
            }
        }
    }
    c.Data["CurrentPlace"] = current_place
    if equip_type != "" {
        query += `INNER JOIN types t ON e.type = t.name `
        query_WHERE += fmt.Sprintf(`WHERE t.name_catalog = "%s" and e.equip_type = "%s" `,
            equip_type, equip_name)
    } else if equip_name != "" {
        query_WHERE += fmt.Sprintf(`WHERE e.equip_type = "%s" `, equip_name)
    }

    if equip_level != "" {
        var levels []struct { Level int }
        o.Raw(fmt.Sprintf(`SELECT l.level FROM levels l WHERE l.value = "%s"`, equip_level)).QueryRows(&levels)
        var values []string
        for _, l := range levels {
            values = append(values, strconv.Itoa(l.Level))
        }
        add_query_WHERE := fmt.Sprintf("e.level IN (%s)", strings.Join(values, ","))
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
