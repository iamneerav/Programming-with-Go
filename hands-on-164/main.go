package main

import "fmt"

func main() {
	result := func(a int, b int) int { // Anonymous function
		return a + b
	}(2, 3)

	fmt.Println(result)

	fmt.Println("--------------------")
	fmt.Println("Another example of anonymous function")

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}
