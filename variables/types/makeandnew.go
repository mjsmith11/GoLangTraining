package main

import (
	"fmt"
)

func main() {
	var age *[]int = new([]int) // new returns a pointer to zero value which is nil for slices
	fmt.Println(age)
	fmt.Println(*age)
	*age = append(*age, 42)
	fmt.Println(age)
	fmt.Println(*age)

	var bday []int = make([]int, 10, 100) //make is only for slice, map, channel. puts zeros in all the spots
	fmt.Println(bday)
	bday[7] = 42
	fmt.Println(bday)
}
