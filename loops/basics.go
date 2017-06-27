package main

import "fmt"

func main() {
	loopOverString()
}

func countTo100() {
	// like a while
	// for condition {}

	// infinite loop
	// for {}

	// to loop over an array, slice, string, or map or reading from a channel
	// for key, value := range myMap

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func loopOverString() {
	myPhrase := "What is a good phrase?"
	for key, value := range myPhrase {
		fmt.Println(key, " - ", string(value))
	}
}
