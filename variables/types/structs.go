package main

import "fmt"

type person struct { // unexported
	name string
	age  int
}

// *** cannot use make with a struct
func main() {
	p1 := person{"James", 20}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	fmt.Println("-------------------")

	p2 := person{name: "Matt", age: 24}
	fmt.Println(p2)

	fmt.Println("-------------------")

	p3 := &person{name: "Joe"} // creating without age gets 0 for age
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)

	fmt.Println("-------------------")

	p4 := new(person) //pointer to 0 value person
	fmt.Println(p4)
	p4.name = "Ralph"
	p4.age = 27
	fmt.Println(p4)
	fmt.Printf("%T\n", p4)
	fmt.Println(p4.name)
	fmt.Println(p4.age)
}
