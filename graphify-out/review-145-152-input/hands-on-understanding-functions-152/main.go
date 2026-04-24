/*
In this Functions exercise, we will learn about Wrapper Function
*/

/*
Wrapper function wraps around another function to enhance its functionality without modifying the original function.

It allows us to add additional behavior or logic before or after the execution of the original function.

This is great for logging, error handling, or any other cross-cutting concerns that we want to apply to multiple functions without changing their implementation.
*/

/*
Reflecting to the older example of wrapper interface we did in last exercise before understanding in depth.
This is an example of wrapper func for logging

A wrapper function is a function that adds extra behavior around existing functionality without changing the original code.

In this example, the wrapper function is writeLog(s fmt.Stringer), which adds logging behavior.

The types book and count both implement the String() string method, so they satisfy the fmt.Stringer interface.

Because both types satisfy fmt.Stringer, the same writeLog function can accept both and log their string representations.

This approach makes code reusable, avoids repetition, and shows how interfaces help one function work with multiple types.

*/

package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string { // Any type of book is of type stringer interface because it has the method String() which returns a string representation of the type book
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string { // Similarly, any type of count is of type stringer interface because it has the method String() which returns a string representation of the type count
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func writeLog(s fmt.Stringer) { //  This is a wrapper function for logging. The parameter s is of type fmt.Stringer which means it can accept any type that implements the String() method, which is the stringer interface. So, we can pass both book and count types to this function because they both implement the String() method.
	log.Println(s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var c count = 42

	writeLog(b) // We pass the variable b of type book to the writeLog function, which accepts a parameter of type fmt.Stringer. Since book implements the String() method, it satisfies the fmt.Stringer interface, and we can pass it to the writeLog function. The writeLog function will then call the String() method of the book type to get its string representation and log it.

	writeLog(c) // Similarly, we pass the variable c of type count to the writeLog function. Since count also implements the String() method, it satisfies the fmt.Stringer interface, and we can pass it to the writeLog function. The writeLog function will then call the String() method of the count type to get its string representation and log it.
}

/*

Coming to this lecture, we will learn about wrapper function for error handling.

1. This example teaches a wrapper function for error handling.

2. The wrapper function is `readFile(fileName string) ([]byte, error)`. `readFile()` calls `os.ReadFile(fileName)` to read the file contents.

3. If an error happens, `readFile()` does not return the raw error directly, instead, it wraps the error using `fmt.Errorf(...)` and adds a clearer message. This makes debugging easier because the error shows where the problem happened.

4. In `main()`, the code calls `readFile("poem.txt")`. `main()` checks whether `err != nil`. If there is an error, `log.Fatalf(...)` prints the error and stops the program. If there is no error, the code prints the file contents as bytes and also as a string.

5. This example shows how wrapper functions can make code cleaner and improve error handling.

6. Main takeaway: a wrapper function adds extra behavior around another function without changing the original function itself.

*/

/*

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("error in main in readFile: %s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error in readfile: %s", err)
	}
	return xb, nil
}

*/
