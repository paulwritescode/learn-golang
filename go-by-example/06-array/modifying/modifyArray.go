package main

import "fmt"

func modArray(array *[5]int) {
	array[4] = 58093
}

func main() {
	array := [5]int{2, 3, 4, 5, 6}
	fmt.Println(array)
	modArray(&array)
	fmt.Println(array)
}

/*
using the pointer to avoind changing a copy of the array
*/
