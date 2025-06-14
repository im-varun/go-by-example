package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if e != nil {
		panic(e)
	}
	fmt.Println(t1)

	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	if e != nil {
		panic(e)
	}
	fmt.Println(t2)

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)
}
