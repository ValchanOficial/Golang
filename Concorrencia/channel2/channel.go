package main

import (
	"fmt"
	"time"
)

//Tipo Channel (canal) - é a forma de comunicação entre goroutines

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) //assincronismo

	a, b := <-c, <-c //recebendo os dados do canal //sincronismo, só retorna os dados quando a e b estejam prontos
	fmt.Println(a, b)
	fmt.Println(<-c)
	//fmt.Println(<-c) //fatal error: all goroutines are asleep - deadlock!
}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //envando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}
