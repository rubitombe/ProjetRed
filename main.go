package main

import (
	ProjetRed "ProjetRed/src"
	"fmt"
)

var player *ProjetRed.Character
var inventory []string

func main() {
	fmt.Println("Bienvenue dans le jeu Projet Red!")
	fmt.Println("Veuillez entrer le nom de votre personnage :")
	var nameRunes string
	fmt.Scanln(&nameRunes)

	c1 := ProjetRed.CharacterCreation(nameRunes)
	fmt.Println("Nom :", c1.Name)
	fmt.Println("Classe :", c1.Class)
	fmt.Println("Niveau :", c1.Level)
	fmt.Println("Vie max :", c1.PointOfLifeMax)
	fmt.Println("Vie actuelle :", c1.PointOfLifeCurrent)
	fmt.Println("Inventaire :", c1.Inv)

	fmt.Println("Nom :", player.Name)
	fmt.Println("Classe :", player.Class)
	fmt.Println("Niveau :", player.Level)
	fmt.Println("Vie max :", player.PointOfLifeMax)
	fmt.Println("Vie actuelle :", player.PointOfLifeCurrent)
	fmt.Println("Inventaire :", player.Inv)

	fmt.Println("Sort :", c1.Spell)
}

// func main() {

// 	c1 := ProjetRed.InitCharacter(
// 		"Veyra",
// 		"Fairy",
// 		0,
// 		100,
// 		50,
// 		[]string{"Life's Potion", "Knife", "Death's Potion"},
// 	)
// 	fmt.Println("Nom :", c1.Name)
// 	fmt.Println("Classe :", c1.Class)
// 	fmt.Println("Niveau :", c1.Level)
// 	fmt.Println("Vie max :", c1.PointOfLifeMax)
// 	fmt.Println("Vie actuelle :", c1.PointOfLifeCurrent)
// 	fmt.Println("Inventaire :", c1.Inv)

// }
