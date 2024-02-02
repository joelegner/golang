package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/joe", joeHandler)
	http.HandleFunc("/template", templateHandler)
	fmt.Println("Browse to http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, html(head(), body("Hello, World!")))
}

func joeHandler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, html(head(), body("Joe Legner is Cool!")))
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

func html(head string, body string) string {
	return fmt.Sprintf(`
<!doctype html>
<html lang="en">
%s
%s
</html>
`, head, body)
}

func head() string {
	return `
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Bootstrap demo</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
    <link rel="stylesheet" href="styles.css">
</head>
`
}

func body(message string) string {
	return fmt.Sprintf(`
<body>
<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
  <div class="container">
    <a class="navbar-brand" href="/">Intranet</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="/joe">Joe is Co¸ol</a>
        </li>
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="/template">Template</a>
        </li>
        <li class="nav-item">
          <a class="nav-link active" href="https://joelegner.github.io/themanual">The Manual</a></li>
        </li>
      </ul>
      <form class="d-flex" role="search">
        <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
        <button class="btn btn-outline-success" type="submit">Search</button>
      </form>
    </div>
  </div>
</nav>

<div class="container my-5">
  <h1>%s</h1>
  <div class="col-lg-8 px-0">
    <p class="fs-5">Welcome to Joe Legner's Intranet site he build with <a href="https://go.dev">Go</a>.</p>
  </div>
</div>

<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>
<script src="main.js"></script>
</body>
`, message)
}
