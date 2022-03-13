package main

import (
	"fmt"
	"strings"
)

const elem = 10

//func main() {
//	/*
//		Even and Odd
//	*/
//	array := [elem]int{1, 55, 22, 44, 11, 2, 0, 3, 2, 1}
//	even, odd := evenAndOdd(array)
//	fmt.Printf("Even numbers:%v and odd numbers:%v", even, odd)
//}
//
//func evenAndOdd(arr [elem]int) ([]int, []int) {
//	var evens = []int{}
//	var odds = []int{}
//	for _, i2 := range arr {
//		if i2%2 == 0 {
//			evens = append(evens, i2)
//		} else {
//			odds = append(odds, i2)
//		}
//	}
//	return evens, odds
//}
func main() {
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}
	parseTest(sentences, chars)
}
func parseTest(sentences [4]string, chars [5]rune) {
	for i := 0; i < len(sentences); i++ {
		sentences[i] = strings.ToUpper(sentences[i])
		for j := 0; j < len(chars); j++ { //fmt.Println(string(sentences[i]), string(chars[j]))
			if strings.Contains(string(sentences[i]), string(chars[j])) {
				index := strings.Index(sentences[i], string(chars[j]))
				fmt.Printf("Subarray %v have str (%v) at pos: %v\n",
					sentences[i], string(chars[j]), index)
			}
		}
	}
}
