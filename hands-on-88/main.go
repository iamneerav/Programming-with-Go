package main

import (
	"fmt"
)

func main() {

	xi := []int{42, 43, 44, 45, 46, 47}

	for i, x := range xi {
		fmt.Printf("index: %v \t value: %v\n", i, x)
	}
}