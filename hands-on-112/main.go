package main

import (
	"fmt"
)

func main() {

	x := []int{42,43,44,45,46,47,48,49,50,51}
	y := append(x,52)
	fmt.Println(y)

	fmt.Println("----------")

	z := append(y,53,54,55)
	fmt.Println(z)

	fmt.Println("----------")

	a := []int{56,57,58,59,60}
	z = append(z,a...)
	
	fmt.Println(z)
	}