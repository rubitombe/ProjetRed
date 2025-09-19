package ProjetRed

import (
	"fmt"
)

func accessInventoryMenu() {

	fmt.Println("Inventaire :")
	fmt.Println("Knife")
	fmt.Println("Potion de vie")
	fmt.Println("Potion de mort ")
}

func Menu(player *Character) {
	var choix int

	for {

		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Accéder au forgeron")
		fmt.Println("5. Quitter")

		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Veuillez entrer un chiffre valide !")

			var discard string
			fmt.Scanln(&discard)
			return
		}

		switch choix {
		case 1:
			DisplayInfo(player)
		case 2:
			accessInventoryMenu()
		case 3:
			var someInt int = 0
			Merchant(player, &someInt)
		case 4:
			forgeron()
			var recipes = map[string]map[string]int{
				"Chapeau de paille": {
					"Plume de faucon": 1,
					"Peau de serpent": 1,
				},
				"Robe Dior": {
					"Plume de faucon": 1,
					"Fourrure d'ours": 1,
				},
				"Bottes à talons": {
					"Cuir de dragon":  1,
					"Peau de serpent": 1,
				},
			}
		case 5:
			fmt.Println("Au revoir et bonne continuation !")
			return
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}
