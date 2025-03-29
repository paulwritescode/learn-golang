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
	// & always prints the address a value is stored in

	num := 42
	var ptr *int = &num // ptr now points to num
	fmt.Print(ptr)

	// Question 1
	x := 24
	p := &x
	fmt.Println("Quiz: What is stored in the following")
	fmt.Printf("x = %d\n", x)
	fmt.Printf("p = %d\n", p)
	fmt.Printf("*p = %d\n", *p)
	/*
		Answer
		x= 24
		p is storing the address of x (&)
		*p (dereferencing p gives us the value that is stored in the address that p stores (&x))
	*/

	// Question 2

	*p = 100
	fmt.Println("Quiz 2: After *p = 100, what is the value of x?")
	fmt.Printf("x = %d\n", x)
	fmt.Printf("p= %d\n", p)

	/*
		p holds &x the address to x
		*p holds the value in x
		changeing *p changes the value in x
		so....
		x=32
		p = &x
		*p = 32 (x=32)
		*p = 100 will change x from 32 to 100
	*/

	//question 3
	type Person struct {
		name string
		age  int
	}

	person := Person{"Alice", 30}
	personPtr := &person

	fmt.Println("Quiz 3: How do we access the name field in the personPtr?")
	fmt.Printf("using dot notation: %s\n", personPtr.name)
	fmt.Println(person)
	fmt.Println(personPtr)
	fmt.Printf("Another way to print the name: using dereferencing (*personPtr).name %s\n", (*personPtr).name)

	//question 4

	a := [4]int{1, 2, 4, 5}
	b := &a[0]
	c := &a[1]

	fmt.Println(a)
	fmt.Printf("b= %v\nc= %v\n", b, c)

	//question 5
	var ptr_new *int
	fmt.Printf("ptr = %v\n", ptr_new)
}
