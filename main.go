package main

import (
    _ "WebShopOnBeeGO/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
    models "WebShopOnBeeGO/models"
)

func init() {
    orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
    orm.RegisterModel(new(models.Equipment))
    orm.RegisterModel(new(models.User))
    orm.RegisterModel(new(models.Catalogs))
    orm.RegisterModel(new(models.CatalogsTreePath))
    orm.RegisterModel(new(models.Type))
    orm.RegisterModel(new(models.Nation))
    orm.RegisterModel(new(models.Level))
    //orm.RunSyncdb("default", true, true)
}

func main() {
    beego.Run()
}
