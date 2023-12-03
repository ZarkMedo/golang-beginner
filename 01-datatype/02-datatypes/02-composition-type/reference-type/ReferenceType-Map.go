// section of ReferenceType.go -> Map section

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Map
	// create a map
	map1 := map[string]string{"a": "1", "b": "2", "c": "3"}
	// create a map with make function
	map2 := make(map[string]string)
	// add all of the map1's elements to map2
	for key, value := range map1 {
		map2[key] = value
	}
	fmt.Println("============Map start ============")
	fmt.Println(map1)
	fmt.Println(map2)
	// fmt.Printf("map1 equals map2: %v\n", map1 == map2) // map can't be compared with ==, it will be compared with the pointer
	fmt.Printf("map1 equals map2: %v\n", reflect.DeepEqual(map1, map2)) // use reflect.DeepEqual to compare the map
	map2["d"] = "4"
	fmt.Printf("map1 equals map2: %v\n", reflect.DeepEqual(map1, map2)) // use reflect.DeepEqual to compare the map
	fmt.Println("============Map end ============")

	// how to loop the map
	// 1. use for range
	fmt.Println("============Map for range start ============")
	for key, value := range map1 {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	// Map also is Set in Go
	// create a set
	set := map[string]bool{"a": true, "b": true, "c": true} // the value is bool, it is Set
	fmt.Println("============Set start ============")
	fmt.Println(set)
	fmt.Println("============Set end ============")
}
