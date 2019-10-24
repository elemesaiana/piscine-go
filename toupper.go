package piscine

func ToUpper(s string) string {

	a := []rune(s)

	for i, str := range s {
		if str >= 97 && str <= 122 && i >= 0 {
			a[i] = rune(str - 32)

		}

	}
	stringg := string(a)
	return stringg
}
