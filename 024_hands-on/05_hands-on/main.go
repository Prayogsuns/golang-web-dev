package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func main() {
	//http.Handle("/", http.FileServer(http.Dir("./starting-files")))
	http.Handle("/pics/", http.FileServer(http.Dir("starting-files/public")))
	http.HandleFunc("/", dogs)


	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "")
	if err != nil {
		log.Println(err)
	}

}
