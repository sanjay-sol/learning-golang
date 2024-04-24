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
	fmt.Println(name, name3, name2)

	// shorthand
	age := 20 //! cant do outside the function
	fmt.Println(age)

	fmt.Printf("name %s \n", name)
	fmt.Printf("name %T", name)

	// multiple variables

	var (
		name4 = "John Doe"
		age2  = 20
	)

	fmt.Println(name4, age2)

	// boolean

	var isCool = true

	fmt.Println(isCool)

	// format specifiers

	fmt.Printf("age %d \n", age2)
	fmt.Printf("age %0.3f \n", 20.5)
	fmt.Printf("age %t \n", true)
	fmt.Printf("age %T \n", age2)

	var str = fmt.Sprintf("name %s", name)
	fmt.Println(str)

	// type conversion

	var age3 = 20
	var age4 = 20.5

	fmt.Println(float64(age3) + age4)

}
