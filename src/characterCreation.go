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
	return &Character{Name: name}, true

}
