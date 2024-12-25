package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var names = []string{"Tuwei", "Marion", "Chepkoech"}

	fmt.Println("These are the names", names)

	names = append(names, "Mbugua")

	fmt.Println("These are the new names", names)

	names = append(names[1:3]) // This is how we do a slice
	fmt.Println(names)
	heros := make([]string, 3, 3)
	/**
	This is how we create a slice, type,length, capacity
	The legth and capacity does not matter since go recognizes that its a string and it will not be capped!

	When we make() a slice, it allocates/creates an array with length  and capacity but returns a slice.
	**/
	heros[0] = "Scarlet"
	heros[1] = "DeadPool"
	heros[2] = "Human-Torch"
	fmt.Println(heros)
	heros = append(heros, "Wolverine")
	fmt.Println(heros)

	/**
	Sorting is good to do
	. Copy the excisting slice to a new slice and work with the new slice
	**/

	numbers := []int{3, 56, 2, 9, 2}
	isSorted := sort.IntsAreSorted(numbers)
	fmt.Println(isSorted)
	SORTED := make([]int, len(numbers))
	copy(SORTED, numbers)

	// Method 1
	sort.Ints(SORTED)
	fmt.Println("This is the sorted slice using method 1", SORTED)

	// Method 2
	METHOD_2 := make([]int, len(numbers))
	copy(METHOD_2, numbers)
	slices.Sort(METHOD_2)
	fmt.Println("This is the sorted slice using method 2", METHOD_2)
}
