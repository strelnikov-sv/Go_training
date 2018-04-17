package main

import (
	"fmt"
	"math"
)

const number = 2
const diff = 0.0000001

func sqrt(x float64) (float64, int) {

	z := 1.0
	zcon := 0.0
	k := 1

	for i := 0; i < k; i++ {
		zcon = z - (z*z-x)/(2*z)
		z = zcon
		if zcon-math.Sqrt(number) > diff {
			k++
		}
	}

	return zcon, k
}

func main() {
	bestsqrt := math.Sqrt(number)
	newsqrt, times := sqrt(number)
	delta := math.Abs(newsqrt - bestsqrt)
	fmt.Println("Функция math.Sqrt", bestsqrt)
	fmt.Println("Моя функция", newsqrt)
	fmt.Println("Разница", delta)
	fmt.Println("Кол-во итераций", times)

}
