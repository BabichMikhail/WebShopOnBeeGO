package controllers

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type AddProductController struct {
    BaseController
}

func (c *AddProductController) Get() {
    EquipId := c.Ctx.Input.Param(":equip_id")
    P := c.GetSession("Purchases")
    Purchases := map[string]int{}
    if P != nil {
        Purchases = P.(map[string]int)
    }
    Purchases[EquipId]++
    c.SetSession("Purchases", Purchases)
    S := c.GetSession("Sum")
    Count := c.GetSession("PurchaseCount")
    var Sum int64
    Sum = 0
    PCount := 0
    o := orm.NewOrm()
    var eq struct {
        Price int
    }
    o.Raw("SELECT e.price FROM equipments e WHERE e.equip_id = ?", EquipId).QueryRow(&eq)
    if S == nil {
        Sum = orm.ToInt64(eq.Price)
        PCount = 1
    } else {
        Sum = orm.ToInt64(S) + orm.ToInt64(Purchases[EquipId]*eq.Price)
        PCount = Count.(int) + 1
    }
    c.SetSession("Sum", Sum)
    c.SetSession("PurchaseCount", PCount)
    c.RedirectOnLastPage()
}
