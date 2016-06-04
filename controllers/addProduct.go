package controllers

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type AddProductController struct {
    BaseController
}

func (c *AddProductController) Get() {
    Equip_id := c.Ctx.Input.Param(":equip_id")
    P := c.GetSession("Purchases")
    Purchases := map[string]int{}
    if P != nil {
        Purchases = P.(map[string]int)
    }
    Purchases[Equip_id]++
    c.SetSession("Purchases", Purchases)
    S := c.GetSession("Sum")
    Count := c.GetSession("PurchaseCount")
    Sum := 0
    PCount := 0
    o := orm.NewOrm()
    var eq struct {
        Price int
    }
    o.Raw("SELECT e.price FROM equipments e WHERE e.equip_id = ?", Equip_id).QueryRow(&eq)
    if S == nil {
        Sum = eq.Price
        PCount = 1
    } else {
        Sum = S.(int) + Count.(int)*eq.Price
        PCount = Count.(int) + 1
    }
    c.SetSession("Sum", Sum)
    c.SetSession("PurchaseCount", PCount)
}
