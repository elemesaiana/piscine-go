package piscine

func AppendRange(min, max int) []int {
	var arr []int
	if min < max && min != max {
		for i := min; i < max; i++ {
			arr = append(arr, i)
		}
	}
	return arr
}
