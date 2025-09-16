package main

func main() {
	c1 := character{"Veyra", "Fairy", 0, 100, 50, []string{{"Life's Potion", 1}, {"Knife", 1}, {"Death's Potion", 1}}}
	return &c1
	fmt.println(c1.name)

}
