package main

import "fmt"

type Customer struct {
	Name string
	Age  int
}

func (p *Customer) GetAge() int {
	return p.Age
}
func (p Customer) GetName() string {
	return p.Name
}

func main() {
	c := Customer{Name: "Jalpesh", Age: 35}
	fmt.Println(c)
	fmt.Println(c.GetAge())
	fmt.Println(c.GetName())
}
