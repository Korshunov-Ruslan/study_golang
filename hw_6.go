package main

//func main() {
//	fmt.Println("Задание 1. Написание простого цикла\n")
//	var userInput int
//	fmt.Println("Enter any int number")
//	fmt.Scan(&userInput)
//	for i := 0; i != userInput; {
//		if userInput < 0 {
//			fmt.Println(i)
//			i -= 1
//		} else {
//			fmt.Println(i)
//			i += 1
//		}
//	}
//}
//func main() {
//	fmt.Println("Задание 2. Сумма двух чисел по единице\n")
//	var firstInput, secondInput int
//	fmt.Println("Enter 2 int numbers")
//	fmt.Scan(&firstInput, &secondInput)
//	var sum int = firstInput + secondInput
//	for firstInput != sum {
//		firstInput += 1
//		fmt.Println(firstInput)
//	}
//	fmt.Println("Сумма двух чисел достигнута!")
//}
//func main() {
//	fmt.Println("Задание 3. Расчёт суммы скидки\n")
//	var productPrice, discount int
//	var maxDiscountRub, maxDiscountPercnt int = 2000, 30
//	fmt.Println("Введите цену товара и скидку")
//	fmt.Scan(&productPrice, &discount)
//	var currentPercent = maxDiscountRub / 30 * 100
//	if discount > maxDiscountRub || currentPercent > maxDiscountPercnt {
//		fmt.Println("Вы запросили слишком большую скидку! Максимальная скидка:", maxDiscountRub, "руб и не более",
//			maxDiscountPercnt, "%")
//		if discount > maxDiscountRub {
//			fmt.Println("Максимальная возможная скидка:", maxDiscountRub, "рублей.")
//		} else if currentPercent > maxDiscountPercnt {
//			fmt.Println("В данном случае ваше скидка будет:", maxDiscountPercnt, "% или", productPrice/100*30, "рублей")
//		} else {
//			fmt.Println("Такая скидка возможна!")
//		}
//	}
//}
//func main() {
//	fmt.Println("Задание 4. Предыдущее занятие на if")
//	var a, b, c int = 0, 0, 0
//	for i := 0; i < 1111; i++ {
//		for a != 10 {
//			a++
//			if a == 10 {
//				break
//			}
//		}
//		for b != 100 {
//			b++
//			if b == 100 {
//				break
//			}
//		}
//		for c != 1000 {
//			c++
//			break
//		}
//	}
//	fmt.Println(a, b, c)
//}
//func main() {
//	fmt.Println("Задание 5 (по желанию). Задача на постепенное наполнение корзин разной ёмкости")
//	var firstValue, secondValue, thirdValue int
//	fmt.Println("Введите ёмкости для трёх корзин")
//	fmt.Scan(&firstValue, &secondValue, &thirdValue)
//	var countForFirst, countForSecond, countForThird int = 0, 0, 0
//	var sumValue int = firstValue + secondValue + thirdValue
//	for i := 0; i != sumValue; i++ {
//		if countForFirst != firstValue {
//			fmt.Println("Заполняем первую корзину")
//			countForFirst++
//			continue
//		}
//		if countForSecond != secondValue {
//			fmt.Println("Заполняем вторую корзину")
//			countForSecond++
//			continue
//		}
//		if countForThird != thirdValue {
//			fmt.Println("Заполняем третью корзину")
//			countForThird++
//			continue
//		}
//		if countForFirst+countForSecond+countForThird == sumValue {
//			break
//		}
//	}
//	fmt.Println(countForFirst, countForSecond, countForThird)
//}
