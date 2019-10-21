package main

import "fmt"

func RecursiveFactorial(nb int) int {
	var res int
	res = 1
	if nb == 0 {
		return 1
	} else if nb > 0 && nb < 18 {
		for i := 1; i < nb+1; i++ {

			res = res * i

		}

	} else {
		res = 0
	}
	return res
}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
