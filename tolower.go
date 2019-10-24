package piscine

func ToLower(s string) string {
	str := []rune(s)

	for i, l := range s {
		if l >= 65 && l <= 90 && i >= 0 {
			str[i] = rune(l + 32)
		}
	}
	result := string(str)
	return result
}
