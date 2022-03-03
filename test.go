package main

import "fmt"

func maximun(arr [3]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func matrixMax(matrix [4][3]int) [4]int {
	var result [4]int
	for i := 0; i < len(matrix); i++ {
		min := maximun(matrix[i])
		fmt.Println(min)
		result[i] = min
	}
	return result
}
func diagonal(matrix [4][3]int) int {
	a := 0
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i][i])
	}
	return a
}
func main() {
	a := [4][3]int{
		{3, 1, 3},
		{2, 2, 2},
		{1, 1, 0},
		{100, 100, 100},
	}
	fmt.Println(diagonal(a))
}
