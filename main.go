package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	// variables

	var name string = "John Doe"
	const name2 = "Jane Doe2"
	var name3 string

	name3 = "Jane Doe3"
	// name2 = "Jane Doe4" // cannot assign to name2
	fmt.Println(name,name3,name2);

	// shorthand
	age := 20 //! cant do outside the function
	fmt.Println(age)

}
