package piscine

func AlphaCount(str string) int {
	counter := 0

	for index, letter := range str {

		if (index >= 0 && letter >= 65 && letter <= 90) || (index >= 0 && letter >= 97 && letter <= 122) {

			counter++
		}

	}

	return counter
}
