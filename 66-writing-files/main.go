package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	d1 := []byte("hello, world!")
	err := os.WriteFile("data1.txt", d1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("data2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)

	n4, err := w.WriteString("buffered\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
