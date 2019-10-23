package main

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	val := []rune(s)
	v := n - 1
	return val[v]

}
func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
