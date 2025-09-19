package ProjetRed

import (
	"fmt"
)

func CharacterCreation(name string) *Character {
	var nameRunes []rune

	for {
		if name == "" {
			fmt.Print("Veuillez entrer le nom de votre personnage : ")
			fmt.Scanln(&name)
		}

		nameRunes = []rune(name)

		if len(nameRunes) == 0 {
			fmt.Println("Le nom ne doit pas être vide.")
			name = ""
			continue
		}
		if len(nameRunes) < 4 || len(nameRunes) > 15 {
			fmt.Println("Le nom doit contenir entre 4 et 15 lettres.")
			name = ""
			continue
		}
		if nameRunes[0] < 'A' || nameRunes[0] > 'Z' {
			fmt.Println("Le nom doit commencer par une lettre majuscule.")
			name = ""
			continue
		}
		if len(nameRunes) > 1 && (nameRunes[1] < 'a' || nameRunes[1] > 'z') {
			fmt.Println("Le nom doit contenir des lettres minuscules après la première lettre.")
			name = ""
			continue
		}
		valid := true
		for _, c := range nameRunes {
			if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
				fmt.Println("Le nom ne doit contenir que des lettres.")
				valid = false
				break
			}
		}
		if !valid {
			name = ""
			continue
		}

		break
	}

	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Fée (100 PV)")
	fmt.Println("2. Nain (70 PV)")
	fmt.Println("3. Gobelin (80 PV)")
	fmt.Println("4. Loup (120 PV)")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choice)

	var class string
	var maxLife int

	switch choice {
	case 1:
		class = "Fée"
		maxLife = 100
	case 2:
		class = "Nain"
		maxLife = 70
	case 3:
		class = "Gobelin"
		maxLife = 80
	case 4:
		class = "Loup"
		maxLife = 120
	default:
		fmt.Println("Choix invalide, par défaut vous serez Humain.")
		class = "Humain"
		maxLife = 100
	}

	return &Character{
		Name:               name,
		Class:              class,
		Level:              1,
		PointOfLifeMax:     maxLife,
		PointOfLifeCurrent: maxLife / 2,
		Inv:                []string{},
		Spell:              []string{"Boule d'attieke", "Magic Locs", "Cataclysme"},
		SpellEquipped:      "Coup de tête",
	}
}
