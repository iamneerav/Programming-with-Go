package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}

	x:=[][]string{jb,jm}

	fmt.Println(x)
	
	for i:= range x {
		fmt.Println(x[i])
	}
}