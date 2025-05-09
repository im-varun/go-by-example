package main

import (
	"fmt"
	"iter"
	"slices"
)

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

func (list *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := list.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func main() {
	list := List[int]{}

	list.Push(10)
	list.Push(20)
	list.Push(30)
	list.Push(40)

	for e := range list.All() {
		fmt.Println(e)
	}

	all := slices.Collect(list.All())
	fmt.Println("all:", all)
}
