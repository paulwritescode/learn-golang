package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// When you take input from the user it is always in string format

func main() {

	// var myInput string
	// fmt.Println("We are taking user input from terminal!")
	// fmt.Scanln(&myInput)
	// fmt.Println(&myInput)

	// var name string = "Paul"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter full name: ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	// we first have to create our reader

	reader := bufio.NewReader(os.Stdin)
	// prompt the user to enter something
	fmt.Println("Enter your name:")
	// capture the user input to a variable using the reader
	// the reader returns two items, error and the input
	my_name, _ := reader.ReadString('\n')
	fmt.Println(my_name)

	// reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Rating: ")
	myRating, _ := reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println(myNumRating + 2)
}
