package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(foo(a...))

	fmt.Println("--------------------")

	fmt.Println(bar(a))

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func bar(a []int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
