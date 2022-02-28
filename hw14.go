package main

import "fmt"

///*
//First ex
//*/
//func isEven(a int) bool {
//	if a%2 == 0 {
//		return true
//	} else {
//		return false
//	}
//}
//func main() {
//	fmt.Println(isEven(1))
//}
/*
Second ex
*/

//func getCoord(a int) (int, int) {
//	x := rand.Intn(a)
//	y := rand.Intn(a)
//	return x, y
//}
//func changeCoord(x, y int) (int, int) {
//	newX := x*2 + 10
//	newY := -3*y - 5
//	return newX, newY
//}
//func main() {
//	a := 25
//	for i := 0; i < 3; i++ {
//		fmt.Println(changeCoord(getCoord(a)))
//	}
//}

/*
third ex
*/

//func multiplication(a int) (b int) {
//	b = a * 3
//	return
//}
//func sum(b int) (c int) {
//	c = b + 56
//	return
//}
//func main() {
//	var a int = 101
//	fmt.Println("Start number:", a)
//	result := sum(multiplication(a))
//	fmt.Printf("Nubmer after all operations: %v", result)
//}

/*
fourth ex
*/

const val = 50
const second = 200
const third = 300
const first = 100

func f1(a int) (res int) {
	res = a + first
	return
}

func f2(a int) (res int) {
	res = a + second
	return
}

func f3(val int) (res int) {
	res = third + val
	return
}

func main() {
	fmt.Println(f1(1))
	f1(1)
	fmt.Println(f2(val))
	fmt.Println(f3(111))
	fmt.Println(f3(f2(f1(val))))
}
