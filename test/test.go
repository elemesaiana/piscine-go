package main

import "fmt"

func main() {

	aString := "Hello"

	aStringChangeble := []byte(aString)

	aStringChangeble[0] = 'A'

	aStringFinalized := string(aStringChangeble)

	fmt.Println(aString)

	fmt.Println(aStringChangeble)

	fmt.Println(aStringFinalized)

}
