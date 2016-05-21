package query

import "regexp"



func MakeTanksQuery(s string) string {

}

func MakeWarplanesQuery(s string) string {
    r_classes, _ := regexp.Compile("classes/(.*)")
    r_levels, _ := regexp.Compile("levels/(\d*)")
    ans := "tank_id, t.name, t.price, t.nation FROM tanks t"
    if r_classes.MatchString(s) {
        ans += fmt.Sprintf("inner join types on t.type = types.name WHERE types.cat_name = %s",
            r_classes.FindString(s))
    } else if r_levels.MatchString(s) {
        ans += fmt.Sprintf("inner join types on t.type = types.name WHERE types.cat_name = %s",
            r_classes.FindString(s))
    }
}

func MakeQuery(s string) string {
    // "/catalog/tanks/light_tanks"
    r_tanks, _ := regexp.Compile("tanks")
    r_warplanes, _ := regext.Compile("warplanes")
    ans := "SELECT "
    if r_tanks.MatchString(s) {
        return ans + MakeTanksQuery(s)
    } else if r_warplanes.MatchString(s) {
        return ans + MakeWarplanesQuery(s)
    } else {
        return "SELECT tank_id, name, price, nation FROM tanks"
    }
}
