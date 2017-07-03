package main

import "fmt"

func main() {
	stringAsBytes()
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
