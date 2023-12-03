// section of ReferenceType.go -> Function section

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Function
	// create a function
	function := func() {
		fmt.Println("Hello World")
	}
	// call the function
	function()
	fmt.Println("============Function start ============")
	fmt.Printf("function type: %v, function value: %v\n", reflect.TypeOf(function).Kind(), function)
	fmt.Println("============Function end ============")

	// function with arguments
	functionWithArguments := func(name string) {
		fmt.Printf("Hello %v\n", name)
	}
	// call the function
	functionWithArguments("John")
	// function with return value
	functionWithReturnValue := func(name string) string {
		return fmt.Sprintf("Hello %v", name)
	}
	// call the function
	fmt.Println(functionWithReturnValue("John"))

	// function on struct
	// call the function
	person := Person{
		name: "John",
		age:  35,
	}
	person.sayHello()
}

// function on struct
// define a type
type Person struct {
	name string
	age  int
}

// define a function on type
func (person Person) sayHello() {
	fmt.Printf("Hello, my name is %v\n", person.name)
}
