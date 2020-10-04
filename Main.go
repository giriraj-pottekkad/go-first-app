package main

import (
	arrays "firstapp/arrays"
	maps "firstapp/maps"
	"fmt"
)

func main() {
	var i bool
	a := "I am not a good programmer"
	v := 6 + 2i
	fmt.Println("Hello Go....!")
	fmt.Printf("%v \n", i)
	fmt.Println(a)
	fmt.Println(imag(v))
	arrays.PrintArrays()
	maps.PopulateMaps()
	fmt.Println(maps.Weeks[1])
}
