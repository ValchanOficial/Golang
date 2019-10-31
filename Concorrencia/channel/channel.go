package main

import "fmt"

func main() {
	//criando canal
	ch := make(chan int, 1)

	ch <- 1 //envia dados para o canal (escrita)
	<-ch    //recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
