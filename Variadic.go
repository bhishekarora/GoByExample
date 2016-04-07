package main

import "fmt"

func sum(nubs ...int) {

	fmt.Print(nubs, " ")
	total := 0
	for _, num := range nubs {
		total += num
	}
	fmt.Println(total)

}
func main() {

	sum(1, 2)
	sum(1, 2, 3)
	test := []int{1, 2, 3, 4}
	sum(test...)

}
