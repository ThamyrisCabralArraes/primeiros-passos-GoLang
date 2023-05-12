package main

import "fmt"

func while() {
	var soma int = 2

	for soma < 500 {
		soma += soma
	}
	fmt.Println("soma", soma)
}
