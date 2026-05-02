// Function Closure Example

package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(a float64) func() float64 {
	var c float64 // intializes c to 0
	return func() float64 {
		c++                   // increments c by 1 each time the function is called. First time c = 0 +1 = 1 second time c = 1 + 1 = 2, third time c = 2 + 1 = 3 and so on.
		return math.Pow(a, c) // returns a to the power of c. First time it returns 2 to the power of 1 = 2, second time it returns 2 to the power of 2 = 4, third time it returns 2 to the power of 3 = 8 and so on.
	}
}
