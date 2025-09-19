package main

import (
	ProjetRed "ProjetRed/src"
	"fmt"
)

var player *ProjetRed.Character
var inventory []string

func main() {

	player := ProjetRed.InitCharacter(
		"Veyra",
		"Fairy",
		0,
		100,
		50,
		[]string{"Life's Potion", "Knife", "Death's Potion"},
		[]string{"Magic Locs", "Iced Attieke", "Cataclysm"},
	)

	fmt.Println("Nom :", player.Name)
	fmt.Println("Classe :", player.Class)
	fmt.Println("Niveau :", player.Level)
	fmt.Println("Vie max :", player.PointOfLifeMax)
	fmt.Println("Vie actuelle :", player.PointOfLifeCurrent)
	fmt.Println("Inventaire :", player.Inv)

	ProjetRed.Menu(player)
}
