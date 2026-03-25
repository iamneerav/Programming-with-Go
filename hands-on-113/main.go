package main

import (
	"fmt"
)	

func main() {

	x := []int{42,43,44,45,46,47,48,49,50,51}
	
	y := x[:3]
	
	z := x[6:]

	x = append(y,z...)

	// We can do like this as well

	// x = append(x[:3], x[6:]...)

	fmt.Println(x)
	}