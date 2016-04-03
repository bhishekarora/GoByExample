package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func anotherPrint() {
	fmt.Println("Void kind of print function")
}

func main() {
	res := plus(5, 10)
	fmt.Println(res)
	anotherPrint()
}
