package main

import (
	"fmt"
)

func main() {

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	for i, a := range x {
		fmt.Printf("a = %v - type = %T\n", a, a)
		x[i] = a
	}

}
