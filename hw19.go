package main

import "fmt"

//func getNewArray(firArr []int, secArr []int) [9]int {
//	sort.Ints(firArr)
//	sort.Ints(secArr)
//	var newArr [9]int
//	copy(newArr[:], firArr[:])
//	copy(newArr[4:], secArr[:])
//	return newArr
//}
//func main() {
//	/*
//		Function that merges 2 sorted arrays.
//	*/
//	firstArray := []int{99, 12, 33, 1}
//	secondArray := []int{1, 4, 2, 3, 100}
//	fmt.Println(getNewArray(firstArray, secondArray))
//}

func bubleSort(arr [6]int) [6]int {
	for i := 0; i < len(arr)-1; i++ {
		for k := 0; k < len(arr)-i-1; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
	}
	return arr
}
func main() {
	/*
		Bubble sort
	*/
	arr := [6]int{1000, 5, 500, 100, 94, 999}
	fmt.Println(bubleSort(arr))
}
