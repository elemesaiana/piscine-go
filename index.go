package main

import "fmt"

func Index(s string, toFind string) int {

	c := 0
	k := 0
	index := -1

	for range s {
		c++
	}
	for range toFind {
		k++
	}
	for i := 0; i <= c-k; i++ {
		if s[i:i+k] == toFind {
			index = i
			break
		}
	}

	return index
}
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
