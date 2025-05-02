package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	seq1 := intSeq()
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())

	seq2 := intSeq()
	fmt.Println(seq2())

	seq3 := intSeq()
	fmt.Println(seq3())
	fmt.Println(seq3())
}
