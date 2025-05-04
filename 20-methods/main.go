package main

import "fmt"

type rectange struct {
	width  int
	height int
}

func (r *rectange) area() int {
	return r.width * r.height
}

func (r rectange) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rectange{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
