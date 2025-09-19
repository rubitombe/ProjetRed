package ProjetRed

import ("fmt
")

func Forgeron(player *Character) {

	equipements := map[int]string{
		1: "Chapeau de paille",
		2: "Robe dior",
		3: "Bottes à talons",
	}

	for {
		fmt.Println("\n=== Forgeron ===")
		fmt.Println("0. Retour au menu principal")
		for i := 1; i <= len(equipements); i++ {
			fmt.Printf("%d. %s (5 pièces d'or)\n", i, equipements[i])
		}
		fmt.Printf("Vous avez %d pièces\n", player.Gold)

		var choix int
		fmt.Print("Que voulez-vous fabriquer aujourd'hui ? ")
		_, err := fmt.Scanln(&choix)
		if err != nil {
			fmt.Println("Erreur, réessayez.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if choix == 0 {
			fmt.Println("Vous quittez le forgeron, à bientot !")
			return
		}

		objet, exists := equipements[choix]
		if !exists {
			fmt.Println("Choix invalide, réessayez.")
			continue
		}

		if player.Gold < 5 {
			fmt.Println("Vous avez besoin de plus de pièces pour fabriquer cet objet.")
			continue
		}

		player.Gold -= 5
		player.Inv = append(player.Inv, objet)
		fmt.Printf("Vous avez fabriqué : %s\n", objet)
		fmt.Printf("Il vous reste %d pièces.\n", player.Gold)
	}
}








