package main

import (
	"fmt"
	"math/rand"
	"time"
)

const el = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [el]int
	for i := 0; i < 10; i++ {
		arr[i] = 10*i + rand.Intn(10)
	}
	fmt.Println(arr)
	value := arr[rand.Intn(el)]
	fmt.Println(value)
	fmt.Println(binarySearch(arr, value))
	value = rand.Intn(500)
	fmt.Println(value)
	fmt.Println(binarySearch(arr, value))
}
func binarySearch(arr [el]int, value int) int {
	index := -1
	min := 0
	max := el - 1
	for max >= min {
		middle := (max + min) / 2
		if arr[middle] == value {
			index = middle
			break
		} else if arr[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return index
}
