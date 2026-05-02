package main

import "fmt"

func main() {
	fmt.Println(Paradise("Hawaii"))
}

// Paradise is a function that takes a location as an argument and returns a string describing the idea of paradise.
func Paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
