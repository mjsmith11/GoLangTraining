package main

import "fmt"

func main() {

}

func stringAsBytes() {
	intro := "Four score and seven years ago..."
	fmt.Println(intro)
	fmt.Println([]byte(intro))

	fmt.Println(len(intro)) //number of bytes not characters.
	fmt.Println(len([]byte(intro)))
}
