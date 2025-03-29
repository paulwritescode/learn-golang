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
	fmt.Println("Hello there, this is the rater project that we are redoing! Damn I had fogotten about golang, but here we go!")

	var name string
	var user_rating string

	// we now take the input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name:")
	name, _ = reader.ReadString('\n') // this return 2 things, so the _ is to ignore it

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating between 1 - 5:")
	user_rating, _ = reader.ReadString('\n') // but this needs to be a number and remove trailing space
	num_rating, _ := strconv.ParseFloat(strings.TrimSpace(user_rating), 64)

	// give the user ouptut
	fmt.Printf("Name: %v\nRating given:%v\nTime recorded:%v\n", name, num_rating, time.Now())

	switch {
	case num_rating == 5:
		fmt.Println("heck yeah! 5")
	case num_rating == 4:
		fmt.Println("we are at 4")
	case num_rating == 3:
		fmt.Println("3")
	case num_rating == 2:
		fmt.Println("2")
	default:
		fmt.Println("1")
	}
}

// what else is there to think about, we are doing this golang thing and it will work
