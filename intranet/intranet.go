package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	// Add some routes
	http.HandleFunc("/", templateHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/aci318-14", aci318_14Handler)

	// Give instructions
	fmt.Println("Browse to http://localhost:8080")

	// Start listening for requests
	http.ListenAndServe(":8080", nil)
}

func aci318_14Handler(w http.ResponseWriter, r *http.Request) {
	// Open file
	f, err := os.Open("resources/aci318-14.pdf")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	//Set header
	w.Header().Set("Content-type", "application/pdf")

	//Stream to response
	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl, _ := template.ParseFiles("templates/main.html")
	tmpl.Execute(w, data)
}
