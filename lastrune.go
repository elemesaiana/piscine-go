package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	v := ' '
	count := 0
	for i := range s {
		if i >= 0 {
			count++
		}

	}
	for i, v := range s {

		if v >= 0 && v <= 1000000000 && i == count-1 {
			return v

		}
	}
	return v

}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
