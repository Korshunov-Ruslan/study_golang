package main

import (
	"fmt"
)

//import "fmt"
//
//func getArgument(a, b int) {
//	var count int
//	if a > b {
//		for i := b; i < a+1; i++ {
//			if i%2 == 0 {
//				count++
//			}
//		}
//	} else {
//		for i := a; i < b+1; i++ {
//			if i%2 == 0 {
//				count++
//			}
//		}
//	}
//	fmt.Println(count)
//}
//func main() {
//	getArgument(1, 12)
//}

func invertNumbers(a, b *int) {
	c := *a
	*a = *b
	*b = c
	fmt.Println(*a, *b)
}

func main() {
	first := 1
	var abc int = 2
	fmt.Println(first, abc)
	invertNumbers(&first, &abc)
}
