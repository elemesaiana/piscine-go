package main

import "fmt"

func IterativeFactorial(nb int) int {
	var res int
	res = 1
	for i := 0; i < nb+1; i++ {
		if i < 0 || i > 18 {
			res = 0
		} else if i == 0 || i == 1 {
			res = 1
		} else {

			res = res * (i)

		}

	}
	return res
}

func main() {
	arg := 5
	fmt.Println(IterativeFactorial(arg))
}
