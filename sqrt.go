package piscine

import (
	"math"
)

func Sqrt(nb int) int {

	var a int
	a = 0
	for i := 1; i <= nb; i++ {
		a = i * i

		if a == nb {

			a = int((math.Sqrt(float64(nb))))
			return a
		}
	}
	return 0
}
