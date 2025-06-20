package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("path:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	dir, base := filepath.Split(p)
	fmt.Println("dir:", dir)
	fmt.Println("base:", base)

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println("file extension:", ext)
	fmt.Println("file name:", strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
