package main

import "fmt"

func main() {
	createAMap()
}

func createAMap() {
	words := make(map[string]string)
	words["clock"] = "tells time"
	words["wall"] = "makes a building"
	words["owl"] = "bird"

	fmt.Print(words)
}
