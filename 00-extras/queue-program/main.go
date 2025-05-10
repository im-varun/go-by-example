package main

import "fmt"

type node struct {
	data int
	next *node
}

type Queue struct {
	head *node
	tail *node
	size int
}

func (queue *Queue) Enqueue(data int) {
	newNode := &node{data: data, next: nil}

	if queue.IsEmpty() {
		queue.head = newNode
		queue.tail = newNode
		queue.size++
		return
	}

	queue.tail.next = newNode
	queue.tail = newNode
	queue.size++
}

func (queue *Queue) Dequeue() int {
	if queue.IsEmpty() {
		return -1
	}

	dequeuedValue := queue.head.data

	queue.head = queue.head.next
	if queue.head == nil {
		queue.tail = nil
	}
	queue.size--

	return dequeuedValue
}

func (queue *Queue) Peek() int {
	if queue.IsEmpty() {
		return -1
	}

	return queue.head.data
}

func (queue *Queue) IsEmpty() bool {
	return queue.head == nil && queue.tail == nil
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) Print() {
	currentNode := queue.head

	fmt.Print("[ ")
	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println("]")
}

func main() {
	queue := Queue{}

	queue.Print()
	fmt.Println("size:", queue.Size())
	fmt.Println("queue is empty:", queue.IsEmpty())

	queue.Enqueue(40)
	queue.Enqueue(30)
	queue.Enqueue(20)
	queue.Enqueue(10)

	queue.Print()
	fmt.Println("size:", queue.Size())
	fmt.Println("queue is empty:", queue.IsEmpty())

	fmt.Println("peek top value:", queue.Peek())

	fmt.Println("Dequeue top value:", queue.Dequeue())
	queue.Print()
	fmt.Println("size:", queue.Size())

	v1 := queue.Dequeue()
	v2 := queue.Dequeue()
	fmt.Println("v1:", v1, "v2:", v2)

	queue.Print()
	fmt.Println("size:", queue.Size())

	queue.Dequeue()

	queue.Print()
	fmt.Println("size:", queue.Size())
	fmt.Println("queue is empty:", queue.IsEmpty())

	fmt.Println("dequeue on empty queue (-1 means queue is empty):", queue.Dequeue())
	fmt.Println("peek on empty queue (-1 means queue is empty):", queue.Peek())
}
