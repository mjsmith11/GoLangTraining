package main

import "fmt"

/* no default fallthrough
fallthrough is optional
you must explicitly state fallthrough
break statements aren't needed
*/

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("Have you no friends?")
	}

	switch "Jenny" {
	case "Tim", "Jenny": //Is it tim or jenny
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}

	myFriendsName := "Jo"
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Marcus", myFriendsName == "Madhi":
		fmt.Println("Both names start with M")
	}

	switchOnType(4)

}

func switchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting x is of this type
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Unknown")
	}
}
