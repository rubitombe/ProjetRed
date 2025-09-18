package ProjetRed

import (
	"fmt"
)

var inventory []string

func addInventory(item string) {
	inventory = append(inventory, item)
	fmt.Printf("Vous avez obtenu : %s\n", item)
}

func removeInventory(item string) {
	for i, v := range inventory {
		if v == item {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Printf("%s a été retiré de l’inventaire.\n", item)
			return
		}
	}
	fmt.Println("Item introuvable dans l’inventaire.")
}

func showInventory() {
	fmt.Println("=== Inventaire ===")
	if len(inventory) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	for i, v := range inventory {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func merchant() {
	fmt.Println("=== Marchand : Jerry ===")
	fmt.Println("Jerry : Bonjour aventurier ! Voici ce que je propose :")
	fmt.Println("1. Potion de vie (gratuit)")
	fmt.Println("0. Retour")

	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		addInventory("Potion de vie")
		fmt.Println("Jerry : Merci pour ton achat, reviens quand tu veux !")
	case 0:
		fmt.Println("Vous quittez le marchand Jerry.")
	default:
		fmt.Println("Choix invalide.")
	}
}

func mainMenu() {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Marchand Jerry")
		fmt.Println("2. Voir l’inventaire")
		fmt.Println("0. Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			merchant()
		case 2:
			showInventory()
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func main() {
	mainMenu()
}
