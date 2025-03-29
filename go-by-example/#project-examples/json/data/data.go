package data

import "time"

type MyData struct {
	Email           string `json:"my_email"`
	CurrentDatetime string `json:"current_datetime"`
	FirstName       string `json:"f_name"`
	LastName        string `json:"last_name"`
}

func Data() MyData {
	today := time.Now().Format("2006")
	SlugDummy := MyData{
		Email:           "kinyattipaul@gmail.com",
		CurrentDatetime: today,
		FirstName:       "Paul",
		LastName:        "Kinyatti",
	}

	return SlugDummy
}
