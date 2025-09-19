package ProjetRed

import (
	"fmt"
)

func Merchant(player *Character, someInt *int) {
	fmt.Println("=== Marchand : Jerry ===")
	fmt.Println("Jerry : Bonjour aventurier ! Voici ce que je propose :")
	fmt.Println("1. Potion de vie (gratuit)")
	fmt.Println("2. Livre de Sort : Boule de Feu")
	fmt.Println("0. Retour")

	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		AddInventory("Potion de vie")
		fmt.Println("Jerry : Merci pour ton achat, reviens quand tu veux !")
	case 2:
		AddInventory("📖 Livre de Sort : Boule de Feu")
		fmt.Println("Jerry : Tu viens d'acquérir un nouveau sort !")
	case 0:
		fmt.Println("Vous quittez le marchand Jerry.")
	default:
		fmt.Println("Choix invalide.")
	}
}

// AddInventory ajoute un item à l'inventaire global
func AddInventory(item string) {
	inventory = append(inventory, item)
	fmt.Printf("✅ Vous avez obtenu : %s\n", item)
}

// RemoveInventory retire un item de l'inventaire global
func RemoveInventory(item string) {
	for i, v := range inventory {
		if v == item {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Printf("🗑️ %s retiré de l’inventaire.\n", item)
			return
		}
	}
	fmt.Println("❌ Item introuvable dans l’inventaire.")
}

// ShowInventory affiche l’inventaire global
func ShowInventory() {
	fmt.Println("=== Inventaire ===")
	if len(inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	for i, v := range inventory {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

// Inventaire global
var inventory []string
