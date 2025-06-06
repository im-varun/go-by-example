package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a string, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people1 := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people1, func(a Person, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println("sorting people by age in ascending order:", people1)

	people2 := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people2, func(a Person, b Person) int {
		return cmp.Compare(b.age, a.age)
	})
	fmt.Println("sorting people by age in descending order:", people2)
}
