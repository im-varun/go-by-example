package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	individualArg := os.Args[2]
	fmt.Println(individualArg)
}
