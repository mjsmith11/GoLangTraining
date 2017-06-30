package main

import "fmt"

func main() {
	evensTo1000()
}

func countTo100() {
	// like a while
	// for condition {

	//}

	// infinite loop
	// for {

	//}

	// to loop over an array, slice, string, or map or reading from a channel
	// for key, value := range myMap

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func loopOverString() {
	//the value is a rune which is an alias for int32 used when it is representing a UTF-8 character

	myPhrase := "What is a good phrase?"
	for key, value := range myPhrase {
		fmt.Println(key, " - ", string(value))
	}

	//for i := 0; i < len(myPhrase); i++ {
	//	fmt.Println(i, " - ", string(myPhrase[i]))

	//}
}

func evensTo1000() {
	for i := 2; i <= 1000; i += 2 {
		fmt.Println(i)
	}
}
