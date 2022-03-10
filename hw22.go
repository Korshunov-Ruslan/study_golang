package main

import "fmt"

//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
const elem = 12

//
//func main() {
//	/*
//		Populate an array with unordered numbers based on a random number generator. Enter the number.
//		The program must find this number in the array and output how many numbers are in the array after the entered one. If there is no entered number in the array, output 0.
//		For the convenience of checking, implement displaying the array on the screen.
//	*/
//	var value int
//	fmt.Println("Enter a number what you want to find")
//	fmt.Scan(&value)
//	_, number := (findNumberAndIndex(arrayGenerator(), value)
//  fmt.Println(number)
//}
//
//func arrayGenerator() [elem]int {
//	rand.Seed(time.Now().UnixNano())
//	var array [elem]int
//	for i := 0; i < elem; i++ {
//		array[i] = rand.Intn(elem * 12)
//	}
//	fmt.Println(array)
//	return array
//}
func findNumberAndIndex(arr [elem]int, value int) (index int, result int) {
	for i := 0; i < elem; i++ {
		if arr[i] == value {
			result = len(arr[i:]) - 1
			index = i
		}
	}
	return
}

func main() {
	/*
		Fill in an ordered array of 12 elements and enter a number.
		It is necessary to implement the search for the first occurrence of a given number in an array.
		The complexity of the algorithm should be minimal.
	*/
	array := [elem]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	neededValue := 2
	index, _ := findNumberAndIndex(array, neededValue)
	fmt.Println(index)
}
