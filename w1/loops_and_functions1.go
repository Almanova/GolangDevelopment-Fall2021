package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	prev, z := 2.0, 1.0
	for math.Abs(prev - z) >= 1e-9 {
		prev, z = z, z - (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Printf("My function: %g\nmath.Sqrt(): %g\n", 
		Sqrt(15), math.Sqrt(15))
}
