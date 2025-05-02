package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for index, num := range nums {
		fmt.Println("index:", index, "num:", num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cat"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range kvs {
		fmt.Println("key:", key)
	}

	for _, value := range kvs {
		fmt.Println("value:", value)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
