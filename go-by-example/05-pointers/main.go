package main

import "fmt"

func main() {
	var value float64 = 99.2
	valueRef := &value
	fmt.Println("This & points to the reference(location address)")
	fmt.Println(valueRef)
	fmt.Println("This * points to the actual value stored in the &location")
	fmt.Println(*valueRef)

	// * always prints the value stored
	// & always prints the addressed a value is stored in

	num := 42
	var ptr *int = &num // ptr now points to num
	fmt.Println(*ptr)
}
