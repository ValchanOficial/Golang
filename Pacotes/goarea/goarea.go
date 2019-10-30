// //package area dentro de C:\Users\user\go\src\github.com\cod3rcursos\area

package goarea

import "math"

//PI : é uma proporção númerica deifinida pela relação entre
// o perímetro de uma circuferência e seu diâmetro
const PI = 3.1416

//Circ : é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

//Rect : é responsável por calcular a área de de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visível!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
