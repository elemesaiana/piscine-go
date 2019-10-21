package piscine

func Sqrt(nb int) int {

	var res int
	res = 0
	for i := 1; i < nb; i++ {
		res = i * i
		if res == nb {
			return i
		}
	}
	return 0
}
