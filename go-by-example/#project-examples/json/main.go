package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulwritescode/json/api/example/data"
	"github.com/paulwritescode/json/api/example/encode"
)

func main() {
	fmt.Println("Server Running on Port 3000")
	http.HandleFunc("/", V2)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func V2(w http.ResponseWriter, r *http.Request) {
	myslug := data.Data()
	finalslug := encode.EncodeJson(myslug)
	fmt.Fprint(w, string(finalslug))
}
