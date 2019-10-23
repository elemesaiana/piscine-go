package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {

	var slice []int

	if n >= 0 || n < 2147483647 {
		if n == 0 {
			z01.PrintRune(rune(n + 48))
		} else {
			for n > 0 {
				slice = append(slice, n%10)
				n /= 10
			}

		}

		for _, digit := range slice {

			z01.PrintRune(rune(digit + 48))

		}

	}
}
