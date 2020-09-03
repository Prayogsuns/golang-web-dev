package main

import (
	"html/template"
	//"fmt",
	"log"
	"net/http"
)


var ghtm *template.Template

func init() {
	ghtm = template.Must(template.ParseFiles("ghtm.gohtml"))
}


func main() {

	rootH := func(w http.ResponseWriter, res *http.Request) {
		err := ghtm.Execute(w, "Index")
		if err != nil {
			log.Println(err)
		}
	}
	dogH := func(w http.ResponseWriter, res *http.Request) {
		err := ghtm.Execute(w, "Dog")
		if err != nil {
			log.Println(err)
		}
	}
	meH := func(w http.ResponseWriter, res *http.Request) {
		err := ghtm.Execute(w, "Me")
		if err != nil {
			log.Println(err)
		}
	}

	http.HandleFunc("/", rootH)
	http.HandleFunc("/dog/", dogH)
	http.HandleFunc("/me/", meH)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
