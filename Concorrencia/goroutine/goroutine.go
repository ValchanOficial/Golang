package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//--------------função serial

	// executa os dois fale()
	// fale("Maria", "Ola Mundo!", 3)
	// fale("João", "Hello World", 1)

	//--------------função executada de forma independente

	// o main() executa primeiro que o fale()
	// go fale("Maria", "Ola Mundo!", 3)
	// go fale("João", "Hello World", 3)

	// o main executa o fale("João") e o fale("Maria") até o João terminar
	// go fale("Maria", "Ola Mundo!", 3)
	// fale("João", "Hello World", 3)

	//fmt.Println("Fim!")
}
