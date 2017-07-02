package main

import "fmt"

func main() {

}

func stringAsBytes() {
	intro := "Four score and seven years ago..."
	fmt.Println(intro)
	fmt.Println([]byte(intro))

	fmt.Println(len(intro))
	fmt.Println(len([]byte(intro)))
}
