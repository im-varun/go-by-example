package main

import "fmt"

const s string = "constant"

func main() {
	fmt.Println(s)

	const n1 int = 4
	fmt.Println(n1)

	const n2 = 8
	fmt.Println(n2)
}
