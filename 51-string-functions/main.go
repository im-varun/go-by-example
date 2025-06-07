package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("contains:", strings.Contains("test", "es"))
	fmt.Println("count:", strings.Count("test", "t"))
	fmt.Println("has prefix:", strings.HasPrefix("test", "te"))
	fmt.Println("has suffix:", strings.HasSuffix("test", "st"))
	fmt.Println("index:", strings.Index("test", "t"))
	fmt.Println("join:", strings.Join([]string{"a", "b", "c", "d"}, "-"))
	fmt.Println("repeat:", strings.Repeat("v", 4))
	fmt.Println("replace:", strings.Replace("foo", "o", "0", -1))
	fmt.Println("replace:", strings.Replace("foo", "o", "0", 1))
	fmt.Println("split:", strings.Split("a-b-c-d", "-"))
	fmt.Println("to lower:", strings.ToLower("TEST"))
	fmt.Println("to upper:", strings.ToUpper("test"))
}
