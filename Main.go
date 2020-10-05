package main

import (
	arrays "firstapp/arrays"
	loops "firstapp/loops"
	maps "firstapp/maps"
	structs "firstapp/structs"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//Variables
	var i bool
	a := "I am not a good programmer"
	v := 6 + 2i
	fmt.Println("Hello Go....!")
	fmt.Printf("%v \n", i)
	fmt.Println(a)
	fmt.Println(imag(v))

	//Arrays
	arrays.PrintArrays()

	//Maps
	maps.PopulateMaps()
	fmt.Println(maps.Weeks[1])

	//Struct
	car := structs.DisplayVehicles()
	fmt.Println(car)

	//Looping
	output := loops.Looping()
	for i, v := range output {
		if v != 0 {
			fmt.Println(i, v)
		}

	}

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		panic("Error connecting to the url")
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Robots.tx cannot be read")
	}

	fmt.Printf("%s", robots)
}
