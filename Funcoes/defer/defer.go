package main

import "fmt"

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!") //última ação executada dentro do método
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return numero
	}
	fmt.Println("Valor baixo...")
	return numero
}
