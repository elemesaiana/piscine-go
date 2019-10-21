package piscine

func Fibonacci(index int) int {

	var f1, f2, fib int
	f1 = 1
	f2 = 1
	fib = 1
	if index > 0 {
		for i := 3; i <= index; i++ {
			fib = f1 + f2
			f1 = f2
			f2 = fib

		}
		return fib
	} else if index == 0 {
		return 0
	} else {
		return -1
	}

}
