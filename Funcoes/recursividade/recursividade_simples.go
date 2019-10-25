package main

import "fmt"

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}
