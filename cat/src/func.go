package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main(){
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_xz := max(x, z)
	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))

	x_add_y, x_multiply_y := SumAndMultiply(x, y)
	fmt.Printf("x_add_y:%d, x_multiply_y:%d", x_add_y, x_multiply_y)
}

func SumAndMultiply(A, B int)(int, int){
	return A+B, A*B
}