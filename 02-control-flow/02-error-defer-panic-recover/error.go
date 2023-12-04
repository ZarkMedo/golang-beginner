// error panic recover
// error.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	// error
	fmt.Println("error start")
	err := errors.New("error message")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("error end")

	// use the method with error return
	fmt.Println("use the method with error return start")
	result, err := sum(1, -2) // this will return error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

// method with error return
func sum(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b should not be negative")
	}
	return a + b, nil
}
