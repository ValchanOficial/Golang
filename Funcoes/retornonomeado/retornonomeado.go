package main

import "fmt"

func main() {
	x, y := 1, 2
	fmt.Println("Antes:", x, ",", y)
	x, y = trocar(x, y)
	fmt.Println("Depois:", x, ",", y)
}

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return
}
