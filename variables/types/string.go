package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringConversion()
}

func stringAsBytes() {
	intro := "Four score and seven years ago..."
	bs := []byte(intro)
	fmt.Printf("%T\n", intro)
	fmt.Println(intro)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for _, v := range bs {
		fmt.Printf("%#x\n", v)
	}
}

func length() {
	fmt.Println(len("Hello")) //length is bytes not characters. It matches here because english characters are at the beginning of UTF-8
	fmt.Println(len("世界"))    //Chinese Characters take 3 bytes each
}

func indexAccess() {
	greeting := "Hello" //slice of bytes because that's what strings are
	fmt.Println(greeting)
	fmt.Println(greeting[0])
	fmt.Println(greeting[4])
	fmt.Println("---------------------")
	fmt.Println(greeting[:4])  //beginning up to and not including 4
	fmt.Println(greeting[0:4]) //same as above
	fmt.Println(greeting[1:4]) //IKncludes 1
	fmt.Println(greeting[1:])  //1 to end
}

func stringConversion() {
	if i, err := strconv.Atoi("-42"); err == nil {
		fmt.Printf("%v\t%T\n", i, i)
	}

	s := strconv.Itoa(-42) //Sprint() will do this too
	fmt.Printf("%v\t%T\n", s, s)

}
