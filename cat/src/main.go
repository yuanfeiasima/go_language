package main

import (
	"./goroutine"
	"fmt"
)
func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go goroutine.Sum(a[:len(a)/2], c)
	go goroutine.Sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Printf(x, y, x+y)
}
