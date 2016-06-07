package controllers

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "C"
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
    var Sum C.ulonglong
    Sum = 0
    PCount := 0
    o := orm.NewOrm()
    var eq struct {
        Price int
    }
    o.Raw("SELECT e.price FROM equipments e WHERE e.equip_id = ?", EquipId).QueryRow(&eq)
    if S == nil {
        Sum = C.ulonglong(eq.Price)
        PCount = 1
    } else {
        Sum = S.(C.ulonglong) + C.ulonglong(Purchases[EquipId]*eq.Price)
        PCount = Count.(int) + 1
    }
    fmt.Println(Sum, PCount, eq.Price)
    c.SetSession("Sum", Sum)
    c.SetSession("PurchaseCount", PCount)
    c.RedirectOnLastPage()
}
