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
	http.HandleFunc("/", pics)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("starting-files/public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pics(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "")
	if err != nil {
		log.Println(err)
	}
}
