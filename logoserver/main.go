package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var content string = `
	<h1>List of URLs in logos</h1>
	<ul>
	<li><a href="https://yahoo.com">Coffee</a></li>
	<li>Tea</li>
	<li>Milk</li>
  	</ul>`
    fmt.Fprintf(w, html(head(""), body(content)))
}

func main() {
	fmt.Println("logoserver 0.1")
	fmt.Println("Serving on http://localhost.com:8080...")
	fmt.Println("Use CTRL+C to stop the server.")
	http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
