package main

import "fmt"

func appends() {
	var nome []string

	var nome2 = append(nome, "joao")
	nome2 = append(nome2, "Maria")

	fmt.Println(nome)
	fmt.Println(nome2)
}
