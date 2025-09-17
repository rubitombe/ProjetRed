package ProjetRed

import (
	"fmt"
)

func DisplayInfo(c *Character) {
	fmt.Println("\n=== Informations du personnage ===")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Class)
	fmt.Println("Niveau :", c.Level)
	fmt.Printf("Points de vie : %d/%d\n", c.PointOfLifeCurrent, c.PointOfLifeMax)
	fmt.Println("Inventaire :")
	if len(c.Inv) == 0 {
		fmt.Println("  (vide)")
	} else {
		for i, item := range c.Inv {
			fmt.Printf("  %d. %s\n", i+1, item)
		}
	}
	fmt.Println("=================================")
}
