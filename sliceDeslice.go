package main

import "fmt"

func sliceDeSlice() {
	tabuleiro := [][]string{
		[]string{"x", "O"},
		[]string{"x", "x"},
		[]string{"o", "O"},
	}

	fmt.Println(tabuleiro)
}
