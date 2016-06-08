package models

type TankCharacteristics struct {
    Description     string  `type:"string" key:"Описание" index:"0"`
    Is_premium      bool    `type:"bool" key:"Премиум" index:"5"`
    Level           int     `type:"int" key:"Уровень" index:"3"`
    Name            string  `type:"string" key:"Название" index:"1"`
    Nation          string  `type:"string" key:"Нация" index:"2"`
    Price           int     `type:"int" key:"Стоимость" index:"6"`
    Type            string  `type:"string" key:"Тип" index:"4"`
    Weight          int     `type:"int" key:"Вес" index:"8"`
    Max_weight      int     `type:"int" key:"Максимальный вес" index:"9"`
    Armor           string  `type:"string" key:"Бронирование" index:"10"`
    Hp              int     `type:"int" key:"Очки прочности" index:"7"`
    Speed_forward   int     `type:"int" key:"Скорость вперёд" index:"11"`
    Speed_backward  int     `type:"int" key:"Скорость назад" index:"12"`
}

type WarplaneCharacteristics struct {
    Description     string  `type:"string" key:"Описание" index:"0"`
    Is_premium      bool    `type:"bool" key:"Премиум" index:"5"`
    Level           int     `type:"int" key:"Уровень" index:"3"`
    Name            string  `type:"string" key:"Название" index:"1"`
    Nation          string  `type:"string" key:"Нация" index:"2"`
    Price           int     `type:"int" key:"Стоимость" index:"6"`
    Type            string  `type:"string" key:"Тип" index:"4"`
    Weight          int     `type:"int" key:"Вес" index:"8"`
    Hp              int     `type:"int" key:"Очки прочности" index:"7"`
    Speed_ground    int     `type:"int" key:"Скорость у поверхности земли" index:"9"`
    Maneuverability int     `type:"int" key:"Манёвренность" index:"10"`
    Max_speed       int     `type:"int" key:"Максимальная скорость" index:"11"`
    Stall_speed     int     `type:"int" key:"Скорость сваливания" index:"12"`
    Optimal_height  int     `type:"int" key:"Оптимальный вес" index:"13"`
    Roll_maneuver   int     `type:"int" key:"Скорость вращения" index:"14"`
    Dive_speed      int     `type:"int" key:"Скорость пикирования" index:"15"`
    Opt_maneuver_speed int  `type:"int" key:"Оптимальная скорость маневрирования" index:"16"`
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
