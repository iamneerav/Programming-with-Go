package main

import (
	"fmt"
)

func main() {

	result := LearningGo("Go is interesting")

	fmt.Println(result)

}

func LearningGo(x string) string {
	return x
}
