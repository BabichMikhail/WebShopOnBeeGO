package controllers

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/fatih/structs"
    "WebShopOnBeeGO/models"
    "reflect"
    "strconv"
    "time"
    "fmt"
)

type CheckoutController struct {
    BaseController
}

func (c *CheckoutController) Get() {
    o := orm.NewOrm()
    P := c.GetSession("Purchases")
    if P != nil {
        goods := P.(map[string]int)
        var sum int64
        sum = 0
        if c.GetSession("Sum") != nil {
            sum = orm.ToInt64(c.GetSession("Sum"))
        }
        count := 0
        if c.GetSession("PurchaseCount") != nil {
            count = c.GetSession("PurchaseCount").(int)
        }
        fmt.Println("Sum =", sum)
        purchase := models.Purchase{ Sum: sum, Count: count, Date: time.Now()}
        o.Begin()
        id, err := o.Insert(&purchase)
        if err == nil {
            o.Commit()
        } else {
            o.Rollback()
        }
        query := ""
        for k := range goods {
            if query != "" {
                query += ", "
            }
            query += k
        }
        var Costs []struct{ Equip_id int; Price int; Name string }
        num, err := o.Raw(fmt.Sprintf(`SELECT e.equip_id, e.name, e.price FROM equipments e WHERE e.equip_id IN (%s)`, query)).QueryRows(&Costs)
        fmt.Println(num, " ", err)
        for _, value := range Costs {
            fmt.Println(value.Equip_id, id, value.Price)
            _, err := o.Raw(fmt.Sprintf(`INSERT INTO goods (price, count, equip_id, purchase_id, name) VALUES (%d, %d, %d, %d, "%s")`,
                value.Price, goods[strconv.Itoa(value.Equip_id)], value.Equip_id, id, value.Name)).Exec();
            if err != nil {
                fmt.Println(err)
            }
        }
        c.SetSession("Purchases", nil)
        c.SetSession("Sum", 0)
        c.SetSession("PurchaseCount", 0)
    }
    c.Redirect("/webshop", 302)
}

type PurchaseController struct {
    BaseController
}

func (c *PurchaseController) Post() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    P := c.GetSession("Purchases")
    if P != nil {
        purchases := P.(map[string]int)
        query := ""
        for k := range purchases {
            if query != "" {
                query += ", "
            }
            query += k
        }
        c.SetSession("Purchases", purchases)
        Sum := 0
        PCount := 0
        o := orm.NewOrm()
        var eqs[]struct{Equip_id int; Price int}
        o.Raw(fmt.Sprintf("SELECT e.equip_id, e.price FROM equipments e WHERE e.equip_id IN (%s)", query)).QueryRows(&eqs)
        for _, value := range eqs {
            Sum += value.Price*purchases[strconv.Itoa(value.Equip_id)]
            PCount += purchases[strconv.Itoa(value.Equip_id)]
        }
        c.SetSession("Sum", Sum)
        c.SetSession("PurchaseCount", PCount)
    }
    c.Redirect("/webshop/purchases/checkout", 302)
}

func (c *PurchaseController) Get() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    P := c.GetSession("Purchases")
    if P != nil {
        purchases := P.(map[string]int)
        i := 0
        query := ""
        for k := range purchases {
            if i != 0 {
                query += " OR "
            }
            query += "e.equip_id = " + k
            i++
        }
        o := orm.NewOrm()
        var equipments []models.EquipInTable
        o.Raw(fmt.Sprintf(`SELECT e.equip_id, e.name, e.price, e.small_image, n.name_i18n as "nation" FROM equipments e ` +
            `INNER JOIN nations n ON n.name = e.nation WHERE %s`, query)).QueryRows(&equipments)
        ext_equipments := make([]models.ExtEquipInTable, len(equipments))
        for i, eq := range equipments {
            ext_equipments[i].Equip_id = eq.Equip_id
            ext_equipments[i].Small_image = eq.Small_image
            ext_equipments[i].Name = eq.Name
            ext_equipments[i].Nation = eq.Nation
            ext_equipments[i].Price = eq.Price
            ext_equipments[i].Count = purchases[strconv.Itoa(eq.Equip_id)]
        }
        c.Data["Equipment"] = ext_equipments
    }
    c.TplName = "purchases.tpl"
}

func (c *PurchaseController) Change() {
    equip_id := c.Ctx.Input.Param(":equip_id")
    add := c.Ctx.Input.Param(":add")
    is_add := add == "1"
    P := c.GetSession("Purchases")
    var purchases map[string]int
    if P != nil {
        purchases = P.(map[string]int)
        if is_add {
            purchases[equip_id]++
        } else {
            purchases[equip_id]--
            if purchases[equip_id] == -1 {
                purchases[equip_id] = 0
            }

        }
    } else {
        purchases = map[string]int{}
        if is_add {
            purchases[equip_id] = 1
        } else {
            purchases[equip_id] = 0
        }
    }
    c.SetSession("Purchases", purchases)
    fmt.Println(equip_id, add, is_add)
    c.UpdatePurchases()
    c.RedirectOnLastPage()
}

func (c *PurchaseController) Delete() {
    P := c.GetSession("Purchases")
    if P == nil {
        c.RedirectOnLastPage()
        return
    }
    delete(P.(map[string]int), c.Ctx.Input.Param(":equip_id"))
    c.UpdatePurchases()
    c.RedirectOnLastPage()
}

func TankCharacteristics(EquipId string) (t reflect.Type, names []string, chs map[string]interface{}) {
    o := orm.NewOrm()
    names = structs.Names(models.TankCharacteristics{})
    t = reflect.TypeOf(models.TankCharacteristics{})
    var ch models.TankCharacteristics
    o.Raw(fmt.Sprintf(`SELECT e.description, e.is_premium, e.level, e.name, ` +
        `n.name_i18n as "nation", e.price, e.type, e.weight, e.max_weight, e.armor, e.hp, ` +
        `e.speed_forward, e.speed_backward FROM tanks e inner join nations n on ` +
        `e.nation = n.name WHERE e.equip_id = %s`, EquipId)).QueryRow(&ch)
    chs = structs.Map(ch)
    return
}

func WarplaneCharacteristics(EquipId string) (t reflect.Type, names []string, chs map[string]interface{}) {
    o := orm.NewOrm()
    names = structs.Names(models.WarplaneCharacteristics{})
    t = reflect.TypeOf(models.WarplaneCharacteristics{})
    var ch models.WarplaneCharacteristics
    o.Raw(fmt.Sprintf(`SELECT e.description, e.is_premium, e.level, e.name, ` +
        `n.name_i18n as "nation", e.price, e.type, e.weight, e.hp, e.speed_ground, ` +
        `e.maneuverability, e.max_speed, e.stall_speed, e.optimal_height, e.roll_maneuver, ` +
        `e.dive_speed, e.opt_maneuver_speed FROM warplanes e inner join nations n on ` +
        `e.nation = n.name WHERE e.equip_id = %s`, EquipId)).QueryRow(&ch)
    chs = structs.Map(ch)
    return
}

func GetCharacteristics(EquipType string, EquipId string) []struct{Key string; Value string} {
    t, names, characteristics := func (EquipType string) (reflect.Type, []string, map[string]interface{}) {
        if EquipType == "tanks" {
            return TankCharacteristics(EquipId)
        } else if EquipType == "warplanes" {
            return WarplaneCharacteristics(EquipId)
        } else {
            return nil, nil, nil
        }
    }(EquipType)
    descriptions := make([]struct{Key string; Value string}, len(names))
    for _, fieldName := range names {
        field, _ := t.FieldByName(fieldName)
        index, _ := strconv.Atoi(field.Tag.Get("index"))
        key := field.Tag.Get("key")
        var value string
        switch field.Tag.Get("type") {
        case "int":
            i, _ := characteristics[field.Name].(int)
            value = strconv.Itoa(i)
        case "bool":
            b, _ := characteristics[field.Name].(bool)
            if b {
                value = "Да"
            } else {
                value = "Нет"
            }
        case "string":
            value, _ = characteristics[field.Name].(string)
        }
        descriptions[index].Key = key
        descriptions[index].Value = value
    }
    return descriptions
}

type GoodsInfoController struct {
    BaseController
}

func (c *GoodsInfoController) Get() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    o := orm.NewOrm()
    equip_id := c.Ctx.Input.Param(":equip_id")
    var equip struct { Equip_type string; Image string }
    o.Raw(fmt.Sprintf("SELECT e.equip_type, e.image FROM equipments e WHERE e.equip_id = %s", equip_id)).QueryRow(&equip)
    descriptions := GetCharacteristics(equip.Equip_type, equip_id)
    P := c.GetSession("Purchases")
    var purchases map[string]int
    if P != nil {
        purchases = P.(map[string]int)
    } else {
        purchases = map[string]int{}
    }
    c.Data["IsCount"] = purchases[equip_id] > 0

    c.Data["Description"] = descriptions[0].Value
    c.Data["Characteristics"] = descriptions[1:]
    c.Data["Image"] = equip.Image
    c.Data["Equip_id"] = equip_id
    c.TplName = "goodsInfo.tpl"
}
