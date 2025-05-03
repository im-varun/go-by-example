package main

import "fmt"

func main() {
	const s = "golang"

	fmt.Println("len:", len(s))

	for i := range len(s) {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()

	for index, value := range s {
		fmt.Printf("%c at index %d\n", value, index)
	}
}
