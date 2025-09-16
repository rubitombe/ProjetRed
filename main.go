package piscine

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
	inv                []object
}

func main() {
	c1 := character{"Veyra", "Fairy", 0, 100, 50, []object{{"Life's Potion", 1}, {"Knife", 1}, {"Death's Potion", 1}}}
	return &c1
	println(c1.name)

}
