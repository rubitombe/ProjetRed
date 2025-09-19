package ProjetRed

type Object struct {
	Name     string
	Quantity int
}

type Character struct {
	Name               string
	Class              string
	Level              int
	PointOfLifeMax     int
	PointOfLifeCurrent int
	Inv                []string
	Skills             []string
	Spell              []string
	SpellEquipped      string
}

func InitCharacter(name string, class string, level int, hpMax int, hpCurrent int, inv []string, Skills []string) *Character {

	c := &Character{
		Name:               name,
		Class:              class,
		Level:              level,
		PointOfLifeMax:     hpMax,
		PointOfLifeCurrent: hpCurrent,
		Inv:                inv,
	}
	c.Skills = append(c.Skills, "👊 Coup de poing")
	return c

}
