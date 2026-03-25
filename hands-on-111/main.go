package main

import (
	"fmt"
)

func main() {

	x := []int{42,43,44,45,46,47,48,49,50,51}

	for i:= 0; i < len(x); i++ {
		fmt.Println(x[i])
		
	}

	fmt.Println("----------------")

	fmt.Println("Creating new Slices")

	y := x[:5] // 
	fmt.Println(y)

	z := x[5:]
	fmt.Println(z)

	a := x[2:7] // Slicing includes source but doesn't exclude destination. This is why always calculate indices rather than position.
	fmt.Println(a)

	b := x[1:6]
	fmt.Println(b)
	
	}