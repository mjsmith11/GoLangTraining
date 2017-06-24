package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // & is memory address of
	fmt.Println("Hello ", name)
}