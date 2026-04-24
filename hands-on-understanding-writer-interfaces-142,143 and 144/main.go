/*
In this Functions exercise, we will learn about the following topics:

1. Writer Interface and writing to a file

2. Writer Interface and writing to a buffer

3. Writing to either a file or a buffer

*/

/*

	**Exploring Writer Interface and writing to a file**

	Refer to this link for information on Writer Interface: https://pkg.go.dev/io#Writer

	Lets try to use that 'Write' method

*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	Name string
}

func (p Person) Writertest(w io.Writer) error { // So, we refer to Writer interface which is used by "io" package
	_, err := w.Write([]byte(p.Name)) // We use the Write method of the Writer interface to write the Name field of the Person struct to the provided writer (w). The Name is converted to a byte slice using []byte() before being written.
	return err                        // As this return int and error

}

func main() {
	f, err := os.Create("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := Person{Name: "John Doe"}
	p.Writertest(f)             //  writes John Doe into file
	fmt.Println([]byte(p.Name)) // Give the byte slice of the Name field of the Person struct, which is printed to Console

	// Another way to write to a file using Writer Interface

	// s := []byte("Hello, World!")

	// _, err = f.Write(s) // We have to declare _, err as it returns int and error in function

	// if err != nil {
	// 	log.Fatal(err)
	// }

	/*
		**Exploring Writer Interface and writing to a buffer**

		Buffer: Data stored in a buffer. Buffer can be different types but lets take an example of Byte Buffer which is a type of buffer that holds data in memory as a sequence of bytes. It is commonly used for efficient manipulation and storage of binary data, such as when working with files, network connections, or other I/O operations.

		In this example, we will use the bytes.Buffer type from the bytes package, which provides a convenient way to work with byte buffers in Go.
	*/

	b := bytes.Buffer{} // package.type. Here package is "bytes" and type is "Buffer". We create a new instance of bytes.Buffer using a struct literal. This initializes an empty buffer that we can write to.

	b.Write([]byte{'h', 'e', 'l', 'l', 'o'}) // We use the Write method of the bytes.Buffer type to write the byte slice "Hello, World!" into the buffer. The string is converted to a byte slice using []byte() before being written.

	fmt.Println(b.String()) // We use the String() method of the bytes.Buffer type to get the contents of the buffer as a string, which is printed to Console

}

/*
Exploring Writing to either a file or a buffer

We also covered this with the first example of Writer Interface and writing to a file. We can write to a buffer by using the same Write method of the Writer interface. The only difference is that we need to provide the appropriate buffer when calling the Write method.

var b bytes.Buffer

p.Writertest(&b) // We pass the address of the bytes.Buffer instance to the Writertest method, which allows us to write to the buffer instead of a file.

fmt.Println(b.String()) // We use the String() method of the bytes.Buffer type to get the contents of the buffer as a string, which is printed to Console

*/
