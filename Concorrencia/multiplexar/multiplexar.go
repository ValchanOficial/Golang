package main

import (
	"fmt"

	html "github.com/ValchanOficial/Golang/HTML"
)

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br/portal", "https://www.google.com"),
		html.Titulo("https://www.linkedin.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar - misturar(mensagens) em um canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}
