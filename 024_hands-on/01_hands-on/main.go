package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", func(w http.ResponseWriter, req *http.Request){
		http.ServeFile(w, req, "dog.jpg")
	})


	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Foo ran")
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "")
	if err != nil {
		log.Println(err)
	}
}
