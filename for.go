package main

import "fmt"

func foreach() {
	soma := 0

	for i := 0; i <= 10; i++ {
		soma += i
	}
	fmt.Println(soma)
}
