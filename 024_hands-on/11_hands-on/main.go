package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("starting-files/templates/*.gohtml"))
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "")
	if err != nil {
		log.Println(err)
	}
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", "")
	if err != nil {
		log.Println(err)
	}
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", "")
	if err != nil {
		log.Println(err)
	}
}

func apply(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", "")
		if err != nil {
			log.Println(err)
		}
	} else {
		err := tpl.ExecuteTemplate(w, "apply.gohtml", "")
		if err != nil {
			log.Println(err)
		}
	}

}
