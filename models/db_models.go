package models

import "time"

type User struct {
    Id              int
    Login           string    `orm:"size(64);unique" valid:"Required"`
    Password        string    `orm:"size(64)"        valid:"Required;MinSize(3)"`
    Rights          int
}

func (user *User) TableName() string {
    return "users"
}

type CatalogsTreePath struct {
    Id              int
    Ctpid           int
    Name            string
    Name_i18n       bool
}

func (catalogstreepath *CatalogsTreePath) TableName() string {
    return "catalogstreepath"
}

type Catalogs struct {
    Id              int
    Cid             int
    Ancestor        int
    Descendant      int
}

func (catalogs *Catalogs) TableName() string {
    return "catalogs"
}

type Equipment struct {
    Id              int
    Description     string
    Equip_id        int
    Equip_type      string
    Image           string
    Is_gift         bool
    Is_premium      bool
    Level           int
    Name            string
    Nation          string
    Price           int
    Short_name      string
    Small_image     string
    Type            string
}

func (equipment *Equipment) TableName() string {
    return "equipments"
}

type Tank struct {
    Id              int
    Equip_id        int
    TankCharacteristics
}

func (tank *Tank) TableName() string {
    return "tanks"
}

type Warplane struct {
    Id              int
    Equip_id        int
    WarplaneCharacteristics
}

func (warplane *Warplane) TableName() string {
    return "warplanes"
}

type Purchase struct {
    Id          int
    Sum         int64
    Count       int
    Date        time.Time
}

func (p *Purchase) TableName() string {
    return "purchases"
}

type Good struct {
    Id          int
    Equip_id    int
    Price       int
    Name        string
    Count       int
    Purchase_id int
}

func (g *Good) TableName() string {
    return "goods"
}

type Level struct {
    Id              int
    Value           string
    Level           int
}

func (t *Level) TableName() string {
    return "levels"
}

type Nation struct {
    Id              int
    Name            string
    Name_i18n       string
}

func (t *Nation) TableName() string {
    return "nations"
}

type Type struct {
    Id              int
    Name            string
    Name_catalog    string
}

func (t *Type) TableName() string {
    return "types"
}
