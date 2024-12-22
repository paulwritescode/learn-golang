package backend

import (
	"fmt"
	"time"
)

func Backend(name string, myNumRating float64) {
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
