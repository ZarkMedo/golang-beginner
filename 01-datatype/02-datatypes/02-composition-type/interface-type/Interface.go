// Interface is a type that defines a set of methods.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Interface
	// create a interface
	interface1 := interface{}("Hello World")
	fmt.Println("============Interface start ============")
	fmt.Printf("interface1 type: %v, interface1 value: %v\n", reflect.TypeOf(interface1).Kind(), interface1)
	fmt.Println("============Interface end ============")

	// use the interface for struct
	// create a rect
	rect := rect{
		width:  10,
		height: 5,
	}
	// create a circle
	circle := circle{
		radius: 5,
	}
	// call the measure function
	measure(rect)
	measure(circle)

}

// define the interface for struct

type geometry interface {
	area() float64
	perim() float64
}

// define a rect
type rect struct {
	width, height float64
}

// define a circle
type circle struct {
	radius float64
}

// define a function on rect
func (r rect) area() float64 {
	return r.width * r.height
}

// define a function on rect
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// define a function on circle
func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

// define a function on circle
func (c circle) perim() float64 {
	return 2 * c.radius * 3.14
}

// define a function to measure the geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
