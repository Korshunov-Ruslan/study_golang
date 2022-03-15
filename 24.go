package main

import "fmt"

//func main() {
//	arr := [10]int{99, 2, 1, 3, 4, 2, 2, 100, 1, 2}
//	fmt.Println(InsertionSort(arr))
//}
//func InsertionSort(ar [10]int) [10]int {
//	for i := 0; i < len(ar); i++ {
//		x := ar[i]
//		j := 1
//		for j >= 1 && ar[j-1] > x {
//			j--
//			fmt.Println(ar[i])
//			ar[j] = ar[j-1]
//		}
//		ar[j] = x
//	}
//	return ar
//}

func main() {
	array := [10]int{1, -5, 2, 1, 5, 1000, 7, 2, 100, 1}
	anon := func(arr [10]int) [10]int {
		for i := 10; i > 0; i-- {
			for j := 1; j < i; j++ {
				if arr[j-1] > arr[j] {
					temp := arr[j]
					arr[j] = arr[j-1]
					arr[j-1] = temp
				}
			}
		}
		return arr
	}
	fmt.Println(anon(array))
}
