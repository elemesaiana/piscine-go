package piscine

import (
	"sort"

	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {

	var a, b, c int
	var slice []int

	if n >= 0 || n < 2147483647 {
		if n == 0 {
			z01.PrintRune(rune(n + 48))
		} else {
			a = n % 10
			b = n / 10 % 10
			c = n / 100 % 10
			slice = append(slice, c, b, a)

			sort.Ints(slice)

			for _, digit := range slice {

				z01.PrintRune(rune(digit + 48))

			}

		}

	}

}
