package main

import (
	"fmt"
)

func Sqrt(square complex128) complex128 {
	var z complex128 = 100

	for i := 0; i < 200; i++ {
		if z*z == square {
			return z
		}

		z -= (z*z - square) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(0.5))
}
