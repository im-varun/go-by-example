package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func main() {
	fmt.Println(person{"Alice", 20})

	fmt.Println(person{name: "John", age: 24})
	fmt.Println(person{age: 30, name: "Trent"})

	fmt.Println(person{name: "Fred"})
	fmt.Println(person{age: 18})

	fmt.Println(&person{name: "Ann", age: 42})

	fmt.Println(newPerson("David", 48))

	s := person{name: "Sean", age: 36}
	fmt.Println("name:", s.name)
	fmt.Println("age:", s.age)

	sp := &s
	fmt.Println("name:", sp.name)
	fmt.Println("age:", sp.age)

	sp.age = 32
	fmt.Println("name:", sp.name)
	fmt.Println("age:", sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
