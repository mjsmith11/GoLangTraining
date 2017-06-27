package main

import "fmt"
import "reflect"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	//equivalent
	//var b = &a
	//b := &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Printf("%T\n", b)

	*b = 42
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a)

	//b is a pointer to the memory address where an int is stored
	//b's type is int pointer

}
