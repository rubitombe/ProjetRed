package main

type character struct {
	name               string
	class              string
	level              int
	pointOfLifeMax     int
	pointOfLifeCurrent int
	inv                []string
}

func main() {
	c1 := character{"Veyra", "Fairy", 0, 100, 50, []string{"Life's Potion", "Knife", "Death's Potion"}}
	println(c1.name)
}
