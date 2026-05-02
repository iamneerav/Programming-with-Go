package main

import (
	"fmt"
)

func main() {

	fmt.Println(location("Maldives"))

}

func location(x string) string {
	return fmt.Sprintf("I am in %s", x)
}
