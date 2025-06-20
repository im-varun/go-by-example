package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)
	defer os.Remove(f.Name())

	fmt.Println("temp file name:", f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dirname, err := os.MkdirTemp("", "sampledir")
	check(err)
	defer os.RemoveAll(dirname)

	fmt.Println("temp dir name:", dirname)

	fname := filepath.Join(dirname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
