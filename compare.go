package piscine

func Compare(a, b string) int {

	c := 0
	k := 0
	index := -1

	for range a {
		c++
	}
	for range b {
		k++
	}
	for i := 0; i <= c-k; i++ {
		if a == b {
			index = 0
		} else if a < b {
			index = -1

		} else {
			index = 1
		}

	}

	return index
}
