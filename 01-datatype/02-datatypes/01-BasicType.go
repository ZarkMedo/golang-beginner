// primative data types in go:
// 1. bool
// 2. string
// 3. int int8 int16 int32 int64
// 4. uint uint8 uint16 uint32 uint64 uintptr
// 5. float32, float64
// 6. complex64, complex128
// 7. byte
// 8. rune

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// bool
	var isTrue bool = true // <==> isTrue := true
	fmt.Println(isTrue)

	// string
	var name string = "John"         // <==> name := "John"
	var nameOneQuote string = `John` // <==> nameOneQuote := `John`, raw string literal
	fmt.Println("============string start ============")
	fmt.Println(name)
	fmt.Println(nameOneQuote)
	fmt.Println("============string end ============")

	// int Signed
	var age int = 35 // <==> age := 35 // int is 32 or 64 bit depends on the system
	fmt.Println("============int start ============")
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))
	fmt.Println("============int end ============")
	// uint Unsigned
	var ageUnsigned uint = 35 // <==> ageUnsigned := 35, just accept positive number
	fmt.Println("============uint start ============")
	fmt.Println(ageUnsigned)
	fmt.Println(reflect.TypeOf(ageUnsigned))
	fmt.Println("============uint end ============")

	// float 32
	var weight float32 = 80.5 // !equals  weight := 80.5, because 80.5 is float64
	weightNoation := 80.5
	fmt.Println("============float32 start ============")
	fmt.Println(weight)
	fmt.Println(reflect.TypeOf(weight))
	fmt.Println(reflect.TypeOf(weightNoation))
	fmt.Println("============float32 end ============")

	// float64
	var weight64 float64 = 180.5 // <==> height := 180.5
	weight64Noation := 180.5     // <==> heightNoation := 180.5
	fmt.Println("============float64 start ============")
	fmt.Println(weight64)
	fmt.Println(reflect.TypeOf(weight64))
	fmt.Println(reflect.TypeOf(weight64Noation))
	fmt.Println("============float64 end ============")

	// complex
	var complexNumber complex64 = 1 + 2i // <==> complexNumber := 1 + 2i
	fmt.Println(complexNumber)

	// byte
	var byteValue byte = 255 // <==> byte is alias for uint8, meant to store ascii characters
	fmt.Println("============byte start ============")
	fmt.Println(byteValue)
	fmt.Printf("byteValue type: %T, bytes size: %d\n", byteValue, unsafe.Sizeof(byteValue))
	fmt.Println("============byte end ============")

	// rune
	var runeValue rune = 255 //  rune is alias for int32, meant to store unicode characters
	fmt.Println("============rune start ============")
	fmt.Println(runeValue)
	fmt.Printf("runeValue type: %T, bytes size: %d\n", runeValue, unsafe.Sizeof(runeValue))
	fmt.Println("============rune end ============")

	// uintptr
	var uintptrValue uintptr = 255 // uintptr is an unsigned integer large enough to store the uninterpreted bits of a pointer value
	// on 32 bit system, uintptr is 32 bits or 4 bytes
	// on 64 bit system, uintptr is 64 bits or 8 bytes
	fmt.Println("============uintptr start ============")
	fmt.Println(uintptrValue)
	fmt.Printf("uintptrValue type: %T, bytes size: %d\n", uintptrValue, unsafe.Sizeof(uintptrValue))
}
