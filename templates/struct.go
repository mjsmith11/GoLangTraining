package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person        //Embedded type-Double zero is outer type. Person is inner type. Everything in Person gets Promoted to DoubleZero
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	} //all fields can be accessed directly from outer type variable

	tpl, err := template.ParseFiles("structtpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
