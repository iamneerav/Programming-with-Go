package main

import (
	"fmt"
	"math/rand"
)

func main() {
	 x := rand.Intn(10)
	 y := rand.Intn(10)	

	 fmt.Printf("x is %v \t y is %v\n", x, y)

	 // If statement

	 if x <4 && y <4 {
		fmt.Println("x and y are less than 4")
	 } else if x >6 && y >6 {
		fmt.Println("x and y are greater than 6")
	 } else if x >= 4 && x <=6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	 } else if y !=5 {
		fmt.Println("y is not equal to 5")
	 } else {
		fmt.Println("None of previous cases were met")
	 }
}