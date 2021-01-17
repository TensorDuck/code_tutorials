package main

import "fmt"

func fib(x int, y int) int{
	return x + y
}

func main() {
	x := 1
	y := 1
	z := 0
	for z < 100{
		z = fib(x, y)
		x = y
		y = z
		fmt.Println(z)
	}
}