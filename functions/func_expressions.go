package main

import "fmt"

func main() {
	greeting := func() { // scope of greeting is func main
		fmt.Println("Hello World!")
	}
	greeting()

	fmt.Printf("%T\n", greeting)

	half := func(n int) (int, bool) { //anonymous function - doesn't have name
		return n / 2, n%2 == 0
	}

	fmt.Println(half(2))
	fmt.Printf("%T\n", half)
}
