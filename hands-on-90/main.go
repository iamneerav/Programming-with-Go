package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	m1 := m["James"]
	fmt.Println(m1)

	// Comma Ok Idiom/statement

	if x, ok := m["James"]; ok {
		fmt.Println("This map contains James and his age is is", x)
	}

	if x, ok := m["Q"]; !ok {
		fmt.Println("This map does not contain Q as its value is", x)
	}
}