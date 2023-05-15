package main

import (
	"fmt"
	"math"
)

func soma(x, y int) int {
	return x + y
}

func calcular(a int) (int, int) {
	var quadrado = a * a
	var cubo = a * a * a

	return quadrado, cubo
}

// pode ser das duas maneiras, só em funcoes curtas
func calcular2(a int) (quadrado int, cubo int) {
	quadrado = a * a
	cubo = a * a * a

	return quadrado, cubo
}

func main() {
	const (
		nome   string = "Thamyris"
		idade  int    = 34
		aberto        = true
	)

	var Numero = math.Max(3, 5) // nao pega com const

	fmt.Println(calcular(2))
	fmt.Println(calcular2(2))
	fmt.Println("maior numero", Numero)
	fmt.Println("meu nome é", nome, "tenho", idade, "anos")
	fmt.Println("valor da soma é", soma(2, 3))

	// fmt.Printf("tipo: %T Valor: %v", aberto, nome)

	ifelse()
	foreach()
	while()
	array()
	slice()
}
