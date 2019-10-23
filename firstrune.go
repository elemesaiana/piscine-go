package piscine

func FirstRune(s string) rune {
	v := ' '
	for i, v := range s {

		if v >= 0 && v <= 1000000000 && i == 0 {
			return v

		}
	}
	return v

}
