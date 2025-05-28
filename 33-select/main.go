package main

import (
	"fmt"
	"time"
)

func one(c1 chan<- string) {
	time.Sleep(1 * time.Second)
	c1 <- "one"
}

func two(c2 chan<- string) {
	time.Sleep(2 * time.Second)
	c2 <- "two"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go one(c1)
	go two(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
