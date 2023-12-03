// section of ReferenceType.go -> Pointer section

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Pointer
	// create a pointer
	pointer := new(int) // it will return a pointer, the value is 0
	fmt.Println("============Pointer start ============")
	fmt.Printf("pointer type: %v, pointer value: %v\n", reflect.TypeOf(pointer).Kind(), pointer)
	// create a pointer with &
	number := 1
	pointer2 := &number
	fmt.Printf("pointer2 type: %v, pointer2 value: %v\n", reflect.TypeOf(pointer2).Kind(), pointer2)
	fmt.Println("============Pointer end ============")

	// Pointer for Arguments, change the value by pointer
	fmt.Println("============Pointer for Arguments start ============")
	number1 := 1
	pointer3 := &number1
	fmt.Printf("number1 pointer: %p, number1 value: %v\n", pointer3, *pointer3) // check the pointer of the number and value
	returnValue := transferPointer(pointer3)

	fmt.Printf("pointer3 pointer: %p, pointer3 value: %v\n", &pointer3, *pointer3) // check the pointer of the number and value
	fmt.Printf("returnValue pointer: %p, returnValue value: %v\n", &returnValue, *returnValue) // check the pointer of the number and value
	fmt.Println("============Pointer for Arguments end ============")

}

// Pointer for Arguments function
func transferPointer(number *int) *int {
	*number = 2
	return number
}