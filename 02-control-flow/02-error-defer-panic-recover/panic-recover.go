// panic and recover
// feature to handle unexpected error
package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

// recover will be executed after panic
func f() {
	defer func() {
			if r := recover(); r != nil {
					fmt.Println("Recovered in f", r) // r is the error message from panic
			}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.") // it will be executed after recover
}
// panic will stop the program, but recover will be executed after panic
func g(i int) {
	if i > 3 {
			fmt.Println("Panicking!")
			panic(fmt.Sprintf("%v", i)) // this will stop the program, fmt.Sprintf("%v", i) will be the error message
	}
	defer fmt.Println("Defer in g", i) // this will be executed after panic, but before recover
	fmt.Println("Printing in g", i)
	g(i + 1)
}