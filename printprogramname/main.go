package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	for _, value := range arg {
		z01.PrintRune(value)
	}
	z01.PrintRune(10)
}
