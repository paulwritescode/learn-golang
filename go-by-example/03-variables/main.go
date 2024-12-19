package main

import "fmt"

// wehn outside the function scope initialize variable with keyword var
var username = "manqi qode"
var isLogged = false
var age = 21

func main() {
	// String
	fmt.Println("We are working with variables.")
	fmt.Printf("First Variable(string): %s \n", username)
	// NB: There is a diffrence between the two methods of print

	// Boolean
	fmt.Printf("Second variable(boolean): %t \n", isLogged)

	// short-hadn initialisation of variable :=
	isLogged := true
	fmt.Printf("Short-hand declaration of variable( := ): %t", isLogged)

	// integer
	fmt.Printf("Third Variable (integer): %d \n", age)
	fmt.Printf("Doing addition (integer): %d \n", age+1)
	age := age + 1
	fmt.Printf("Short-hand declaration (integer): %d \n", age)
	// uint
	var smallnumber uint8 = 255
	fmt.Printf("uint8: %d \n", smallnumber)

	// float
	var smallFloat float32 = 255.45676543456787654323456
	fmt.Printf("float32: %f \n", smallFloat)
	var smallFloat1 float64 = 255.45676543456787654323456
	fmt.Printf("float64: %f \n", smallFloat1)

	// default valuse and aliases
	fmt.Println()
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("This variable of type: %T \n", anotherVariable)
	fmt.Println()
	var anotherVariable1 string
	fmt.Println(anotherVariable1)
	fmt.Printf("This variable of type: %T \n", anotherVariable1)
	// The default value that goes into the variable

}
