package main

type character struct {
	name               string
	class              string
	level              int
	pointOfLifeMax     int
	pointOfLifeCurrent int
	inv                []object
}

func main() {
	c1 := character{"Veyra", "Fairy", 0, 100, 50, []object{"Life's Potion", "Knife", "Death's Potion"}}
	println(c1.name)
}
