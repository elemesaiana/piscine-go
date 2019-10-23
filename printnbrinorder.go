package piscine

import (
	"github.com/01-edu/z01"
)

func sort(mas []int) []int {
	len := 0
	for range mas {
		len = len + 1
	}
	for i := len; i > 0; i-- {
		for j := 1; j < i; j++ {
			if mas[j] < mas[j-1] {
				a := mas[j]
				mas[j] = mas[j-1]
				mas[j-1] = a
			}
		}
	}
	return mas
}

func PrintNbrInOrder(n int) {

	var slice []int
	if n == 0 {
		z01.PrintRune(rune(0 + 48))
	}

	for n > 0 {

		slice = append(slice, n%10)
		n /= 10
	}

	for _, digit := range sort(slice) {

		z01.PrintRune(rune(digit + 48))

	}

}
