package main

import "fmt"

func main() {
	createAMap()
}

func createAMap() {
	words := make(map[string]string, 5) // 5 is a capacity hint
	// words := map[int]string
	words["clock"] = "tells time"
	words["wall"] = "makes a building"
	words["owl"] = "bird"

	// for k, v := range words

	fmt.Println(words)

	delete(words, "owl")

	fmt.Println(words)

	if val, exists := words["2"]; exists {
		//do something only if there is something at index "2"
	} else {
		// val would be zero value here
	}
}
