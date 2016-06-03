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

type TankCharacteristics struct {
    Description     string
    Is_premium      bool
    Level           int
    Name            string
    Nation          string
    Price           int
    Type            string
    Weight          int
    Max_weight      int
    Armor           string
    Hp              int
    Speed_forward   int
    Speed_backward  int
}

type Tank struct {
    Id              int
    Equip_id        int
    TankCharacteristics
}

func (tank *Tank) TableName() string {
    return "tanks"
}

type WarplaneCharacteristics struct {
    Description     string
    Is_premium      bool
    Level           int
    Name            string
    Nation          string
    Price           int
    Type            string
    Weight          int
    Hp              int
    Speed_ground    int
    Maneuverability int
    Max_speed       int
    Stall_speed     int
    Optimal_height  int
    Roll_maneuver   int
    Dive_speed      int
    Opt_maneuver_speed int
}

type Warplane struct {
    Id              int
    Equip_id        int
    WarplaneCharacteristics
}

func (warplane *Warplane) TableName() string {
    return "warplanes"
}

type Type struct {
    Id              int
    Name            string
    Name_catalog    string
}

func (t *Type) TableName() string {
    return "types"
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

type Fields struct {
    Id          string
    Name        string
    Price       string
    Nation      string
}

type EquipInTable struct {
    Equip_id    int
    Small_image string
    Name        string
    Nation      string
    Price       int
}

type ExtEquipInTable struct {
    EquipInTable
    Count       int
}

type Purchase struct {
    Id          int
    Sum         int
    Count       int
    Date        time.Time
}

func (p *Purchase) TableName() string {
    return "purchases"
}

type Good struct {
    Id          int
    Cost        int
    Count       int
    Purchase_id int
}

func (g *Good) TableName() string {
    return "goods"
}
