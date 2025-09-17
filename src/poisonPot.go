package ProjetRed

import "fmt"

func poisonPot(c *Character, pot string) {

	for i, item := range c.Inv {
		if item == "Poison's Potion" {
			c.PointOfLifeCurrent -= 15
			if c.PointOfLifeCurrent < 0 {
				c.PointOfLifeCurrent = 0
			}
			c.Inv = append(c.Inv[:i], c.Inv[i+1:]...)
			fmt.Println("Vous avez utilisé une potion de poison. Votre vie actuelle est maintenant de", c.PointOfLifeCurrent)
			return
		}
		if !(item == "Poison's Potion") && i == len(c.Inv)-1 {
			fmt.Println("Vous n'avez pas de potion de poison dans votre inventaire.")
			return
		}
	}
}
