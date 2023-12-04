// error panic recover
// error.go
package main

import (
	"fmt"
	"os"
)

func main() {

	// defer
	fmt.Println("defer start")
	deferReturn() // output: 0 because defer will be executed after return
	fmt.Println("defer end")

	// defer for loop
	fmt.Println("defer for loop start")
	deferForLoop() // output: 4 3 2 1 0 because defer will be executed after for loop, first in last out
	fmt.Println("defer for loop end")
	// defer surrounding function
	fmt.Println("defer surrounding function start")
	deferSurroundingFunction(1) // output: 1 then 2 because defer will be executed after surrounding function
	fmt.Println("defer surrounding function end")

	// defer use case
	fmt.Println("defer use case start")
	file, err := openFile("defer.go")	// it will close the file after return
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())
	fmt.Println("defer use case end")
}


// method with defer
func deferReturn() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// method with defer for loop
func deferForLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// method with defer for surrounding function
func deferSurroundingFunction(i int) {
	defer func ()  {
		i++
		fmt.Println(i)
	}() // this will be executed after surrounding function
	fmt.Println(i)
}

// defer usecase, open file
func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
			return nil, err
	}

	defer file.Close()

	// Do something with the open file

	return file, nil
}