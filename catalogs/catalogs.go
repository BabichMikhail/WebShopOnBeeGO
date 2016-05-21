package catalogs

import "fmt"

var Ready_cat string

type Catalog struct {
    Cid         int     `orm:"cid,int,cid"`
    Name        string  `orm:"name,text,name"`
    Name_i18n   string  `orm:"name_i18n,text,name_i18n"`
    Ancestor    int     `orm:"ancestor,int,ancestor"`
    Descendant  int     `orm:"descendant,int,descendant"`
}

func CreateCatalog(cat []Catalog, s string, cat_id int) string {
    count := 0
    var name, name_i18n string
    for _, value := range cat {
        if value.Cid == cat_id {
            count++
            name = value.Name
            name_i18n = value.Name_i18n
        }
    }
    s += "/" + name
    ans := fmt.Sprintf(`<li><a href="%s">%s</a>`, s, name_i18n)
    if count > 0 {
        ans += `<ul>`
    }
    for _, value := range cat {
        if value.Cid == cat_id && value.Descendant != 0 {
            ans += CreateCatalog(cat, s, value.Descendant)
        }
    }
    if count > 0 {
        ans += `</ul></li>`
    }
    return ans
}

func GetCatalog(cat []Catalog) string {
    if Ready_cat == `` {
        used_cid := map[int]bool{}
        Ready_cat += `<div><ul><li><a href="/webshop/catalog">Каталог</a></li>`
        for _, value := range cat {
            if value.Ancestor == 0 && used_cid[value.Cid] == false {
                used_cid[value.Cid] = true
                Ready_cat += CreateCatalog(cat, `/webshop/catalog`, value.Cid)
            }
        }
        Ready_cat += `</ul></div></div>`
    }
    return Ready_cat
}
