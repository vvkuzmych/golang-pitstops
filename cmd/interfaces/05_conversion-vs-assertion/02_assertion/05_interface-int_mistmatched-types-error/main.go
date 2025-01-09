package main

import "fmt"

func main() {
	var val interface{} = 7

	// Type assertion to extract the integer value
	intVal, ok := val.(int)
	if ok {
		fmt.Println(intVal + 6)
	} else {
		fmt.Println("val is not an integer")
	}
}
