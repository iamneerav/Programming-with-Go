package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var i int
	for i = 0; i <42; i++ {	
		x := rand.Intn(5)

		fmt.Printf("x is %v\n", x)

		// Switch statement
		
		switch x {
		case 0:
			fmt.Println("x is 0")	
		case 1:
			fmt.Println("x is 1")
		case 2:
			fmt.Println("x is 2")
		case 3:
			fmt.Println("x is 3")
		case 4:
			fmt.Println("x is 4")
		}}
	}