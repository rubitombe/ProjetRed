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

var Inventory = map[string]Item{
	"Death's Potion": {Quantite: 1, Degats: 15, Durabilite: 100},
	"Life's Potion":  {Quantite: 2, Effet: "soin", Valeur: 20},
	"Knife":          {Quantite: 1, Degats: 10, Durabilite: 80},
}

func AccessInventory(inv map[string]Item, player *Character) {
	fmt.Println("=== Inventaire du joueur ===")
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

	// Demande au joueur quel item utiliser
	fmt.Println("Quel objet voulez-vous utiliser ? (Tapez le nom exact ou '0' pour revenir)")
	var choix string
	fmt.Scan(&choix)

	if choix == "0" {
		return // retourne au menu principal
	}

	useItem(inv, choix, player) // ✅ Appelle la fonction useItem
}

func checkInventory(inventory []string) bool {
	if len(inventory) >= 10 {
		fmt.Printf("⚠️ Inventaire plein (%d/10).\n", len(inventory))
		return false
	}
	return true
}

func addInventory(item string, inventory *[]string) {
	if len(*inventory) >= 10 {
		return
	}
	*inventory = append(*inventory, item)
	fmt.Printf("✅ Vous avez obtenu : %s\n", item)
}

func removeInventory(item string, inventory *[]string) {
	for i, v := range *inventory {
		if v == item {
			*inventory = append((*inventory)[:i], (*inventory)[i+1:]...)
			fmt.Printf("🗑️ %s retiré de l’inventaire.\n", item)
			return
		}
	}
	fmt.Println("❌ Item introuvable dans l’inventaire.")
}

func showInventory(inv map[string]Item, player *Character) {
	fmt.Println("\n🎒 === Inventaire de", player.Name, "=== 🎒")
	if len(inv) == 0 {
		fmt.Println("Inventaire vide...")
		return
	}

	i := 1
	for nom := range inv {
		fmt.Printf(" %d. %s\n", i, nom)
		i++
	}

	var choix string
	fmt.Println("\n👉 Quel objet voulez-vous utiliser ?")
	fmt.Scan(&choix)

	useItem(inv, choix, player)
}

func useItem(inv map[string]Item, itemName string, player *Character) {
	item, exists := inv[itemName]
	if !exists {
		fmt.Println("❌ Item introuvable dans l’inventaire.")
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
		spellBook(player, "Boule de Feu") // ✅ le sort est appris si pas déjà connu

	default:
		fmt.Printf("Vous utilisez %s.\n", itemName)
	}

	// Diminue la quantité et supprime si zéro
	item.Quantite--
	if item.Quantite <= 0 {
		delete(inv, itemName)
	} else {
		inv[itemName] = item
	}
}
