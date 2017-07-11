package main

import (
	"fmt"
)

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // run at the last possible moment before the calling function returns // useful for closing files, releasing locks etc
	hello()
}

//f, err := os.Open(filename)
//if err != nil {
//	return "", err
//}
// defer f.Close()
//// do something useful with f and it will be closed when you return no matter the path taken, fall off the end of method, or corresponding gorouting is panicking
