package main

import (
	"fmt"
	"net/http"
)

// Information is
type Information struct {
	Name      string
	Surname   string
	BirthDate string
	LiveCity  string
}

func (i Information) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i.Name = "Mert"
	i.Surname = "Acel"
	i.BirthDate = "15.02.1994"
	i.LiveCity = "Bursa"

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path : ", r.URL.Path[1:])
	fmt.Fprint(w, "<table border=\"1\"><tr><td >İsim</td><td >Soyisim</td><td>Doğum Tarihi</td><td >Doğum Yeri</td></tr><tr><td>"+i.Name+"</td><td >"+i.Surname+"</td><td>"+i.BirthDate+"</td><td>"+i.LiveCity+"</td></tr></table>")
}

func main() {
	var info Information
	http.ListenAndServe("localhost:8080", info)
}
