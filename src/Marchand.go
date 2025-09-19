package ProjetRed

import (
	"fmt"
)

var inventory []string

func AddInventory(item string) {
	inventory = append(inventory, item)
	fmt.Printf("Vous avez obtenu : %s\n", item)
}

func RemoveInventory(item string) {
	for i, v := range inventory {
		if v == item {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Printf("%s a été retiré de l’inventaire.\n", item)
			return
		}
	}
	fmt.Println("Item introuvable dans l’inventaire.")
}

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

func merchant(_ *Character, gold *int) {
	items := map[int]struct {
		Name  string
		Price int
	}{
		1: {"Potion de vie", 0},
		2: {"Potion de vie", 9},
		3: {"Potion de mort", 12},
		4: {"Livre de sort: Boule d'attieke", 31},
		5: {"Peau de serpent", 11},
		6: {"Fourrure ours polaire", 8},
		7: {"Cuir de dragon", 7},
		8: {"Plume de faucon", 5},
	}

	fmt.Println("=== Marchand : Jerry ===")
	fmt.Println("Jerry : Bonjour aventurier ! Voici ce que je propose :")
	fmt.Println("1. Potion de vie (gratuit)")
	fmt.Println("0. Retour")

	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		AddInventory("Potion de vie")
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
			player := &Character{}
			gold := 50
			merchant(player, &gold)
		case 2:
			ShowInventory()
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
