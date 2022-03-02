package main

import (
	"fmt"
	"math"
)

const size = 10

func findTwoSmallest(input [size]int) (int, int) {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	secondMin := math.MaxInt
	for i := 0; i < len(input); i++ {
		if input[i] < secondMin && input[i] > min {
			secondMin = input[i]
		}
	}
	return min, secondMin
}

func twoSmallestNew(input [size]int) (int, int) {
	min := math.MaxInt
	secondMin := math.MaxInt

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			secondMin = min
			min = input[i]
		} else if input[i] < secondMin {
			secondMin = input[i]
		}
	}
	return min, secondMin
}

func findTwoLargestNumbers(input [size]int) (int, int) {
	max := math.MinInt
	secondMax := math.MinInt
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			secondMax = max
			max = input[i]
		} else if input[i] > secondMax {
			secondMax = input[i]
		}
	}
	return max, secondMax
}
func main() {
	number := [size]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(findTwoSmallest(number))
	fmt.Println(twoSmallestNew(number))
	fmt.Println(findTwoLargestNumbers(number))
}
