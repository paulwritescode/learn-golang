/**

SAMPLE PROGRAM THAT TAKES USER INPUT, STORES IN A SLICE AND LOOPS OVER THE SLICE IN OUTPUT


package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var longIn = []string{"Hello", "World"}
	fmt.Println()
	fmt.Println("NB: To Understand slices, first understand arrays")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	longIn = append(longIn, input)
	for _, value := range longIn {
		fmt.Println(value)
	}
	fmt.Println("I know how to take user input in golang")
}

**/

// 1. An array has a length and element type

// ie var name [4]string
// type = string
// lenght = 4

// 2. A dedicatedtslice has the following specification
// []T
// T = type

// 	3. A slice can be created using
// 	. letters := []string {"a"}
// 	. var letters []string
// 	. func make([]T, len, cap) []T
// 	=> When we use the make, it takes a type, length and an optional capacity. It allocates an array and returns a slice that refers to that array!

// 	=> The make() can be used to create both slices and maps;