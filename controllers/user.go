package controllers

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/BabichMikhail/WebShopOnBeeGO/models"
    "crypto/sha1"
    "fmt"
)

type UserController struct {
    BaseController
}

func (c *UserController) Get() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    c.SetAuthorized()
    if c.GetSession("authorized") == nil || c.GetSession("authorized").(bool) == false {
        c.Redirect("/webshop/login", 302)
        return
    }
    c.SetCatalog()
    c.SetPurchases()
    var oldPurchases []models.Purchase
    o := orm.NewOrm()
    o.Raw(`SELECT * FROM purchases p WHERE p.purchaser_id = ?`, c.GetSession("userId").(int)).QueryRows(&oldPurchases)
    c.Data["UserPurchases"] = oldPurchases
    c.TplName = "user.tpl"
}

func (c *UserController) Post() {
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    c.TplName = "user.tpl"
    if c.GetSession("authorized") == nil || c.GetSession("authorized").(bool) == false {
        c.Redirect("/webshop/login", 302)
        return
    }
    userId := c.GetSession("userId").(int)
    oldPassword := c.GetString("OldPassword")
    newPassword := c.GetString("Password")
    newRePassword := c.GetString("Repassword")
    o := orm.NewOrm()
    var user models.User
    o.Raw(`SELECT * FROM users u WHERE u.Id = ?`, userId).QueryRow(&user)
    if fmt.Sprintf("%x", sha1.Sum([]byte(oldPassword))) != user.Password {
        c.Data["Error"] = true
        c.Data["Err_msg"] = "Неверный пароль"
        return
    }
    if newPassword != newRePassword {
        c.Data["Error"] = true
        c.Data["Err_msg"] = "Пароли не совпадают"
        return
    }
    o.Raw(`UPDATE users SET password = ? WHERE id = ?`, fmt.Sprintf("%x", sha1.Sum([]byte(newRePassword))), userId).Exec()
    c.Redirect("/webshop/user", 302)
}
