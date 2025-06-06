package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strs:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("ints:", ints)

	sorted := slices.IsSorted(ints)
	fmt.Println("sorted:", sorted)
}
