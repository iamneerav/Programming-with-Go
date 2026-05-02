/*

Call a function.
That function returns another function.
Store the returned function in a variable.
Call that returned function.

*/

package main

import "fmt"

func main() {

	result := address("Toronto", "Ontario")

	fmt.Println(result())

}

func address(city string, state string) func() string {
	return func() string {
		return city + ", " + state + "," + " " + "Canada"
	}
}
