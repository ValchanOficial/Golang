package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1300.32,
			"Guga Pereira":   1200.43,
		},
		"J": {
			"José João": 1234.43,
		},
		"P": {
			"Paula Miranda": 2132.21,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
