package main

import (
	"fmt"
)

//functions are types - they can be put in variables, passed to other functions, returned, declared inside other functions

func main() {
	greeting()
	greet("Jane")
	greet("John") //John is the argument
	greet2("John", "Doe")

	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)

	data := []float64{43, 56, 87, 23, 45, 57}
	m := average(data...)
	fmt.Println(m)
}

//no arg
func greeting() {
	fmt.Println("Hello everyone")
}

//1 arg
func greet(name string) {
	//name is the parameter
	fmt.Println(name)
}

//2 arg
func greet2(fname, lname string) { //both are string
	fmt.Println(fname, lname)
}

//with return
func getGreeting(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

//named return not reccomended
func getGreeting2(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

//multiple returns. Typically one value and an error
func getFullNames(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

// variadic
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

//excercise
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
