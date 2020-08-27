package main

import (
	"log"
	"os"
	"text/template"
)

type Address struct {
	Location string
	City string
	Zip int
	Region string
}

type Hotel struct {
	Name string
	Address
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := []Hotel{
			Hotel{"X", Address{"UpTown", "California", 110001, "Northern"}},
			Hotel{"Y", Address{"MidTown", "California", 110002, "Central"}},
			Hotel{"Z", Address{"DownTown", "California", 110003, "Southern"}},
		}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
