package main

import "os"

func panicTest() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

func main() {
	// panic("a problem")
	panicTest()
}
