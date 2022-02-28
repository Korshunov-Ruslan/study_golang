package main

import (
	"fmt"
	"math"
)

//func main() {
//	fmt.Println("Разложение ex в ряд Тейлора\n")
//	var x, st float64
//	fmt.Scan(&x, &st)
//	epsilon := 1 / math.Pow(10, st)
//	fmt.Println(epsilon)
//	factorial := 1
//	result := 0.0
//	prevResult := 0.0
//	n := 0
//	for {
//		if n > 0 {
//			factorial *= n
//		}
//		result += math.Pow(x, float64(n)) / float64(factorial)
//		if math.Abs(result-prevResult) < epsilon {
//			break
//		}
//		n++
//		prevResult = result
//	}
//
//}

func main() {
	fmt.Println("Проблемы округления")
	var sum, monthPerc float64
	va
	var year int
	fmt.Scan(&sum, &monthPerc, &year)
	var roundingError float64
	for i := 0; i < 12; i++ {
		newSum := sum + sum*monthPerc/100
		sum = math.Trunc(newSum*100) / 100
		roundingError += newSum - sum
	}
	fmt.Println(sum, roundingError)
	//fmt.Println(a, "\n", a*100, "\n", math.Trunc(a*100)/100)
}
