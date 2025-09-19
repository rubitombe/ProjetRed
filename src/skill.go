package ProjetRed

import (
	"fmt"
)

func showSkills(player *Character) {
	fmt.Println("\n✨ === Grimoire de", player.Name, "=== ✨")
	if len(player.Skills) == 0 {
		fmt.Println("Aucune compétence connue.")
		return
	}
	for i, skill := range player.Skills {
		fmt.Printf(" %d. %s\n", i+1, skill)
	}
}

func spellBook(player *Character, spell string) {
	for _, s := range player.Skills {
		if s == spell {
			fmt.Println("⚠️ Vous connaissez déjà ce sort :", spell)
			return
		}
	}
	player.Skills = append(player.Skills, spell)
	fmt.Println("🔥 Nouveau sort appris :", spell)
}
