package main

import "fmt"

const p = "death and taxes"

const (
	A = iota //0
	B = iota //1
	C = iota //2
	//iota means it increments
)

const (
	D = iota //0
	E = iota //1
	F = iota //2
	//iota means it increments
)

const (
	_ = iota      //0
	G = iota * 10 // 1 * 10
	H = iota * 20 // 2 * 10

)

const (
	_  = iota             //0
	KB = 1 << (iota * 10) // 1 << (1 * 10) bitwise ops
	MB = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {

	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
