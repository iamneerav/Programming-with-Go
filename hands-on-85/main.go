package main

import (
	"fmt"
)

func main() {

	var i int = 0
	for i < 100 {
		// if i % 2 ==0 { // Prints only even numbers
		// fmt.Println(i) 
		// }
		// i++

		if i % 2 !=0 { // Breaks if the number is odd
		fmt.Println(i)
		break 
		}
		i++
	}
	}