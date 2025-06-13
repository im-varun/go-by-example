package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())

	fmt.Println(now.Weekday())
	fmt.Println(then.Weekday())

	fmt.Println("then is before now:", then.Before(now))
	fmt.Println("then is after now:", then.After(now))
	fmt.Println("then is at the same time as now:", then.Equal(now))

	diff := now.Sub(then)
	fmt.Println(diff)
	fmt.Println("duration in hours:", diff.Hours())
	fmt.Println("duration in minutes:", diff.Minutes())
	fmt.Println("duration in seconds:", diff.Seconds())

	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
