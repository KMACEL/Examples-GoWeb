package main

import (
	"html/template"
	"net/http"
)

type PageElem struct {
	Title   string
	Name    string
	Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := PageElem{
		Title:   "Test",
		Name:    "Mert",
		Message: "Hello Word"}
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
