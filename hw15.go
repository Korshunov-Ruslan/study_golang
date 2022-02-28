package main

import "fmt"

func makeArray() [10]int {
	var arr [10]int
	for i, _ := range arr {
		fmt.Println("Введите число")
		fmt.Scan(&arr[i])
	}
	return arr
}
func isNumberEven(arr [10]int) {
	var (
		even    int
		notEven int
	)
	for _, i2 := range arr {
		if i2%2 == 0 {
			even++
		} else {
			notEven++
		}
	}
	fmt.Printf("Кол-во чётных чисел:%v, а кол-во нечётных чисел:%v", even, notEven)
}
func main() {
	isNumberEven(makeArray())
}

//func reverseArray(a []int) []int {
//	var newArr []int
//	for i := range a {
//		n := a[len(a)-1-i]
//		newArr = append(newArr, n)
//	}
//	return newArr
//}
//
//func main() {
//	arr := []int{9, 8, 7, 1, 2}
//	fmt.Println("Array at start:", arr)
//	fmt.Println("Array at the end:", reverseArray(arr))
//}
