package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("two random integers in range [0, 100):", rand.IntN(100), rand.IntN(100))
	fmt.Println("two random integers in range [0,10):", rand.IntN(10), rand.IntN(10))

	fmt.Println("random float in range [0, 1):", rand.Float64())
	fmt.Println("random float in range [5, 10)", (rand.Float64()*5)+5)

	s1 := rand.NewPCG(42, 1024)
	r1 := rand.New(s1)
	fmt.Println(r1.IntN(100), r1.IntN(100))

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Println(r2.IntN(100), r2.IntN(100))
}
