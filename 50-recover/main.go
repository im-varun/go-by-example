package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from error:", r)
		}
	}()

	defer func() {
		fmt.Println("another function in deferred functions stack")
	}()

	mayPanic()

	fmt.Println("after mayPanic()")
}
