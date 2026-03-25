package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for x, y := range m {
		fmt.Printf("key: %v \t value: %v\n", x, y)
	}
	}
	