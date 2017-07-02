package main

import (
	"fmt"
)

func Greeting(prefix string, who ...string) {
	fmt.Println(prefix)
	for _, value := range who {
		fmt.Println(value)
	}
}

func GreetingS(prefix string, who []string) {
	fmt.Println(prefix)
	for _, value := range who {
		fmt.Println(value)
	}
}

func main() {
	Greeting("nobody")
	Greeting("hello: ", "Joe", "Anna", "Eileen")
	s := []string{"James", "Jasmine"}
	Greeting("goodbye:", s...)
	GreetingS("Goodbye", s)
}
