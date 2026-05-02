// Callback example

package main

import (
	"fmt"
)

func main() {

	result := location("Merivale Road", "Ottawa", country)

	fmt.Println(result)

}

func location(street string, city string, c func(a string) string) string {
	return fmt.Sprintf("I am in %s, %s, %s", street, city, c("Canada"))
}

func country(a string) string {
	return a
}
