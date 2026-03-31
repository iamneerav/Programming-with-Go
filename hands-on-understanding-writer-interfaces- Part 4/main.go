/*
In this Functions exercise, we will learn about the following topics:

1. Writer Interface and writing to a file

2. Writer Interface and writing to a buffer

3. Writing to either a file or a buffer

*/

/*
Exploring Writer Interface and writing to a file

Refer to this link for information on Writer Interface: https://pkg.go.dev/io#Writer

Lets try to use that 'Write' method

*/

// 1st example of Writer Interface and writing to a file

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	Name string
}

func (p Person) Writertest(w io.Writer) error { // So, we refer to Writer interface which is used by "io" package
	x, err := w.Write([]byte(p.Name)) // We use the Write method of the Writer interface to write the Name field of the Person struct to the provided writer (w). The Name is converted to a byte slice using []byte() before being written.
	fmt.Println("Bytes written: ", x) // We print the number of bytes written to the writer using fmt.Printf. The variable x contains the number of bytes that were successfully written.
	return err                        // As this return int and error

}

func main() {
	f, err := os.Create("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := Person{Name: "John Doe"}
	fmt.Println(p.Writertest(f)) //  writes John Doe into file
	fmt.Println([]byte(p.Name))  // Give the byte slice of the Name field of the Person struct, which is [74 111 104 110 32 68 111 101] (ASCII values of "John Doe")

}

/*
Exploring Writer Interface and writing to a buffer
*/

/*
Exploring Writing to either a file or a buffer
*/
