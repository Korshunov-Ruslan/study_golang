package main

import "fmt"

func main() {
	const el = 10
	array := [el]int{10, 2, 4, 4, 2, 1, 5, 6, 2, 10}
	fmt.Println("-----UNSORTED-----")
	fmt.Println(array)
	fmt.Println("-----SORTED-----")
	fmt.Println(sort(array))
}
func sort(arr [10]int) [10]int {
	for i := 0; i < 10; i++ {
		minIdx := i
		for j := i; j < 10; j++ {
			if arr[j] <= arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}
