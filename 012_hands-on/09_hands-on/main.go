package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func close(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//fmt.Printf("Order %v of type %T\n", Order, Order)
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer close(f)

	tcr := csv.NewReader(f)


	records, err := tcr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//for _, rec := range records {
	//	fmt.Println(rec)
	//}

	errr := tpl.Execute(os.Stdout, records)
	if errr != nil {
		fmt.Println(fmt.Errorf("Error is %v", errr))
	}


}
