package main

import "fmt"

func main() {
	useCallback()

}
func useMakeEvenGen() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
	fmt.Println(nextEven()) //6

	masEven := makeEvenGenerator()
	fmt.Println(masEven()) //2
	fmt.Println(masEven()) //4
	fmt.Println(masEven()) //6
}

func useCallback() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

func makeEvenGenerator() func() int { //returns a func that returns an int
	i := 0 // i is enclosed by make even generator
	return func() int {
		i += 2
		return i
	}
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
