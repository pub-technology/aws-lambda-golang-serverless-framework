package main

import "fmt"

func doubleMe(x float64) float64 {
	return x * 2
}

func main() {
	fmt.Println("Learn unit test in Go!")
	fmt.Println(doubleMe(21))
}
