package main

import "fmt"

func modifyArray(array *[5]int) {
	array[3] = 300
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	modifyArray(&array)
	fmt.Println(array)
	fmt.Println("modification affects the same array because we are pointing to the original array to be modified using the &")
}
