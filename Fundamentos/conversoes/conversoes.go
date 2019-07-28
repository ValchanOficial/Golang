package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) //1.2

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal) //6

	//cuidado...
	fmt.Println("Teste " + string(97)) //a

	//int pra string
	fmt.Println("Teste " + strconv.Itoa(123)) //123

	//string pra int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") //ou 1
	if b {
		fmt.Println("Verdadeiro")
	}
}
