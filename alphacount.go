package piscine

import (
	"github.com/01-edu/z01"
)

func AlphaCount(str string) int {
	counter := 0

	for index, letter := range str {

		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			z01.PrintRune(rune(index))
			counter++
		}

	}

	return counter
}
