package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perimeter())
}

func detectRectangle(g geometry) {
	if r, ok := g.(rectangle); ok {
		fmt.Println("rectangle with width:", r.width)
		fmt.Println("rectangle with height:", r.height)
	} else {
		fmt.Println("error: not a rectangle")
	}
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius:", c.radius)
	} else {
		fmt.Println("error: not a circle")
	}
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectRectangle(r)
	detectRectangle(c)

	detectCircle(r)
	detectCircle(c)
}
