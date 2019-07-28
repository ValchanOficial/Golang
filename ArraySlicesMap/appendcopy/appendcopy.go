package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	//append e copy não funcionam com array
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) //apenas copia 2 elementos para o slice2
	fmt.Println(slice2)
}
