package main

import "fmt"

func main() {
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 4; i++ {
		fmt.Println(i)
	}

	for i := range 4 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}
}
