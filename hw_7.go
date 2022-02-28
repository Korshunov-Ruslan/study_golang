package main

//func main() {
//	fmt.Println("Задание 1. Зеркальные билеты\n")
//	var min, max int = 100000, 999999
//	var count int
//	for i := min; i <= max; i++ {
//		if i/100000 == (i%100)%10 && (i/10000)%10 == (i%100)/10 && (i/1000)%10 == (i%1000)/100 {
//			count++
//		}
//	}
//	fmt.Printf("Количество зеркальных билетов среди всех шестизначных номеров от %v до %v:%v", min, max, count)
//}

//func main() {
//	fmt.Println("Задание 2. Шахматная доска")
//	var width, height int
//	fmt.Println("Введите ширину")
//	fmt.Scan(&width)
//	fmt.Println("Введите высоту")
//	fmt.Scan(&height)
//	for i := 0; i < height; i++ {
//		if i%2 == 0 {
//			for j := 0; j < width; j++ {
//				if j%2 == 0 {
//					fmt.Print(" ")
//				} else {
//					fmt.Print("*")
//				}
//			}
//			fmt.Println("")
//		} else {
//			for k := 0; k < width; k++ {
//				if k%2 == 0 {
//					fmt.Print("*")
//				} else {
//					fmt.Print(" ")
//				}
//			}
//			fmt.Println("")
//		}
//	}
//}
//func main() {
//	fmt.Println("Задание 3. Вывод ёлочки\n")
//	var height int
//	fmt.Println("Введите высоту ёлочки")
//	fmt.Scan(&height)
//	for i := 0; i < height; i++ {
//		for j := i + 1; j < height; j++ {
//			fmt.Print(" ")
//		}
//
//		for k := 0; k <= i*2; k++ {
//			fmt.Print("*")
//		}
//		fmt.Println()
//	}
//}
//func main() {
//	fmt.Println("Задание 4 (по желанию). Счастливые билеты\n")
//	var min, max int = 100000, 999999
//	var buyTicketsCount int
//	var currentNumber int
//	for i := min; i <= max; i++ {
//		if (i/100000 + (i/10000)%10 + (i/1000)%10) == ((i%100)%10 + (i%100)/10 + (i%1000)/100) {
//			buyTicketsCount = i - currentNumber
//			currentNumber = i
//		}
//	}
//	fmt.Printf("Минимальное количество билетов, которое нужно купить, "+
//		"чтобы среди них оказался счастливый: %v", buyTicketsCount)
//}
