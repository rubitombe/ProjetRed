package ProjetRed

import (
	"fmt"
)

var inventory []string

func AddInventorym(item string) {
	inventory = append(inventory, item)
	fmt.Printf("Vous avez obtenu : %s\n", item)
}

func RemoveInventorym(item string) {
	for i, v := range inventory {
		if v == item {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Printf("%s a été retiré de l’inventaire.\n", item)
			return
		}
	}
	fmt.Println("Item introuvable dans l’inventaire.")
}

func ShowInventorym() {
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
		6: {"Fourrure ours polaire", 7},
		7: {"Cuir de dragon", 3},
		8: {"Plume de faucon", 1},
	}

	fmt.Println("=== Marchand : Jerry ===")
	fmt.Println("Jerry : Bonjour aventurier ! Voici ce que je propose :")
	fmt.Println("0. Retour")
	for i := 1; i <= len(items); i++ {
		fmt.Printf("%d. %s (%d pièces)\n", i, items[i].Name, items[i].Price)
		if items[i].Price > *gold {
			fmt.Printf("   (Vous avez %d pièces.)\n", *gold)
		} else {
			fmt.Printf("   (Prix : %d pièces)\n", items[i].Price)
		}
	}
	var choix int
	fmt.Println("Choisissez un item :")
	fmt.Scan(&choix)
	if items[choix].Price > *gold {
		fmt.Println("   (Vous n'avez pas assez de pièces pour cet article.)")
		return
	}
	if choix == 0 {
		fmt.Println("Vous quittez le marchand Jerry.")
		return
	}
	if choix < 0 || choix > len(items) {
		fmt.Println("   (Choix invalide.)")
		return
	}
	if items[choix].Price <= *gold {
		fmt.Println("   (Vous pouvez acheter cet article.)")
	}
	item, exists := items[choix]
	if !exists {
		fmt.Println("   (Cet article n'existe pas.)")
		return
	}
	*gold -= item.Price
	inventory = append(inventory, items[choix].Name)
	fmt.Printf("Vous avez acheté : %s\n", items[choix].Name)
	fmt.Printf("Il vous reste %d pièces.\n", *gold)
	fmt.Println("Jerry : Merci pour ton achat, reviens quand tu veux !")
}

func MainMenu() {
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Marchand Jerry")
		fmt.Println("2. Voir inventaire")
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
			ShowInventorym()
		case 0:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
