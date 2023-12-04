// Collection/Aggregation or Non-Reference Types
// 1. Array
// 2. Struct

package main

import (
	"fmt"
	"time"
)

func main() {
	// Array
	// elements declared with a specific size
	newArray := []string{"a", "b", "c", "d", "e"} // <==> newArray := [5]string{"a", "b", "c", "d", "e"}
	// point the size, elements are not declared
	newArray1 := [10]string{}
	// add all of the newArray's elements to newArray1
	for i := 0; i < len(newArray); i++ {
		newArray1[i] = newArray[i]
	}
	// we will detail about difference between array and slice in the reference type section

	fmt.Println("============Array start ============")
	fmt.Println(newArray)
	fmt.Println(newArray1)
	fmt.Println("============Array end ============")

	// Array for Arguments
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println("============Array for Arguments start ============")
	// check the pointer of the numbers and value
	fmt.Printf("numbers pointer: %p, numbers value: %v\n", &numbers, numbers)
	// call the transfer function
	// newNumbers := transfer(numbers)
	newNumbers := transfer(numbers)
	// check the pointer of the numbers and value
	fmt.Printf("numbers pointer: %p, numbers value: %v\n", &numbers, numbers)
	fmt.Printf("newNumbers pointer: %p, newNumbers value: %v\n", &newNumbers, newNumbers)
	fmt.Println("============Array for Arguments end ============")

	// Struct
	// define a struct
	type Person struct {
		name string
		age  int
		date time.Time
	}
	// create a struct
	person := Person{
		name: "John",
		age:  35,
		date: time.Now(),
	}
	// create a struct without field name
	person1 := Person{"John", 35, time.Now()}
	// create a struct initialize some fields, the rest will be zero value
	person2 := Person{name:"John", age:35}

	fmt.Println("============Struct start ============")
	fmt.Println(person)
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println("============Struct end ============")

}

// Array for Arguments, change the array value in the function
func transfer(numbers [5]int) [5]int {
	numbers[0] = 100
	return numbers
}
