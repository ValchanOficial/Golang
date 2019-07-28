package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[21321322] = "Pedro"
	aprovados[21421424] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[21421424])
	delete(aprovados, 21321322)
	fmt.Println(aprovados)
}
