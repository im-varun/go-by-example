package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	fmt.Println("exiting without defer statement execution")
	os.Exit(1)
}
