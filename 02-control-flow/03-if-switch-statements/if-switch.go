// if and switch

package main

import (
	"fmt"
)

func main() {
	// if
	fmt.Println("if start")
	if 1 > 0 {
		fmt.Println("1 > 0")
	} else if 1 < 0 {
		fmt.Println("1 < 0")
	} else {
		fmt.Println("1 == 0")
	}
	fmt.Println("if end")

	// switch
	fmt.Println("switch start") // switch will break after case, if you want to continue, use fallthrough
	switch 1 {
	case 0:
		fmt.Println("0")
	case 1:
	case 2:
		fmt.Println("1")
	case 3:
		fmt.Println("3")
		fallthrough // this will continue to next case
	default:
		fmt.Println("default")
	}
	fmt.Println("switch end")

}
