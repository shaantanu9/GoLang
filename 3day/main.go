// Data Types in Go Lang

package main

import "fmt"

func main() {

	// Integer
	var a int = 10
	fmt.Printf("%d\n", a) // %d is for integer
	// Float
	var b float32 = 10.5
	fmt.Printf("%f\n", b) // %f is for float
	// String
	var c string = "Hello World"
	fmt.Printf("%s\n", c) // %s is for string
	// Boolean
	var d bool = true
	fmt.Printf("%t\n", d) // %t is for boolean
	// Character
	var e rune = 'a'
	fmt.Printf("%c\n", e) // %c is for character
	// Array
	var f [3]int = [3]int{1, 2, 3} //[3]int is for array of 3 integers
	fmt.Printf("%v\n", f)          // %v is for array
	// Slice
	var g []int = []int{1, 2, 3} // []int is for slice of integers
	fmt.Printf("%v\n", g)        // %v is for slice
	// Map
	var h map[string]int = map[string]int{"a": 1, "b": 2} // map[string]int is for map of string and integer
	fmt.Printf("%v\n", h)                                 // %v is for map
	// Pointer
	var i *int = new(int) // *int is for pointer to integer
	fmt.Printf("%v\n", i) // %v is for pointer
	// Function
	var j = func() { // func() is for function
		fmt.Println("Hello World")
	}
	j() // %v is for function
	// Interface
	var k interface{} = "Hello World" // interface{} is for interface
	fmt.Printf("%v\n", k)             // %v is for interface type, interface is a type that can hold any type of value like string, int, float, etc.
	// Struct
	var l struct { // struct is for struct
		a int
		b string
	}
	l.a = 1
	l.b = "Hello World"
	fmt.Printf("%v\n", l) // %v is for struct
	fmt.Println(l)        // %v is for struct

	// Constants
	const (
		a1 = 1
		b1 = 2
	)
	fmt.Printf("%d\n", a1) // %d is for integer
	// %d is use for integer
	// and %f is use for float and %s is use for string
	//and %t is use for boolean
	// and %c is use for character
	// and %v is use for all other data types
	// and %T is use for data type
	// and %#v is use for interface type
	// and %#T is use for data type
	// and %#v is use for struct type
	// and %#v is use for function type
	// and %#v is use for pointer type
	// and %#v is use for array type
	// and %#v is use for slice type
	// and %#v is use for map type
	// and %#v is use for interface type
	// and %#v is use for struct type
}
