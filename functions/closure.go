package main

import "fmt"

//closure - one thing enclosing another thing. limit scope
func main() {
	x := 0
	increment := func() int { //increment is enclosed by main as is x. Thus increment can access x
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}
