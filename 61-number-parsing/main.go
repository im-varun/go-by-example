package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	h, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(h)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	c, _ := strconv.Atoi("135")
	fmt.Println(c)

	_, e := strconv.Atoi("hello")
	fmt.Println(e)
}
