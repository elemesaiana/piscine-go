package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	c := 0

	for range b {
		c++
	}
	if c != 0 {
		for index, str := range a {
			if c > 0 && str == b[0] {
				return index

			}
		}
	}

	return -1
}
