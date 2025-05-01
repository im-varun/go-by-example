package main

import "fmt"

func returnOneValue() int {
	return 1
}

func returnTwoValues() (int, int) {
	return 2, 3
}

func returnThreeValues() (int, int, int) {
	return 4, 5, 6
}

func main() {
	a := returnOneValue()
	fmt.Println(a)

	b, c := returnTwoValues()
	fmt.Println(b, c)

	_, d, e := returnThreeValues()
	fmt.Println(d, e)
}
