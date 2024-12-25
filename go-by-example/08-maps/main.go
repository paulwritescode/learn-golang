/**

NOTES
. This is key:value pair
. We use:
		make new
. new => only allocates , there is no initialisation of memory
. make => allocates and initialize memory || No zero memory



**/

package main

import "fmt"

func main() {
	var SCORE = make(map[string]int) // Using this structure is using NEW
	SCORE["Math"] = 100
	SCORE["Eng"] = 86
	SCORE["Scie"] = 80
	SCORE["Scie"] = 50
	SCORE["Scie"] = 84
	fmt.Println(SCORE)

	// we will use the key
	// To access particular value
	GET_MATH := SCORE["Math"]
	fmt.Println("This is the Math:", GET_MATH)

	// To delete
	delete(SCORE, "Scie")
	fmt.Println("We deleted something:", SCORE)

	// loops (for)
	for key, value := range SCORE {
		fmt.Printf("Score of %v is %v\n", key, value)
	}

	// loops (ommiting the key)
	for _, value := range SCORE {
		fmt.Printf("Score is %v\n", value)
	}

}
