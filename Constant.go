package main

import "fmt"
import "math"

const s = "Hello world"

func main() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	const m int = 10
	fmt.Println(m)
}
