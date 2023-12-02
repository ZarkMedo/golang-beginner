// how to declare variables

package main

import (
	"fmt"
)

func main() {
	// var to declare variables
	var name = "John"
	fmt.Print(name)
	// var to declare multiple variables
	var age, height int = 35, 180
	fmt.Println(age, height)
	// declare variables use :=, you can omit the var keyword
	weight := 80
	fmt.Println(weight)
	// declaration without initialization
	var interval int
	interval = 10
	fmt.Println(interval)
	// declare multiple variables without initialization
	var book, author string
	book = "The Hobbit"
	author = "J.R.R Tolkien"
	fmt.Println(book, author)

}
