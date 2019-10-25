package main

import "fmt"

func main() {
	n := 1
	inc1(n)        //por valor
	fmt.Println(n) // 1

	// & usado para obter o endereço da variável
	inc2(&n)       //por referência
	fmt.Println(n) // 2
}

func inc1(n int) {
	n++
}

// um ponteiro é representado por um *
func inc2(n *int) {
	// * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}
