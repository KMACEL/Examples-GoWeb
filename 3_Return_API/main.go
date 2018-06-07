package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Home struct {
	Home string `json:"msg"`
}

//Information is
type Information struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	BirthDate string `json:"birthDay"`
	LiveCity  string `json:"liveCity"`
}

func main() {
	api := "/api"

	http.HandleFunc(api, func(w http.ResponseWriter, r *http.Request) {
		home := Home{Home: "WELCOME GOLANG"}
		convert, _ := json.Marshal(home)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(convert))
	})

	http.HandleFunc(api+"/info", func(w http.ResponseWriter, r *http.Request) {
		info := Information{Name: "Mert", Surname: "Acel", BirthDate: "15.02.1994", LiveCity: "Busa"}
		convert, _ := json.Marshal(info)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(convert))
	})

	http.ListenAndServe(":8080", nil)
}
