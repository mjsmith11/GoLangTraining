package main

import "fmt"

func main() {
	sliceAppend()
}

func createSlices() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2])

	friends := []string{"Rio", "Kai", "Bob"}
	printer(friends...)
}

func printer(str ...string) {
	for k, v := range str {
		fmt.Println(k, v)
	}
}

func makeASlice() {
	customerNumber := make([]int, 3) // 3 is length and capacity

	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length = number of elements refered to by the slice
	// 5 is the capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	greeting = append(greeting, "suprabadham") //to grow the slice

	fmt.Println(greeting[3])
}

func sliceAppend() {
	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}
