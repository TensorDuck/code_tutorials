package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	sq_x := x / 2

	for math.Abs((sq_x*sq_x)-x) > 0.0000001 {
		sq_x -= ((sq_x * sq_x) - x) / (2 * sq_x)
	}

	return sq_x

}

func main() {
	z := sqrt(23)
	fmt.Println(z)
}
