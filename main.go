package main

import "fmt"

func main() {

	c1 := InitCharacter(
		"Veyra",
		"Fairy",
		0,
		100,
		50,
		[]string{"Life's Potion", "Knife", "Death's Potion"},
	)
	fmt.Println("Nom :", c1.Name)
	fmt.Println("Classe :", c1.Class)
	fmt.Println("Niveau :", c1.Level)
	fmt.Println("Vie max :", c1.PointOfLifeMax)
	fmt.Println("Vie actuelle :", c1.PointOfLifeCurrent)
	fmt.Println("Inventaire :", c1.Inv)
}
