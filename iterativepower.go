package piscine

func IterativePower(nb int, power int) int {

	var res int
	res = 1
	if power > 0 {
		for i := 1; i <= power; i++ {
			res = res * nb
		}
	} else {
		res = 1
	}

	return res
}
