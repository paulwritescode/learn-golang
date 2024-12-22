package frontend

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Frontend() (string, float64) {
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

	return name, myNumRating
}
