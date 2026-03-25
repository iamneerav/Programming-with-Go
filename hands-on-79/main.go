package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is the init function")
}

func main() {
	 x := rand.Intn(250)
	//  x := 101
	 fmt.Println(x)

	//  if x >= 0 && x <= 100 {
	// 	fmt.Println ("x is between 0 and 100")
	//  } else if x >= 101 && x <= 200 {
	// 	fmt.Println ("x is between 101 and 200")
	//  } else if x >= 201 && x <= 250 {
	// 	fmt.Println ("x is between 201 and 250")
	//  }

	//  y := rand.Intn(3)
	//  fmt.Println(y)

	// Switch Statement
	switch {
		case x >= 0 && x <= 100 :
			fmt.Println("x is between 0 and 100")
		case x >= 101 && x <= 200 :
			fmt.Println("x is between 101 and 200")
		case x >= 201 && x < 250   :
			fmt.Println("x is between 201 and 250")	
		default:
			fmt.Println("x is greater than 250")	
		}


}