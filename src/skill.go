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

func UseItem(itemName string, player *Character) {
	item, ok := Inventory[itemName]
	if !ok {
		fmt.Println("❌ Item introuvable.")
		return
	}

	switch itemName {
	case "Life's Potion":
		player.PointOfLifeCurrent += item.Valeur
		if player.PointOfLifeCurrent > player.PointOfLifeMax {
			player.PointOfLifeCurrent = player.PointOfLifeMax
		}
		fmt.Printf("✅ Vous utilisez %s et récupérez %d points de vie.\n", itemName, item.Valeur)

	case "Livre de Sort : Boule de Feu":
		spellBook(player, "Boule de Feu") // ✅ ici, spellBook est défini dans le même package

	default:
		fmt.Printf("Vous utilisez %s.\n", itemName)
	}

	item.Quantite--
	if item.Quantite <= 0 {
		delete(Inventory, itemName)
	} else {
		Inventory[itemName] = item
	}
}
