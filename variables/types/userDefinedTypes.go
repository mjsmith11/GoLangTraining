package main

import (
	"fmt"
)

type mySentence string // type mySentence with underlying type string

func main() {
	var message mySentence = "Hello World!"
	fmt.Println(message)
	fmt.Printf("%T\n", message)
	fmt.Printf("%T\n", string(message))
}

/* creating a type only to alias another isn't really something you would do
It is done if you need to attach a method to the type
*/
