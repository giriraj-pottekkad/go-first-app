package loops

import (
	"fmt"
)

var progression = make([]int, 100)

//Looping exported
func Looping() []int {
	var k int = 0
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			progression[k] = i * j
			k = k + 1
		}
	}
	fmt.Println("Output progression populated...")
	return progression
}
