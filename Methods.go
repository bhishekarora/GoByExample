package main

import "fmt"

type Customer struct {
	Name string
	Age  int
}

func (c *Customer) GetAge() int {
	return c.Age
}
func (c Customer) GetName() string {
	return c.Name
}

func main() {
	c := Customer{Name: "Jalpesh", Age: 35}
	fmt.Println(c)
	fmt.Println(c.GetAge())
	fmt.Println(c.GetName())
}
