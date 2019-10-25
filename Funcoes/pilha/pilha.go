package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	f1()
}

func f3() {
	fmt.Println("---------------PrintStack---------------")
	debug.PrintStack()
	fmt.Println("---------------PrintStack---------------")
}

func f2() {
	f3()
}

func f1() {
	f2()
}
