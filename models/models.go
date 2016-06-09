package models

type TankCharacteristics struct {
    Description     string  `type:"string" key:"Описание" index:"0" orm:"description"`
    Is_premium      bool    `type:"bool" key:"Премиум" index:"5" orm:"is_premium"`
    Level           int     `type:"int" key:"Уровень" index:"3" orm:"level"`
    Name            string  `type:"string" key:"Название" index:"1" orm:"name"`
    Nation          string  `type:"string" key:"Нация" index:"2" orm:"nation"`
    Price           int     `type:"int" key:"Стоимость" index:"6" orm:"price"`
    Type            string  `type:"string" key:"Тип" index:"4" orm:"type"`
    Weight          int     `type:"int" key:"Вес" index:"8" orm:"weight"`
    Max_weight      int     `type:"int" key:"Максимальный вес" index:"9" orm:"max_weight"`
    Armor           string  `type:"string" key:"Бронирование" index:"10" orm:"armor"`
    Hp              int     `type:"int" key:"Очки прочности" index:"7" orm:"hp"`
    Speed_forward   int     `type:"int" key:"Скорость вперёд" index:"11" orm:"speed_forward"`
    Speed_backward  int     `type:"int" key:"Скорость назад" index:"12" orm:"speed_backward"`
}

type WarplaneCharacteristics struct {
    Description     string  `type:"string" key:"Описание" index:"0" orm:"description"`
    Is_premium      bool    `type:"bool" key:"Премиум" index:"5" orm:"is_premium"`
    Level           int     `type:"int" key:"Уровень" index:"3" orm:"level"`
    Name            string  `type:"string" key:"Название" index:"1" orm:"name"`
    Nation          string  `type:"string" key:"Нация" index:"2" orm:"nation"`
    Price           int     `type:"int" key:"Стоимость" index:"6" orm:"price"`
    Type            string  `type:"string" key:"Тип" index:"4" orm:"type"`
    Weight          int     `type:"int" key:"Вес" index:"8" orm:"weight"`
    Hp              int     `type:"int" key:"Очки прочности" index:"7" orm:"hp"`
    Speed_ground    int     `type:"int" key:"Скорость у поверхности земли" index:"9" orm:"speed_ground"`
    Maneuverability int     `type:"int" key:"Манёвренность" index:"10" orm:"maneuverability"`
    Max_speed       int     `type:"int" key:"Максимальная скорость" index:"11" orm:"max_speed"`
    Stall_speed     int     `type:"int" key:"Скорость сваливания" index:"12" orm:"stall_speed"`
    Optimal_height  int     `type:"int" key:"Оптимальный вес" index:"13" orm:"optimal_height"`
    Roll_maneuver   int     `type:"int" key:"Скорость вращения" index:"14" orm:"roll_maneuver"`
    Dive_speed      int     `type:"int" key:"Скорость пикирования" index:"15" orm:"dive_speed"`
    Opt_maneuver_speed int  `type:"int" key:"Оптимальная скорость маневрирования" index:"16" orm:"opt_maneuver_speed"`
}

type Fields struct {
    Id          string
    Name        string
    Price       string
    Nation      string
}

type EquipHomePage struct {
    Equip_id    int
    Name        string
    Price       int
    Small_image string
}

type EquipHomePageIsCount struct {
    EquipHomePage
    IsCount     bool
}

type EquipInTable struct {
    Equip_id    int
    Name        string
    Price       int
    Small_image string
    Nation      string
}

type ExtEquipInTable struct {
    EquipInTable
    Count       int
}

type EquipInTableIsCount struct {
    EquipInTable
    IsCount     bool
}

type BaseEquip struct {
    Equip_id    int
    Name        string
    Price       int
}

type Header struct {
    Name        string
    Namei18n    string
    Type        string
    Pattern     string
    IsTextarea  bool
}
