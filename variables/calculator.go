package main

import "fmt"

func main() {
	var numOne, numTwo int
	fmt.Print("Enter a number: ")
	fmt.Scan(&numOne)
	fmt.Print("Enter a second number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne * numTwo)

}
