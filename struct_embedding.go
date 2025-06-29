package main

import "fmt"

type base struct {
	name string
}

func (b base) describe() string {
	return fmt.Sprintf("My name is: %s", b.name)
}

type container struct {
	base
	age int
}

func main() {
	co := container{
		base: base{
			name: "hardik",
		},
		age: 21,
	}

	fmt.Println(co.describe())
	fmt.Println(co.name)
	fmt.Println(co.age)

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println(d.describe())
}
