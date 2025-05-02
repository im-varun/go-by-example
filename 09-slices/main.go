package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[0])

	fmt.Println("len:", len(s), "cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t1 := []string{"g", "h", "i"}
	fmt.Println("dcl:", t1)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t1, t2) {
		fmt.Println("t1 == t2")
	}

	t3 := []string{"a", "b", "c"}
	if !(slices.Equal(t1, t3)) {
		fmt.Println("t1 != t3")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	nums := []int{10, 20, 30, 40, 50}
	for index, value := range nums {
		fmt.Println("index:", index, "value:", value)
	}
}
