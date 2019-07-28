package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      1200.45,
		"Gabriela Silva": 1245.56,
		"Pedro Junior":   1224.43,
	}

	funcsESalarios["Rafael Filho"] = 1234.54
	delete(funcsESalarios, "Inexistente") //não acontece nada

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
