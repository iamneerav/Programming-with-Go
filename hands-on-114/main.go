package main

import (
	"fmt"
)

func main() {

	x:= make ([]string,0,50)
	fmt.Printf("x is %T\n", x)

	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)

	fmt.Println("Adding US states to slice")

	y := append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	fmt.Println(y)
	
	// Getting values according to index position

	for i:=0; i < len(y); i++ {
		fmt.Println(y[i])
	}
	}