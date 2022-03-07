package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2

	for z*z-x < -0.001 || z*z-x > 0.001 {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
