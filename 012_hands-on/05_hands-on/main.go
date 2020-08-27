package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

type Menu struct {
	Bf []Breakfast
	Lch []Lunch
	Dr []Dinner
}

type Breakfast struct {
	Item string
	Num int
}

type Lunch struct {
	Item string
	Num int
}

type Dinner struct {
	Item string
	Num int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	Order := Menu{
			[]Breakfast{
				Breakfast{"Sandwich",2},
				Breakfast{"Orange Juice", 4},
			},
			[]Lunch{
				Lunch{"Norht Indian Thali", 2},
			},
			[]Dinner{
				Dinner{"Veg Biryani", 2},
			},
		}

	//fmt.Printf("Order %v of type %T\n", Order, Order)

	err := tpl.Execute(os.Stdout, Order)
	if err != nil {
		fmt.Println(fmt.Errorf("Error is %v", err))
	}

}
