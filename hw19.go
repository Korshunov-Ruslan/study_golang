package main

import (
	"fmt"
	"sort"
)

func getNewArray(firArr []int, secArr []int) []int {
	sort.Ints(firArr)
	sort.Ints(secArr)
	lenTwoArr := len(firArr) + len(secArr)
	newestArr := make([]int, 0, lenTwoArr)
	n, m := len(firArr), len(secArr)
	i, j := 0, 0
	newestArr = append(newestArr)
	for i < n && j < m {
		if firArr[i] <= secArr[j] {
			newestArr = append(newestArr, firArr[i])
			i++
		} else {
			newestArr = append(newestArr, secArr[j])
			j++
		}
	}
	if len(firArr) < len(secArr) {
		newestArr = append(newestArr, secArr[j:]...)
	} else if len(firArr) > len(secArr) {
		newestArr = append(newestArr, secArr[i:]...)
	} else {
		newestArr = append(newestArr, secArr[i:]...)
		newestArr = append(newestArr, secArr[j:]...)
	}
	return newestArr
}
func main() {
	/*
		Function that merges 2 sorted arrays.
	*/
	firstArray := []int{2, 12, 33, 33}
	secondArray := []int{1, 91, 2, 99}
	fmt.Println(getNewArray(firstArray, secondArray))
}

//
//func bubleSort(arr [6]int) [6]int {
//	for i := 0; i < len(arr)-1; i++ {
//		for k := 0; k < len(arr)-i-1; k++ {
//			if arr[k] > arr[k+1] {
//				arr[k], arr[k+1] = arr[k+1], arr[k]
//			}
//		}
//	}
//	return arr
//}
//func main() {
//	/*
//		Bubble sort
//	*/
//	arr := [6]int{1000, 5, 500, 100, 94, 999}
//	fmt.Println(bubleSort(arr))
//}
