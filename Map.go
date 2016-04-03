package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 5
	m["k2"] = 10

	fmt.Println("Map m", m)

	v1 := m["k2"]
	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("m after delete", m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
