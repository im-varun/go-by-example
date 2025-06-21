package main

import (
	"embed"
	"fmt"
)

//go:embed single_file.txt
var fileString string

//go:embed single_file.txt
var fileByte []byte

//go:embed single_file.txt
//go:embed *.hash
var folder embed.FS

func main() {
	fmt.Println(fileString)
	fmt.Println(string(fileByte))

	content1, _ := folder.ReadFile("file1.hash")
	fmt.Println(string(content1))

	content2, _ := folder.ReadFile("file2.hash")
	fmt.Println(string(content2))
}
