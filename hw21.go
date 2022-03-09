package main

import "fmt"

func main() {
	/*
		S = 2 × x + y ^ 2 − 3/z;  x = int16, y = uint8 ; z = float32.
	*/
	a := func(x int16, y uint8, z float32) float32 {
		return 2*float32(int64(x)) + float32(int64(y)*int64(y)) - 3/float32(z)
	}
	fmt.Println(a(10, 17, 2))
}

//func main() {
//	/*
//		Write a function that takes a function of the form A func (int, int) int as input,
//		and wraps it inside and calls it on exit.
//	*/
//	f := func(a int, b int) int { return a * b }
//	ff := func(a int, b int) int { return a + b }
//	fff := func(a int, b int) int { return a - b }
//	iAcceptfunc(1, 2, f)
//	iAcceptfunc(1, 2, ff)
//	iAcceptfunc(1, 2, fff)
//}
//func iAcceptfunc(a, b int, f func(int, int) int) {
//	defer fmt.Println(f(1, 2))
//
//}
