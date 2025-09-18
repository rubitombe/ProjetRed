package ProjetRed

import (
	"fmt"
)

func TakePot(c *Character, pot string) {

	for i, item := range c.Inv {
		if item == "Life's Potion" {
			c.PointOfLifeCurrent += 20
			if c.PointOfLifeCurrent > c.PointOfLifeMax {
				c.PointOfLifeCurrent = c.PointOfLifeMax
			}
			c.Inv = append(c.Inv[:i], c.Inv[i+1:]...)
			fmt.Println("Vous avez utilisé une potion de vie. Votre vie actuelle est maintenant de", c.PointOfLifeCurrent)
			return
		}
		if !(item == "Life's Potion") && i == len(c.Inv)-1 {
			fmt.Println("Vous n'avez pas de potion de vie dans votre inventaire.")
			return
		}
	}
}
