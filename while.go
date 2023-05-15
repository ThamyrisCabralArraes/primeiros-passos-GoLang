package main

import "fmt"

func while() {
	var soma int = 2

	for soma < 500 {
		soma += soma
	}
	fmt.Println("soma", soma)

	var nome string = "Rosa"

	// switch nao precisa de break
	switch nome {
	case "Rosa":
		fmt.Println("O nome dela é Rosa")
	case "Joao":
		fmt.Println("O nome dela é Joao")
	default:
		fmt.Println("Nao sei quem é")
	}
}
