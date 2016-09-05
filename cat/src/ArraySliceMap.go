package main

import "fmt"

func main(){

	//array
	var arr_int [10]int
	arr_int[0] = 10
	arr_int[1] = 20

	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{1, 2, 3}

	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", c)
}