package main

import (
	"fmt"
	m "math"
)

func main() {
	//declarando variáveis
	const PI float64 = 3.1415
	var raio = 3.2

	//area = PI * m.Pow(raio, 2) -> se já estivesse declarada
	area := PI * m.Pow(raio, 2) // -> cria e atribui valor para variável
	fmt.Println("A área da circunferência é:", area)

	//outra forma de definir const e var
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	//declarando mais de uma var
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
