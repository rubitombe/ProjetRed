package main

import "fmt"

func main() {
	c1 := Character{"Veyra", "Fairy", 0, 100, 50, []string{{"Life's Potion", 1}, {"Knife", 1}, {"Death's Potion", 1}}}
	return &c1
	fmt.Println(c1.name)

}
