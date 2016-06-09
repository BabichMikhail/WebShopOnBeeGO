package controllers

import (
    "github.com/astaxie/beego/orm"
    "WebShopOnBeeGO/models"
    "github.com/fatih/structs"
    "crypto/sha1"
    "strings"
    "reflect"
    "time"
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
    o.Raw(`SELECT g.purchase_id, g.count, g.name, g.price, g.equip_id FROM goods g`).QueryRows(&goods)
    c.Data["Purchases"] = purchases
    c.Data["Goods"] = goods
    c.TplName = "admin.tpl"
}

func (c *AdminController) SaveGoodsInfo() {
    if c.GetSession("authorized") == nil || c.GetSession("userright").(int) != 0 {
        c.Redirect("/webshop", 302)
        return
    }
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    equipId := c.Ctx.Input.Param(":equip_id")
    tableName := c.Ctx.Input.Param(":tableName")
    descriptions := GetCharacteristics(tableName, equipId, true)
    o := orm.NewOrm()
    var pairs []string
    for _, value := range descriptions {
        newValue := c.GetString(value.PropName)
        switch value.Type {
        case "string": newValue = fmt.Sprintf(`"%s"`, newValue)
        case "bool":
            if newValue == "Да" {
                newValue = "1"
            } else {
                newValue = "0"
            }
        }
        pairs = append(pairs, value.PropName + "=" + newValue)
    }
    o.Raw(fmt.Sprintf("UPDATE %s SET %s WHERE equip_id = %s;", tableName, strings.Join(pairs, ","), equipId)).Exec()
    c.RedirectOnLastPage()
}

func (c *AdminController) AdminHome() {
    c.Redirect("/webshop/admin/edittable/users", 302)

    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    if c.GetSession("authorized") == nil || c.GetSession("userright").(int) != 0 {
        c.Redirect("/webshop", 302)
        return
    }
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    c.TplName = "adminhome.tpl"
}

func HeadersInOrder(t reflect.Type, names []string) []models.Header {
    headers := make([]models.Header, len(names))
    for _, fieldName := range names {
        field, _ := t.FieldByName(fieldName)
        index, _ := strconv.Atoi(field.Tag.Get("index"))
        headers[index].Type = field.Tag.Get("type")
        fmt.Println("index = ", field.Tag.Get("index"), fieldName, headers[index].Type)
        headers[index].Name = fieldName
        headers[index].Namei18n = field.Tag.Get("namei18n")
        headers[index].Pattern = field.Tag.Get("pattern")
        formtag := field.Tag.Get("formtag")
        headers[index].IsTextarea = formtag != "" && formtag == "textarea"
    }
    return headers
}

func ToMapStringString(m map[string]interface{}, headers []models.Header) map[string]string {
    ans := map[string]string{}
    for _, header := range headers {
        if m[header.Name] == nil {
            ans[header.Name] = ""
            fmt.Println("Warning: %s is nil", header.Name)
            continue
        }
        switch header.Type {
        case "int64": ans[header.Name] = strconv.FormatInt(orm.ToInt64(m[header.Name]), 10)
        case "int": ans[header.Name] = strconv.Itoa(m[header.Name].(int))
        case "string": ans[header.Name] = m[header.Name].(string)
        case "bool": ans[header.Name] = strconv.FormatBool(m[header.Name].(bool))
        case "time.Time":
            fmt.Println(m[header.Name].(time.Time).String())
            ans[header.Name] =  m[header.Name].(time.Time).String()
        }
    }
    return ans
}

func SetDataRow(bm map[string]interface{}, headers []models.Header) (row []string) {
    m := ToMapStringString(bm, headers)
    for _, header := range headers {
        if header.Name == "Password" {
            row = append(row, "")
        } else {
            row = append(row, m[header.Name])
        }
    }
    return
}

func (c *AdminController) SetEmptyData(bm map[string]interface{}, headers []models.Header) {
    m := ToMapStringString(bm, headers)
    var row []string
    for _, header := range headers {
        row = append(row, m[header.Name])
    }
    c.Data["EmptyData"] = row
}

func (c *AdminController) EditTable() {
    c.SetSession("LastUrl", c.Ctx.Request.URL.String())
    if c.GetSession("authorized") == nil || c.GetSession("userright").(int) != 0 {
        c.Redirect("/webshop", 302)
        return
    }
    c.SetAuthorized()
    c.SetCatalog()
    c.SetPurchases()
    o := orm.NewOrm()
    Name := c.Ctx.Input.Param(":tableName")

    var Headers []models.Header
    TableNames := []string {
        "users", "nations", "types", "levels", "equipments",
        "warplanes", "tanks", "purchases", "goods", "translates",
    }
    var modelType interface{}
    var Data [][]string
    switch Name {
    case "users":
        modelType = models.User{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.User
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
    case "nations":
        modelType = models.Nation{}
        Headers = HeadersInOrder(reflect.TypeOf(models.Nation{}), structs.Names(models.Nation{}))
        var data []models.Nation
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
    case "types":
        modelType = models.Type{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Type
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
    case "levels":
        modelType = models.Level{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Level
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
    case "equipments":
        modelType = models.Equipment{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Equipment
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
        c.Data["IsGoodsPage"] = true;
    case "warplanes":
        modelType = models.Warplane{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Warplane
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
        c.Data["IsGoodsPage"] = true;
    case "tanks":
        modelType = models.Tank{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Tank
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
        c.Data["IsGoodsPage"] = true;
    case "purchases":
        modelType = models.Purchase{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Purchase
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
    case "goods":
        modelType = models.Good{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Good
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
        c.Data["IsGoodsPage"] = true;
    case "translates":
        modelType = models.Translate{}
        Headers = HeadersInOrder(reflect.TypeOf(modelType), structs.Names(modelType))
        var data []models.Translate
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        for _, value := range data {
            Data = append(Data, SetDataRow(structs.Map(value), Headers))
        }
    }
    c.SetEmptyData(structs.Map(modelType), Headers)
    c.Data["TableNames"] = TableNames
    c.Data["Headers"] = Headers
    c.Data["Data"] = Data
    c.TplName = "admintables.tpl"
 }

func (c* AdminController) GetFieldNamesAndValues(names []string, t reflect.Type, id string) (fieldNames []string, values []string) {
    for _, name := range names {
        if name == "Id" {
            continue
        }
        if name == "Password" {
            values = append(values, fmt.Sprintf(`"%x"`, sha1.Sum([]byte(c.GetString(name + id)))))
            fieldNames = append(fieldNames, name)
            continue
        }
        if name == "Date" {
            continue
        }
        newValue := c.GetString(name + id)
        field, _ := t.FieldByName(name)
        switch field.Tag.Get("type") {
        case "string": newValue = fmt.Sprintf(`"%s"`, newValue)
        case "bool":
            if newValue == "true" {
                newValue = "'true'"
            } else {
                newValue = "'false'"
            }
        // todo
        //case "time.Time":
        }
        values = append(values, newValue)
        fieldNames = append(fieldNames, name)
        fmt.Println(c.GetString(name + id), name)
    }
    return
}

func (c* AdminController) DoIfNeed(allFieldNames []string, modelType interface{}, id int, name string) {
    o := orm.NewOrm()
    if c.GetString("checkremove" + strconv.Itoa(id)) == "on"{
        o.Raw(fmt.Sprintf(`DELETE FROM %s WHERE id = %d`, name, id)).Exec()
    } else if c.GetString("checkedit" + strconv.Itoa(id)) == "on" {
        var pairs []string
        fieldNames, values := c.GetFieldNamesAndValues(allFieldNames, reflect.TypeOf(modelType), strconv.Itoa(id))
        for idx := range fieldNames {
            pairs = append(pairs, fieldNames[idx] + "=" + values[idx])
        }
        o.Raw(fmt.Sprintf(`UPDATE %s SET %s WHERE id = %d`, name, strings.Join(pairs, ","), id)).Exec()
    }
}

func (c* AdminController) ApplyChanges() {
    if c.GetSession("authorized") == nil || c.GetSession("userright").(int) != 0 {
        c.Redirect("/webshop", 302)
        return
    }
    Name := c.Ctx.Input.Param(":tableName")
    o := orm.NewOrm()
    var modelType interface{}
    var allFieldNames []string
    switch Name {
    case "users":
        var data []models.User
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.User{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "nations":
        var data []models.Nation
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Nation{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "types":
        var data []models.Type
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Type{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "levels":
        var data []models.Level
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Level{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "equipments":
        var data []models.Equipment
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Equipment{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "warplanes":
        var data []models.Warplane
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Warplane{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "tanks":
        var data []models.Tank
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Tank{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "purchases":
        var data []models.Purchase
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Purchase{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "goods":
        var data []models.Good
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Good{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    case "translates":
        var data []models.Translate
        o.Raw(fmt.Sprintf(`SELECT * FROM %s`, Name)).QueryRows(&data)
        modelType = models.Translate{}
        allFieldNames = structs.Names(modelType)
        for _, value := range data {
            c.DoIfNeed(allFieldNames, modelType, value.Id, Name)
        }
    }
    if c.GetString("checkedit0") == "on" {
        fieldNames, values := c.GetFieldNamesAndValues(allFieldNames, reflect.TypeOf(modelType), "0")
        err, num := o.Raw(fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, Name, strings.Join(fieldNames, ","), strings.Join(values, ","))).Exec()
        fmt.Printf(`INSERT INTO %s (%s) VALUES (%s)`, Name, strings.Join(fieldNames, ","), strings.Join(values, ","))
        fmt.Println(err, num)
    }
    c.RedirectOnLastPage()
}
