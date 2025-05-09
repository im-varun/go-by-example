package main

import "fmt"

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head *element[T]
	tail *element[T]
}

type element[T any] struct {
	val  T
	next *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) AllElements() []T {
	var elements []T

	for e := list.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}

	return elements
}

func main() {
	var s = []string{"food", "bar", "baz"}

	fmt.Println("index of baz:", SliceIndex(s, "baz"))
	fmt.Println("another way for index of baz:", SliceIndex[[]string, string](s, "baz"))

	list := List[int]{}
	list.Push(10)
	list.Push(20)
	list.Push(30)
	list.Push(40)
	fmt.Println("list:", list.AllElements())
}
