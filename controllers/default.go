package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    "WebShopOnBeeGO/models"
    "WebShopOnBeeGO/catalogs"
    "crypto/md5"
    "fmt"
    "regexp"
)

type Fields struct {
    Id          string
    Name        string
    Price       string
    Nation      string
}

type Equipment struct {
    Equip_id    int
    Name        string
    Price       int
    Nation      string
}

type BaseController struct {
    beego.Controller
}

func (c *BaseController) SetAuthorized() {
    authorized := c.GetSession("authorized")
    if authorized != nil && authorized.(bool) {
        c.Data["Authorized"] = true
        c.Data["Username"] = c.GetSession("username")
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
            "SELECT c.cid, c.ancestor, c.descendant, ctp.name, ctp.name_i18n " +
            "    FROM catalogs c inner join catalogstreepath ctp" +
            "    WHERE c.cid = ctp.ctpid").QueryRows(&cats)
        c.SetSession("Catalog", catalogs.GetCatalog(cats))
    }
    c.Data["Catalog"] = c.GetSession("Catalog")
}

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
    h := md5.New()
    data_pwd := []byte(password)
    hash_pwd := fmt.Sprintf("%x", h.Sum(data_pwd))
    user := models.User{Login: login, Password: hash_pwd}
    fmt.Println("pwd_hash ", hash_pwd)
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
    c.Redirect("/webshop", 302)
}

func (c *LoginController) Login() {
    c.TplName = "login.tpl"
    c.SetCatalog()
    if !c.Ctx.Input.IsPost() {
        return
    }
    login := c.GetString("Login")
    password := c.GetString("Password")
    h := md5.New()
    data := []byte(password)
    hash := fmt.Sprintf("%x", h.Sum(data))
    fmt.Println(hash)
    o := orm.NewOrm()
    var user models.User
    o.Raw("SELECT login, password FROM user WHERE login = ?", login).QueryRow(&user)
    if user.Password == hash {
        c.SetSession("authorized", true)
        c.SetSession("username", login)
        c.Redirect("/webshop", 302)
        return
    }
    c.Data["Error"] = true
    c.Data["Err_msg"] = "Неверное имя пользователя или пароль"
}

func (c *LoginController) Logout() {
    c.SetSession("authorized", false)
    c.Redirect("/webshop", 302)
}

type WebShopController struct {
    BaseController
}

func (c *WebShopController) HomePage() {
    c.SetAuthorized()
    c.SetCatalog()
    c.TplName = "main.tpl"
}

func (c *WebShopController) Catalog() {
    c.Data["Fields"] = &Fields {
        Id:     "Код товара",
        Name:   "Название",
        Price:  "Стоимость",
        Nation: "Нация",
    }
    c.SetAuthorized()
    c.SetCatalog()
    var equipments []Equipment
    r_equip_types, _ := regexp.Compile("webshop/catalog/([a-z]+)")
    url := c.Ctx.Request.URL.String()
    var equip_type string
    var is_equip_type bool
    if r_equip_types.MatchString(url) {
        is_equip_type = true
        equip_type = r_equip_types.FindStringSubmatch(url)[1]
    }
    query := `SELECT e.equip_id, e.name, e.price, n.name_i18n as "nation" FROM equipments e `
    var query_WHERE string
    r_types, _ := regexp.Compile(fmt.Sprintf("/webshop/catalog/%s/([a-z_]+)", equip_type))
    if r_types.MatchString(url) {
        query += `INNER JOIN types t ON e.type = t.name `
        query_WHERE += fmt.Sprintf(`WHERE t.name_catalog = "%s" and e.equip_type = "%s"`,
            r_types.FindStringSubmatch(url)[1], equip_type)
    } else if is_equip_type {
        query_WHERE += fmt.Sprintf(`WHERE e.equip_type = "%s" `, equip_type)
    }
    r_levels, _ := regexp.Compile("/webshop/catalog/[a-z]+/[a-z_]+/([a-z0-9_]+)")
    o := orm.NewOrm()
    if r_levels.MatchString(url) {
        value := r_levels.FindStringSubmatch(url)[1];
        var levels []struct { Level int }
        o.Raw(fmt.Sprintf(`SELECT l.level FROM levels l WHERE l.value = "%s"`, value)).QueryRows(&levels)
        add_query_WHERE := ""
        for _, l := range levels {
            if add_query_WHERE != "" {
                add_query_WHERE += "or "
            }
            add_query_WHERE += fmt.Sprintf("e.level = %d ", l.Level)
        }
        if query_WHERE == "" {
            query_WHERE = "WHERE "
        } else {
            query_WHERE += "and "
        }
        query_WHERE += "(" + add_query_WHERE + ")"

    }
    query += "INNER JOIN nations n ON n.name = e.nation "
    o.Raw(query + query_WHERE).QueryRows(&equipments)
    c.Data["Equipment"] = equipments
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
