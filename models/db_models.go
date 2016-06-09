package models

import "time"

type User struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Login           string      `orm:"size(64);unique" valid:"Required" type:"string" index:"1" namei18n:"Логин" pattern:"[a-zA-Z0-9_]+"`
    Password        string      `orm:"size(64)" valid:"Required;MinSize(3)" type:"string" index:"2" namei18n:"Пароль" pattern:"[a-zA-Z0-9_]+"`
    Rights          int         `type:"int" index:"3" namei18n:"Уровень прав" pattern:"[0-3]"`
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
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Description     string      `type:"string" index:"1" namei18n:"Описание" pattern:".*" formtag:"textarea"`
    Equip_id        int         `type:"int" index:"2" namei18n:"Код товара" pattern:"[0-9]+"`
    Equip_type      string      `type:"string" index:"3" namei18n:"Тип товара" pattern:".*"`
    Image           string      `type:"string" index:"4" namei18n:"Большое изображение" pattern:".*"`
    Is_gift         bool        `type:"bool" index:"5" namei18n:"Подарочный" pattern:"true|false"`
    Is_premium      bool        `type:"bool" index:"6" namei18n:"Премиум" pattern:"true|false"`
    Level           int         `type:"int" index:"7" namei18n:"Уровень" pattern:"[0-9]|10" default:"1"`
    Name            string      `type:"string" index:"8" namei18n:"Название" pattern:".*"`
    Nation          string      `type:"string" index:"9" namei18n:"Нация" pattern:".*"`
    Price           int         `type:"int" index:"10" namei18n:"Цена" pattern:"[0-9]+"`
    Short_name      string      `type:"string" index:"11" namei18n:"Сокр. название" pattern:".*"`
    Small_image     string      `type:"string" index:"12" namei18n:"Маленькое изображение" pattern:".*"`
    Type            string      `type:"string" index:"13" namei18n:"Тип" pattern:".+"`
    Have_goods       bool        `type:"bool" index:"14" namei18n:"Наличие товара" pattern:"true|false"`
    Delivery_time    int         `type:"int" index:"15" namei18n:"Срок доставки" pattern:"[0-9]+"`
}

func (equipment *Equipment) TableName() string {
    return "equipments"
}

type Tank struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Equip_id        int         `type:"int" index:"2" namei18n:"Код товара" pattern:"[0-9]+"`
    Description     string      `type:"string" namei18n:"Описание" index:"1" pattern:".*" formtag:"textarea"`
    Is_premium      bool        `type:"bool" namei18n:"Премиум" index:"4" pattern:"true|false"`
    Level           int         `type:"int" namei18n:"Уровень" index:"5" pattern:"[0-9]|10" default:"1"`
    Name            string      `type:"string" namei18n:"Название" index:"6" pattern:".*"`
    Nation          string      `type:"string" namei18n:"Нация" index:"7" pattern:".*"`
    Price           int         `type:"int" namei18n:"Стоимость" index:"8" pattern:"[0-9]+"`
    Type            string      `type:"string" namei18n:"Тип" index:"3" pattern:".*"`
    Weight          int         `type:"int" namei18n:"Вес" index:"9" pattern:"[0-9]+"`
    Max_weight      int         `type:"int" namei18n:"Максимальный вес" index:"10" pattern:"[0-9]+"`
    Armor           string      `type:"string" namei18n:"Бронирование" index:"11" pattern:"[0-9]+\\[0-9]+\\[0-9]+"`
    Hp              int         `type:"int" namei18n:"Очки прочности" index:"12" pattern:"[0-9]+"`
    Speed_forward   int         `type:"int" namei18n:"Скорость вперёд" index:"13" pattern:"[0-9]+"`
    Speed_backward  int         `type:"int" namei18n:"Скорость назад" index:"14" pattern:"[0-9]+"`

}

func (tank *Tank) TableName() string {
    return "tanks"
}

type Warplane struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Equip_id        int         `type:"int" index:"2" namei18n:"Код товара" pattern:"[0-9]+"`
    Description     string      `type:"string" namei18n:"Описание" index:"1" pattern:".*" formtag:"textarea"`
    Is_premium      bool        `type:"bool" namei18n:"Премиум" index:"4" pattern:"true|false"`
    Level           int         `type:"int" namei18n:"Уровень" index:"5" pattern:"[0-9]|10" default:"1"`
    Name            string      `type:"string" namei18n:"Название"index:"6" pattern:".*"`
    Nation          string      `type:"string" namei18n:"Нация" index:"7" pattern:".*"`
    Price           int         `type:"int" namei18n:"Стоимость" index:"8" pattern:"[0-9]+"`
    Type            string      `type:"string" namei18n:"Тип" index:"3" pattern:".*"`
    Weight          int         `type:"int" namei18n:"Вес" index:"9" pattern:"[0-9]+"`
    Hp              int         `type:"int" namei18n:"Очки прочности" index:"10" name_i18n:"Очки прочности" pattern:"[0-9]+"`
    Speed_ground    int         `type:"int" namei18n:"Скорость у поверхности земли" index:"11" pattern:"[0-9]+"`
    Maneuverability int         `type:"int" namei18n:"Манёвренность" index:"12" pattern:"[0-9]+"`
    Max_speed       int         `type:"int" namei18n:"Максимальная скорость" index:"13" pattern:"[0-9]+"`
    Stall_speed     int         `type:"int" namei18n:"Скорость сваливания" index:"14" pattern:"[0-9]+"`
    Optimal_height  int         `type:"int" namei18n:"Оптимальный вес" index:"15" pattern:"[0-9]+"`
    Roll_maneuver   int         `type:"int" namei18n:"Скорость вращения" index:"16" pattern:"[0-9]+"`
    Dive_speed      int         `type:"int" namei18n:"Скорость пикирования" index:"17" pattern:"[0-9]+"`
    Opt_maneuver_speed int      `type:"int" namei18n:"Оптимальная скорость маневрирования" index:"18" pattern:"[0-9]+"`

}

func (warplane *Warplane) TableName() string {
    return "warplanes"
}

type Purchase struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Sum             int64       `type:"int64" index:"1" namei18n:"Сумма" pattern:"[0-9]+"`
    Count           int         `type:"int" index:"2" namei18n:"Количество" pattern:"[0-9]+"`
    Date            time.Time   `type:"time.Time" index:"3" namei18n:"Дата" pattern:".*"`
    Purchaser_id    int         `type:"int" index:"4" namei18n:"ID покупателя" pattern:"[0-9]*"`
    Status          string      `type:"string" index:"5" namei18n:"Статус покупки" pattern:".*"`
}

func (p *Purchase) TableName() string {
    return "purchases"
}

type Good struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Equip_id        int         `type:"int" index:"2" namei18n:"Код товара" pattern:"[0-9]+"`
    Price           int         `type:"int" index:"1" namei18n:"Стоимость" pattern:"[0-9]+"`
    Name            string      `type:"string" index:"3" namei18n:"Название" pattern:".*"`
    Count           int         `type:"int" index:"4" namei18n:"Количество" pattern:"[0-9]+"`
    Purchase_id     int         `type:"int" index:"5" namei18n:"ID покупки" pattern:"[0-9]+"`
    Purchaser_id    int         `type:"int" index:"6" namei18n:"ID покупателя" pattern:"[0-9]*"`
}

func (g *Good) TableName() string {
    return "goods"
}

type Level struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Value           string      `type:"string" index:"1" namei18n:"Значение" pattern:".*"`
    Level           int         `type:"int" index:"2" namei18n:"Уровень" pattern:"[1-9]|10"`
}

func (t *Level) TableName() string {
    return "levels"
}

type Nation struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Name            string      `type:"string" index:"1" namei18n:"Название" pattern:".+"`
    Name_i18n       string      `type:"string" index:"2" namei18n:"Название ру" pattern:".+"`
}

func (t *Nation) TableName() string {
    return "nations"
}

type Type struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Name            string      `type:"string" index:"1" namei18n:"Название" pattern:".+"`
    Name_catalog    string      `type:"string" index:"2" namei18n:"Название в каталоге" pattern:".+"`
}

func (t *Type) TableName() string {
    return "types"
}

type Translate struct {
    Id              int         `type:"int" index:"0" namei18n:"ID" pattern:"[0-9]+"`
    Name            string      `type:"string" index:"1" namei18n:"Название" pattern:".*"`
    Name_i18n       string      `type:"string" index:"2" namei18n:"Региональное название" pattern:".*"`
}

func (t *Translate) TableName() string {
    return "translates"
}
