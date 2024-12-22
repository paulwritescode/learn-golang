package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	// frontend
	// take user name
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name: ")
	name, _ = reader.ReadString('\n')

	// take rating and convert to int/float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate us between 1 and 5")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	// backend
	fmt.Println("###################################################")
	fmt.Printf("Hello %v\nThanks for rating us with %v rating \n\nYour rating was recorded at %v\n\n", name, myNumRating, time.Now().Format(time.Stamp))

	if myNumRating == 5 {
		fmt.Println("Bonus for team for 5 star service")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("We are alwsys improving")
	} else if myNumRating < 3 {
		fmt.Println("Need serious work on our side")
	}
	fmt.Println("###################################################")
}
