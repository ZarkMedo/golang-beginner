// for loop and range

package main

import "fmt"

func main()  {
	// for loop
	fmt.Println("for loop 0 to 4 start")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("for loop 0 to 4 end")

	// range
	fmt.Println("range start")
	nums := []int{1, 2, 3, 4, 5}
	for index, num := range nums {
		fmt.Println(index, num)
	}
	fmt.Println("range end")

	// continue and break
	fmt.Println("continue and break start")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("continue and break end")
}