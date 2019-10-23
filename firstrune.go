package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	v := ' '
	for i, v := range s {

		if v >= 0 && v <= 127 && i == 0 {
			return v

		}
	}
	return v

}
func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
