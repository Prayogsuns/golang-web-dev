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

type rH string
type dH string
type mH string

var rootH rH
var dogH dH
var meH mH

func (rh rH) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	err := ghtm.Execute(w, "Index")
	if err != nil {
		log.Println(err)
	}
}
func (dh dH) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	err := ghtm.Execute(w, "Dog")
	if err != nil {
		log.Println(err)
	}
}
func (mh mH) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	err := ghtm.Execute(w, "Me")
	if err != nil {
		log.Println(err)
	}
}

func main() {


	http.Handle("/", rootH)
	http.Handle("/dog/", dogH)
	http.Handle("/me/", meH)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
