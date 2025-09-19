package ProjetRed

import (
	"fmt"
)

func CharacterCreation(name string) (*Character, bool) {

	nameRunes := []rune(name)
	if len(nameRunes) == 0 {
		fmt.Println("Le nom ne doit pas être vide.")
		return nil, false
	}
	if len(nameRunes) < 4 || len(nameRunes) > 15 {
		fmt.Println("Le nom doit contenir entre 4 et 15 lettres")
		return nil, false
	}
	if nameRunes[0] < 'A' || nameRunes[0] > 'Z' {
		fmt.Println("Le nom doit commencer par une lettre majuscule.")
		return nil, false
	}
	if nameRunes[1] < 'a' || nameRunes[1] > 'z' {
		fmt.Println("Le nom doit contenir des lettres minuscules après la première lettre.")
		return nil, false
	}
	for _, c := range nameRunes {
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			fmt.Println("Le nom ne doit contenir que des lettres.")
			return nil, false
		}
	}

	fmt.Println("Choisissez une classe :")
	fmt.Println("1. Fée (100 PV)")
	fmt.Println("2. Nain (70 PV)")
	fmt.Println("3. Gobelin (80 PV)")
	fmt.Println("4. Loup(120 PV)")
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
		Spell:              "Coup de tête",
	}, true
}
