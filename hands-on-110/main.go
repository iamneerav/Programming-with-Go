package main

import (
	"fmt"
)

func main() {

	x := []int{42,43,44,45,46,47,48,49,50,51}

	for i:= 0; i < len(x); i++ {
		fmt.Println(x[i])
	} 
	
	}