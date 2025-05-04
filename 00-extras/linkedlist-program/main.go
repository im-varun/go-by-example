package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (linkedList *LinkedList) Insert(data int) {
	newNode := &Node{data: data, next: nil}

	if linkedList.head == nil {
		linkedList.head = newNode
		return
	}

	newNode.next = linkedList.head
	linkedList.head = newNode
}

func (linkedList *LinkedList) Contains(data int) bool {
	isFound := false

	currentNode := linkedList.head

	for currentNode != nil && !isFound {
		if currentNode.data == data {
			isFound = true
		}

		currentNode = currentNode.next
	}

	return isFound
}

func (linkedList *LinkedList) Delete(data int) {
	if linkedList.head == nil {
		return
	}

	var previousNode *Node
	currentNode := linkedList.head

	for currentNode != nil && currentNode.data != data {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	if previousNode == nil {
		linkedList.head = currentNode.next
	} else {
		previousNode.next = currentNode.next
	}
}

func (linkedList *LinkedList) Size() int {
	size := 0

	currentNode := linkedList.head

	for currentNode != nil {
		size += 1
		currentNode = currentNode.next
	}

	return size
}

func (linkedList *LinkedList) Print() {
	currentNode := linkedList.head

	fmt.Print("[ ")
	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println("]")

}

func main() {
	linkedList := LinkedList{}

	linkedList.Print()
	fmt.Println("size:", linkedList.Size())

	linkedList.Insert(40)
	linkedList.Insert(30)
	linkedList.Insert(20)
	linkedList.Insert(10)

	linkedList.Print()
	fmt.Println("size:", linkedList.Size())

	fmt.Println("linked list contains 20:", linkedList.Contains(20))
	fmt.Println("linked list contains 50:", linkedList.Contains(50))

	linkedList.Delete(30)
	linkedList.Print()
	fmt.Println("size:", linkedList.Size())
	fmt.Println("linked list contains 30:", linkedList.Contains(30))

	linkedList.Delete(10)
	linkedList.Print()
	fmt.Println("size:", linkedList.Size())
	fmt.Println("linked list contains 10:", linkedList.Contains(10))

	linkedList.Delete(40)
	linkedList.Print()
	fmt.Println("size:", linkedList.Size())
	fmt.Println("linked list contains 40:", linkedList.Contains(40))

	linkedList.Delete(20)
	linkedList.Print()
	fmt.Println("size:", linkedList.Size())
	fmt.Println("linked list contains 20:", linkedList.Contains(20))
}
