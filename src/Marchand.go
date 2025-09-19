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

func Merchantm(player *Character, gold *int) {
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

	for {
		fmt.Println("\n=== Marchand : Jerry ===")
		fmt.Println("0. Retour au menu")
		for i := 1; i <= len(items); i++ {
			fmt.Printf("%d. %s (%d pièces)\n", i, items[i].Name, items[i].Price)
		}
		fmt.Printf("Vous avez %d pièces\n", *gold)

		var choix int
		fmt.Print("Choisissez un item : ")
		fmt.Scan(&choix)

		if choix == 0 {
			fmt.Println("Vous quittez le marchand Jerry.")
			return
		}

		item, exists := items[choix]
		if !exists {
			fmt.Println("Choix invalide, réessayez.")
			continue
		}

		if item.Price > *gold {
			fmt.Println("Vous n'avez pas assez de pièces.")
			continue
		}

		*gold -= item.Price
		player.Inv = append(player.Inv, item.Name)
		fmt.Printf("Vous avez acheté : %s\n", item.Name)
		fmt.Printf("Il vous reste %d pièces.\n", *gold)
	}
}
