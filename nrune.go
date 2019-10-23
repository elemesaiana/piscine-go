package piscine

func NRune(s string, n int) rune {
	count := 0
	for i := range s {
		if i >= 0 {
			count++
		}

	}
	if n-1 >= 0 && count >= n {
		val := []rune(s)
		v := n - 1
		return val[v]
	}
	return 0
}
