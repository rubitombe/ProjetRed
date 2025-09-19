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

func Menu(c1 *Character) {
	var choix int

	for {

		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Accéder au marchand")
		fmt.Println("4. Quitter")
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scan(&choix)
		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide !")

			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch choix {
		case 1:
			DisplayInfo(c1)
		case 2:
			accessInventoryMenu()
		case 3:
			var someInt int = 0
			merchant(c1, &someInt)
		case 4:
			fmt.Println("Au revoir")
			return
		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}
