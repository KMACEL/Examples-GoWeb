package main

import (
	"html/template"
	"net/http"
)

type PageElem struct {
	Title       string
	Name        string
	Message     string
	Information []Information
}

type Information struct {
	Name      string
	Surname   string
	BirthDate string
	BirthCity string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := PageElem{
		Title:   "Test",
		Name:    "Mert",
		Message: "Hello Word",
		Information: []Information{
			Information{
				Name:      "Mert",
				Surname:   "Acel",
				BirthDate: "15.02.1994",
				BirthCity: "Bursa"},
			Information{
				Name:      "Kübra",
				Surname:   "Acel",
				BirthDate: "25.02.1993",
				BirthCity: "Uşak"},
			Information{
				Name:      "Bilge",
				Surname:   "Acel",
				BirthDate: "26.6.2028",
				BirthCity: "Ankara"},
			Information{
				Name:      "Erdem Uraz",
				Surname:   "Acel",
				BirthDate: "05.01.2018",
				BirthCity: "Kocaeli"}}}

	t, _ := template.ParseFiles("test.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
