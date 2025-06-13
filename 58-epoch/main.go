package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println("elapsed time since UNIX epoch in seconds:", now.Unix())
	fmt.Println("elapsed time since UNIX epoch in milliseconds:", now.UnixMilli())
	fmt.Println("elapsed time since UNIX epoch in nanoseconds:", now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
