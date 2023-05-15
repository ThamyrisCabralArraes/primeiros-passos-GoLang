package main

import "fmt"

func array() {
	var numeros [7]int
	numeros[0] = 1
	numeros[1] = 1
	numeros[2] = 2
	numeros[3] = 1
	numeros[4] = 4
	numeros[5] = 1
	numeros[6] = 2
	fmt.Println(numeros)

	var numeros2 = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numeros2)

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])

	}
}
