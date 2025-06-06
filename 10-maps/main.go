package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 5
	m["k2"] = 7

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n1)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n1, n2) {
		fmt.Println("n1 == n2")
	}
}
