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
	Money              string
	Spell              []string
	SpellEquipped      string
}

func InitCharacter(name string, class string, level int, pointOfLifeMax int, pointOfLifeCurrent int, inv []string) *Character {

	return &Character{
		Name:               name,
		Class:              class,
		Level:              level,
		PointOfLifeMax:     pointOfLifeMax,
		PointOfLifeCurrent: pointOfLifeCurrent,
		Inv:                inv,
	}
}
