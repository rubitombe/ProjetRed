package ProjetRed

import (
	"fmt"
)

func isDead(c *Character) {

	if c.PointOfLifeCurrent <= 0 {
		fmt.Println("Character is dead")
		c.PointOfLifeCurrent = 30
	} else {
		fmt.Println("Character is still alive")
	}
}
