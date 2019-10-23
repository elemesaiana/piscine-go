package piscine

func FirstRune(s string) rune {
	v := ' '
	for i, v := range s {

		if v >= 0 && v <= 127 && i == 0 {
			return v

		}
	}
	return v

}
