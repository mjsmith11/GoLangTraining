package main

import "fmt"

func main() {
	byteSliceToString()
}

func numerics() {
	x := 12
	y := 12.1230123
	fmt.Println(y + float64(x))
	fmt.Println(int(y) + x)

}

func runes() {
	var x rune = 'a'
	var y int32 = 'b'
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(string(x))
	fmt.Println(string(y))
	//rune to string
}

func byteSliceToString() {
	fmt.Println(string([]byte{'H', 'e', 'l', 'l', 'o'}))
	fmt.Println([]byte("hello"))
}
