package main

import "fmt"

func slice() {
	var numeros2 = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numeros2[2:5])
	fmt.Println(numeros2[2:])
	fmt.Println(numeros2[:5])

	var nomes = [3]string{
		"Ana",
		"Jose",
		"Maria",
	}

	fmt.Println(nomes[0:2])
	var nomesModificados [3]string = nomes
	nomesModificados[0] = "thamyris"
	fmt.Println("nomes modificados", nomesModificados)
	fmt.Println("nomes originais", nomes)
	fmt.Printf("%T\n", nomes) // \n é a quebra de linha
	fmt.Println("length", len(nomes))

}
