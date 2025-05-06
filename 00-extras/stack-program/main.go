package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (stack *Stack) Push(data int) {
	newNode := &Node{data: data, next: nil}

	if stack.IsEmpty() {
		stack.head = newNode
		stack.size++
		return
	}

	newNode.next = stack.head
	stack.head = newNode
	stack.size++
}

func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		return -1
	}

	poppedValue := stack.head.data

	stack.head = stack.head.next
	stack.size--

	return poppedValue
}

func (stack *Stack) Peek() int {
	if stack.IsEmpty() {
		return -1
	}

	return stack.head.data
}

func (stack *Stack) IsEmpty() bool {
	return stack.head == nil
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Print() {
	currentNode := stack.head

	fmt.Print("[ ")
	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println("]")
}

func main() {
	stack := Stack{}

	stack.Print()
	fmt.Println("size:", stack.Size())
	fmt.Println("stack is empty:", stack.IsEmpty())

	stack.Push(40)
	stack.Push(30)
	stack.Push(20)
	stack.Push(10)

	stack.Print()
	fmt.Println("size:", stack.Size())
	fmt.Println("stack is empty:", stack.IsEmpty())

	fmt.Println("peek top value:", stack.Peek())

	fmt.Println("pop top value:", stack.Pop())
	stack.Print()
	fmt.Println("size:", stack.Size())

	v1 := stack.Pop()
	v2 := stack.Pop()
	fmt.Println("v1:", v1, "v2:", v2)

	stack.Print()
	fmt.Println("size:", stack.Size())

	stack.Pop()

	stack.Print()
	fmt.Println("size:", stack.Size())
	fmt.Println("stack is empty:", stack.IsEmpty())

	fmt.Println("pop on empty stack (-1 means stack is empty):", stack.Pop())
	fmt.Println("peek on empty stack (-1 means stack is empty):", stack.Peek())
}
