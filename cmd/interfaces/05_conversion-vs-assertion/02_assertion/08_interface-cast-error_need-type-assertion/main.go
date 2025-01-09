package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)           // float64
	fmt.Printf("%T\n", int(rem))      // int

	var val interface{} = 7
	fmt.Printf("%T\n", val)           // interface{} (underlying type int)

	// Type assertion to extract the integer value
	if intVal, ok := val.(int); ok {
		fmt.Printf("%T\n", intVal)    // int
	} else {
		fmt.Println("val is not an int")
	}
}
