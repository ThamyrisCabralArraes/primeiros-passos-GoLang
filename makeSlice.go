package main

import "fmt"

func makeSlice() {

	numeros := make([]int, 0, 5)
	fmt.Println(numeros[0:3])

}
