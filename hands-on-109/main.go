package main

import (
	"fmt"
)

func main() {

	x := []int{1,2,3,4,5}
	fmt.Println(x)

	for _, a := range x {
		fmt.Printf("a = %v - type = %T\n",a,a)
		x[a] = a
	}
	
	}