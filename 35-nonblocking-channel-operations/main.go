package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("message received:", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("message sent:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("message received:", msg)
	case sig := <-signals:
		fmt.Println("signal received:", sig)
	default:
		fmt.Println("no activity")
	}

	requests := make(chan string, 1)
	responses := make(chan string, 1)

	req := "hello, world!"
	select {
	case requests <- req:
		fmt.Println("request sent:", req)
	default:
		fmt.Println("no request sent")
	}

	responses <- "200"
	select {
	case res := <-responses:
		fmt.Println("response received:", res)
	default:
		fmt.Println("no response received")
	}
}
