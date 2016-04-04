package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Jalpesh", 35})
	fmt.Println(person{"Vishal", 32})
	fmt.Println(person{name: "Teerth"})
	fmt.Println(&person{name: "Reena", age: 33})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
