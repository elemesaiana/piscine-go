package piscine

func FirstRune(s string) rune {
	v := ' '
	for i, v := range s {

		if (v >= 65 && v <= 90 && i == 0) || (v >= 97 && v <= 122 && i == 0) {
			return v

		}
	}
	return v

}
