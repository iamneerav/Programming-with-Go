/*
In this Functions exercise, we will learn about the following topics:

1. Exploring the stringer interface

2. Expanding on the stringer interface - wrapper func for logging

3. Writer Interface and writing to a file

4. Writer Interface and writing to a buffer

5. Writing to either a file or a buffer

*/

/*

Exploring the stringer interface

Stringer Interface: The stringer interface is a built-in interface in Go that defines a single method, String(), which returns a **STRING representation of the TYPE** that implements it. It is commonly used to provide a human-readable representation of a type when it is printed or logged.

type Stringer interface {
    String() string 		// Method that returns a `string` representation of the type
}

*/

// Example of Stringer Interface

package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintln("Book Title is", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprintln("The count is", strconv.Itoa(int(c)))
}

func main() {

	b1 := book{title: "The Great Gatsby"}
	c1 := count(5)

	log.Println(b1.String()) // We call the function String() + type  struct "book" which has has instance "b1" and it will return the string representation of the type "book" which is "Book Title is The Great Gatsby"
	log.Println(c1.String()) // We call the function String() + type  struct "count" which has has instance "c1" and it will return the string representation of the type "count" which is "The count is 5"
}

/*
2. Expanding on the stringer interface - wrapper func for logging

Implements a wrapper func that takes fmt.Stringer as an argument and logs the string representation of the type that implements the Stringer interface.
*/

// Example of wrapper func for logging

/*

If we see the previous example, We can do like this:

func main() {

	b1 := book{title: "The Great Gatsby"}
	c1 := count(5)

	Case 1: If we the below

	fmt.Println(b1.String())
	fmt.Println(c1.String())

	We get output like this:
		Book Title is The Great Gatsby

		The count is 5

	Case 2: If we add log

	log.Println(b1.String())
	log.Println(c1.String())

	We get output like this:

		2026/03/23 22:11:04 Book Title is The Great Gatsby

		2026/03/23 22:11:04 The count is 5


// Now what if we want to print some message before "Book Title is The Great Gatsby"//

We can expand the functionaily of a function by creating a wrapper function that takes 'fmt.Stringer' as an argument and logs the string representation of the type that implements the Stringer interface.

func loginfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}

// NOTES:
1. Why we use fmt.Stringer as an argument:
	We took fmt.Stringer as an argument which is an interface that has a method String() which returns a string representation of the type that implements it. Also, both book and count uses the String() method to return their string representation, so we can use fmt.Stringer as an argument to log the string representation of both book and count types.

2. Why we call s.String() inside the loginfo function:
	We call s.String() inside the loginfo function, it will call the String() method of the type that implements the Stringer interface (which is used by book and count) and return the string representation of that type, which will be logged by log.Println.

	So, we are using this wrapper function if we need more information in the log message, we can easily add it in the loginfo function without changing the String() method of the types that implement the Stringer interface.

func main() {

	b1 := book{title: "The Great Gatsby"}
	c1 := count(5)

	loginfo(b1)
	loginfo(c1)
}

Output:

2026/03/23 22:11:04 LOG FROM 138 Book Title is The Great Gatsby

2026/03/23 22:11:04 LOG FROM 138 The count is 5

*/

/*

Writer Interface and writing to a file

*/
