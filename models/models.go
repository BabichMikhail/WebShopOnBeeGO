package models

type User struct {
    Id            int64
    Login         string    `orm:"size(64);unique" valid:"Required"`
    Password      string    `orm:"size(64)"        valid:"Required;MinSize(3)"`
}

func (user *User) TableName() string {
    return "user"
}

type CatalogsTreePath struct {
    Id              int     `form:"-"`
    Ctpid           int     `form:"-"`
    Name            string  `form:"name,text,name"`
    Name_i18n       bool    `form:"name_i18n,text,name_i18n"`
}

func (catalogstreepath *CatalogsTreePath) TableName() string {
    return "catalogstreepath"
}

type Catalogs struct {
    Id              int     `form:"-"`
    Cid             int     `form:"-"`
    Ancestor        int     `form:"-"`
    Descendant      int     `form:"-"`
}

func (catalogs *Catalogs) TableName() string {
    return "catalogs"
}

type Equipment struct {
    Id              int     `form:"-"`
    Description     string  `form:"description,text,description"`
    Equip_id        int     `form:"equip_id,int,equip_id"`
    Equip_type      string  `form:"equip_type,text,equip_type"`
    Image           string  `form:"image,text,image"`
    Is_gift         bool    `form:"is_gift,bool,is_gift"`
    Is_premium      bool    `form:"is_premium,bool,is_premium"`
    Level           int     `form:"level,int,level"`
    Name            string  `form:"name,text,name"`
    Nation          string  `form:"nation,text,nation"`
    Price           int     `form:"price,int,price"`
    Short_name      string  `form:"short_name,text,short_name"`
    Small_image     string  `form:"small_image,text,small_image"`
    Type            string  `form:"type,text,type"`
}

func (equipment *Equipment) TableName() string {
    return "equipments"
}

type Type struct {
    Id              int     `form:"-"`
    Name            string  `form:"name,text,name"`
    Name_catalog    string  `form:"name_catalog,text,name_catalog"`
}

func (t *Type) TableName() string {
    return "types"
}

type Level struct {
    Id              int     `form:"-"`
    Value           string  `form:"value,text,value"`
    Level           int     `form:"level,int,level"`
}

func (t *Level) TableName() string {
    return "levels"
}

type Nation struct {
    Id              int     `form:"-"`
    Name            string  `form:"name,text,name"`
    Name_i18n       string  `form:"name_i18n,text,name_i18n"`
}

func (t *Nation) TableName() string {
    return "nations"
}
