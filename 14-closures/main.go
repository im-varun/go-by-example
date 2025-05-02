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

	var sum func(nums ...int) int
	sum = func(nums ...int) int {
		total := 0

		for _, num := range nums {
			total += num
		}

		return total
	}

	fmt.Println("1 + 2 =", sum(1, 2))
	fmt.Println("1 + 2 + 3 =", sum(1, 2, 3))
	fmt.Println("1 + 2 + 3 + 4 =", sum(1, 2, 3, 4))
}
