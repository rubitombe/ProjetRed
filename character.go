package ProjetRed

type object struct {
	name     string
	quantity int
}

type character struct {
	name               string
	class              string
	level              int
	pointOfLifeMax     int
	pointOfLifeCurrent int
	inv                []string
}
