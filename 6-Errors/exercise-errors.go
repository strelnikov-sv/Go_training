package main

import (
	"fmt"
	"math"
)

const diff = 0.0001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	zcon := 0.0
	k := 1
	i := 0

	fmt.Println(z)
	for i := 0; i < k; i++ {
		zcon = z - (z*z-x)/(2*z)
		z = zcon
		if math.Sqrt(z-zcon) > diff {
			k++
		}
		fmt.Println(z)
	}

	fmt.Println(i, "iterations")

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
