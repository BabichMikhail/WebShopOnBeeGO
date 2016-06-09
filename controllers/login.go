package controllers

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "WebShopOnBeeGO/models"
    "crypto/sha1"
    "fmt"
)

type LoginController struct {
    BaseController
}

func (c *LoginController) Signup() {
    c.TplName = "signup.tpl"
    c.SetCatalog()
    if !c.Ctx.Input.IsPost() {
        return
    }
    login := c.GetString("Login")
    password := c.GetString("Password")
    repassword := c.GetString("Repassword")
    if password != repassword {
        c.Data["Error"] = true
        c.Data["Err_msg"] = "Пароли не совпадают"
        return
    }
    data_pwd := []byte(password)
    hash_pwd := fmt.Sprintf("%x", sha1.Sum(data_pwd))
    user := models.User{ Login: login, Password: hash_pwd, Rights: 3 }
    fmt.Println(hash_pwd)
    o := orm.NewOrm()
    o.Begin()
    _, err := o.Insert(&user)
    fmt.Println(err)
    if err == nil {
        o.Commit()
    } else {
        o.Rollback()
    }
    c.SetSession("authorized", true)
    c.SetSession("username", login)
    c.SetSession("userright", 3)
    c.SetSession("userId", user.Id)
    c.RedirectOnLastPage()
}

func (c *LoginController) Login() {
    c.TplName = "login.tpl"
    c.SetCatalog()
    c.SetPurchases()
    if !c.Ctx.Input.IsPost() {
        return
    }
    login := c.GetString("Login")
    password := c.GetString("Password")
    data := []byte(password)
    hash := fmt.Sprintf("%x", sha1.Sum(data))
    fmt.Println(hash)
    o := orm.NewOrm()
    var user models.User
    err := o.Raw(fmt.Sprintf(`SELECT id, login, password, rights FROM users WHERE login = "%s"`, login)).QueryRow(&user)
    fmt.Println(user.Password, hash, user.Login, err)
    if user.Password == hash {
        c.SetSession("authorized", true)
        c.SetSession("userId", user.Id)
        c.SetSession("username", login)
        c.SetSession("userright", user.Rights)
        c.RedirectOnLastPage()
        return
    }
    c.Data["Error"] = true
    c.Data["Err_msg"] = "Неверное имя пользователя или пароль"
}

func (c *LoginController) Logout() {
    c.SetSession("authorized", false)
    c.SetSession("userId", nil)
    c.SetSession("username", nil)
    c.SetSession("userright", 3)
    c.RedirectOnLastPage()
}
