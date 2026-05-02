package main

import (
	"fmt"
)

func main() {

	// defer statements are executed in LIFO order, meaning the last defer statement will be executed first when the function returns.
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	defer fmt.Println(5)

	fmt.Println("Hello") // This runs first as all of above are deferred and will run in LIFO order

}
