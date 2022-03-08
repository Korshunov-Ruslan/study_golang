package main

import "fmt"

func acceptFunc(a, b, c int, F func(int, int) int) int {
	return a + F(b, c)
}
func main() {
	fmt.Println(acceptFunc(10, 20, 30, func(b int, c int) int {
		return b + c
	}))
}
