package main

import "fmt"

func main() {
	var numbers [3]string // to create an array we must set a number in the square brackets [3]

	numbers[0] = "uno"
	numbers[1] = "dos"
	numbers[2] = "tres"

	fmt.Println(numbers)

	var calors = [...]string{"Roho", "gris", "azul", "verde"}
	fmt.Println(calors)
	fmt.Println(len(calors))

	for index, value := range calors {
		fmt.Printf("key %d, value %s\n", index, value)
	}

}

// There is no differnce between these

// var a [3]string                                    // zero-valued
// var b = [3]string{"a", "b", "c"}                  // explicit size
// c := [3]string{"a", "b", "c"}                     // short declaration
// d := [...]string{"a", "b", "c"}                   // let compiler count
