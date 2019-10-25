package main

import "fmt"

func main() {
	x := 20
	fmt.Println("[main] valor de x:", x)

	imprimeX := closure()
	imprimeX()
}

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println("[closure] valor de x:", x)
	}
	return funcao
}
