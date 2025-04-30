package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty array a:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("length of array a:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array b:", b)

	c := [...]int{1, 2, 3, 4}
	fmt.Println("array c:", c)
	fmt.Println("length of array c:", len(c))

	d := [...]int{100, 3: 400, 500}
	fmt.Println("array d:", d)

	var twoD1 [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			twoD1[i][j] = i + j
		}
	}
	fmt.Println("two dimensional array1:", twoD1)

	var twoD2 [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("two dimensional array2:", twoD2)

	twoD3 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("two dimensional array3:", twoD3)
}
