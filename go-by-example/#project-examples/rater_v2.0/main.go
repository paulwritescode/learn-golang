package main

import (
	"fmt"

	"github.com/paulwritescode/rater/v2/backend"
	"github.com/paulwritescode/rater/v2/frontend"
)

func main() {
	fmt.Println("Welcome to our Rater")
	name, myNumRating := frontend.Frontend()
	backend.Backend(name, myNumRating)
}
