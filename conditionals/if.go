package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This code ran")
	}

	if false {
		fmt.Println("This code did not run")
	}

	b := true
	if food := "Chocolate"; b { //initialization; condition - often used to say do this and if ok run this block ("comma ok" idiom)
		// food has scope for only this if statement
		fmt.Println(food)
	}
}
