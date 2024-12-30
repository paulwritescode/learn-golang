package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world! Server on port 3000")
	// 2.Define a server : url + handler function
	h1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Wakanda Forever", Director: "Rusell Brothers"},
				{Title: "Avengers", Director: "Marvel"},
				{Title: "Bleach", Director: "Tite Kubo"},
			},
		}
		templ.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		fmt.Println("Value added to DB")
		htmlStr := fmt.Sprintf("<p class='bg-sky-500 px-4 py-2 text-white rounded-md'>%s-%s</p>", title, director)
		templ, _ := template.New("t").Parse(htmlStr)
		templ.Execute(w, nil)
		// Look into using templates rather than hardcoding the html in the response
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	// 1. Create server
	log.Fatal(http.ListenAndServe(":3000", nil))

}
