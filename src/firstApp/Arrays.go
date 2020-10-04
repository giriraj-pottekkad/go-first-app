package main

import (
	"fmt"
)

func printArrays() {
	numbers := []int{1, 2, 3}
	words := make([]string, 10, 100)
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", numbers[:])
	words[0] = "Giriraj"
	words[1] = "Harikrishnan"
	words[2] = "Padmavathy"
	words[3] = "Geetha"
	fmt.Println(len(words), cap(words))
}
