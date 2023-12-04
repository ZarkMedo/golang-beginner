// Reference Type
// 1. Slice
// 2. Channel
// 3. Map
// 4. Pointer
// 5. Function/Method

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Slice
	// create a slice
	slice := []string{"a", "b", "c", "d", "e"}
	// create a slice with make function
	slice1 := make([]string, 5)
	// add all of the slice's elements to slice1
	copy(slice1, slice)
	fmt.Println("============Slice start ============")
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println("============Slice end ============")
	// slice for Arguments
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("============Slice for Arguments start ============")
	// check the pointer of the numbers and value
	fmt.Printf("numbers pointer: %p, numbers value: %v\n", &numbers, numbers)
	// call the transfer function
	newNumbers := transferSlice(numbers)
	// check the pointer of the numbers and value
	fmt.Printf("numbers pointer: %p, numbers value: %v\n", &numbers, numbers)
	fmt.Printf("newNumbers pointer: %p, newNumbers value: %v\n", &newNumbers, newNumbers)
	fmt.Println("============Slice for Arguments end ============")



	// difference between array and slice
	// array
	array := [5]string{"a", "b", "c", "d", "e"}
	// slice
	slice2 := make([]string, 5) // <==> slice2 := []string{"a", "b", "c", "d", "e"}
	copy(slice2, array[:])
	fmt.Println("============Array and Slice start ============")
	fmt.Printf("array type: %v, slice type: %v\n", reflect.TypeOf(array).Kind(), reflect.TypeOf(slice2).Kind()) // reflect.TypeOf(array).Kind() return the type of the array/slice
	fmt.Printf("array value: %v, slice value: %v\n", array, slice2)
	fmt.Println("============Array and Slice end ============")


	// Channel
	// create a channelUnbuffered
	channelUnbuffered := make(chan string) // if you don't specify the buffer, it will be unbuffered channel, it means it will be blocked until the value is received, it is synchronous, must use goroutine to send and receive value
	// create a channel with buffer
	channelBuffered := make(chan string, 5) // just accept 5 values, if more than 5 values, it will be blocked
	// send a value to the channel, use the goroutine
	go sendEvent(channelUnbuffered, "a")
	// receive a value from the channel
	value := <-channelUnbuffered

	// send a value to the channel, also you can use goroutine
	channelBuffered <- "d" // if you want, go sendEvent(channelBuffer, "d")
	// receive a value from the channel
	value1 := <-channelBuffered
	close(channelUnbuffered)
	close(channelBuffered)
	fmt.Println("============Channel start ============")
	fmt.Printf("value: %v\n", value)
	fmt.Printf("value1: %v\n", value1)
	fmt.Println("============Channel end ============")

	// Range over channel
	channel := make(chan string, 5)
	done := make(chan bool)
	// receive a value from the channel with range
	go func() {
		for value := range channel {
			fmt.Printf("value: %v\n", value)
		}
		done <- true
	}()
	// send a value to the channel
	channel <- "a"
	channel <- "b"
	channel <- "c"
	channel <- "d"
	channel <- "e"
	// close the channel
	close(channel)
	// wait until the goroutine is done
	<-done
}

// transfer function, accept an slice and return a new slice
func transferSlice(numbers []int) []int {
	newNumbers := make([]int, len(numbers))
	copy(newNumbers, numbers)
	return newNumbers
}

// sendEvent function, accept a channel and send a value to the channel
func sendEvent(channel chan string, message string) {
	channel <- message
}