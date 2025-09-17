package ProjetRed

import (
	"fmt"
)

type Item struct {
	Quantite   int
	Degats     int
	Durabilite int
	Effet      string
	Valeur     int
}

var inventory = map[string]Item{
	"Death's Potion": {Quantite: 1, Degats: 15, Durabilite: 100},
	"Life's Potion":  {Quantite: 5, Effet: "soin", Valeur: 20},
	"Knife":          {Quantite: 1, Degats: 10, Durabilite: 80},
}

func accessInventory(inv map[string]Item) {
	fmt.Println("Inventaire du joueur")
	for nom, item := range inv {
		fmt.Printf("\nObjet : %s\n", nom)
		fmt.Printf("  Quantité   : %d\n", item.Quantite)
		if item.Degats > 0 {
			fmt.Printf("  Dégâts     : %d\n", item.Degats)
		}
		if item.Durabilite > 0 {
			fmt.Printf("  Durabilité : %d\n", item.Durabilite)
		}
		if item.Effet != "" {
			fmt.Printf("  Effet      : %s\n", item.Effet)
		}
		if item.Valeur > 0 {
			fmt.Printf("  Valeur     : %d\n", item.Valeur)
		}
	}
	fmt.Println()
}
