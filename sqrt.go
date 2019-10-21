package piscine

func Sqrt(nb int) int {

	var res int
	res = 0
	for i := 0; i < nb; i++ {
		res = i * i
		if res == nb {
			return i
		}
	}
	if nb == 1 {
		return 1
	}
	return 0
}
