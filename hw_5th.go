package main

import "fmt"

//func main() {
//	fmt.Println("Задание 1. Определение координатной плоскости точки\n")
//	var osX, osY int
//	fmt.Println("Введите значение для оси X")
//	fmt.Scan(&osX)
//	fmt.Println("Введите значение для оси Y")
//	fmt.Scan(&osY)
//	if osX > 0 && osY > 0 {
//		fmt.Println("Это 1 четверть.")
//	} else if osX < 0 && osY > 0 {
//		fmt.Println("Это вторая четверть.")
//	} else if osX < 0 && osY < 0 {
//		fmt.Println("Это третья четверть.")
//	} else if osX > 0 && osY < 0 {
//		fmt.Println("Это четвертая четверть.")
//	} else {
//		fmt.Println("Вы вначале координат")
//	}
//}

//func main() {
//	fmt.Println("Задание 2. Проверить, что одно из чисел — положительное\n")
//	var firstValue, secondValue, thirdValue int
//	var positiveValue bool = true
//	fmt.Println("Введите первое число")
//	fmt.Scan(&firstValue)
//	fmt.Println("Введите второе число")
//	fmt.Scan(&secondValue)
//	fmt.Println("Введите третье число")
//	fmt.Scan(&thirdValue)
//	if firstValue > 0 || secondValue > 0 || thirdValue > 0 {
//		fmt.Println(positiveValue)
//	} else {
//		positiveValue = false
//		fmt.Println(positiveValue)
//	}
//}

//func main() {
//	fmt.Println("Задание 3. Проверить, что есть совпадающие числа\n")
//	var firstValue, secondValue, thirdValue int
//	var isSame bool = true
//	fmt.Println("Введите первое число")
//	fmt.Scan(&firstValue)
//	fmt.Println("Введите второе число")
//	fmt.Scan(&secondValue)
//	fmt.Println("Введите третье число")
//	fmt.Scan(&thirdValue)
//	if firstValue == secondValue || secondValue == thirdValue {
//		fmt.Println(isSame)
//	} else if firstValue == thirdValue {
//		fmt.Println(isSame)
//	} else {
//		isSame = false
//		fmt.Println(isSame)
//	}
//}
//func main() {
//	fmt.Println("Задание 4. Сумма без сдачи\n")
//	var productPrice, firstCoin, secondCoin, thirdCoin float64
//	fmt.Println("Введите стоимость товара")
//	fmt.Scan(&productPrice)
//	fmt.Println("Введите номинал 1 монеты")
//	fmt.Scan(&firstCoin)
//	fmt.Println("Введите номинал 2 монеты")
//	fmt.Scan(&secondCoin)
//	fmt.Println("Введите номинал 3 монеты")
//	fmt.Scan(&thirdCoin)
//	var isOk bool = true
//	if productPrice/firstCoin == 1 || productPrice/(secondCoin) == 1 {
//		fmt.Println(isOk)
//	} else if productPrice/thirdCoin == 1 {
//		fmt.Println(isOk)
//	} else if productPrice/(firstCoin+secondCoin) == 1 || productPrice/(firstCoin+thirdCoin) == 1 {
//		fmt.Println(isOk)
//	} else if productPrice/(secondCoin+thirdCoin) == 1 {
//		fmt.Println(isOk)
//	} else if productPrice/(firstCoin+secondCoin+thirdCoin) == 1 {
//		fmt.Println(isOk)
//	} else {
//		isOk = false
//		fmt.Println(isOk)
//	}
//}
//func main() {
//	fmt.Println("Задание 5. Определение максимальных процентов\n")
//	var firstRate, secondRate, thirdRate float64
//	fmt.Println("Введите первую ставку:")
//	fmt.Scan(&firstRate)
//	fmt.Println("Введите вторую ставку:")
//	fmt.Scan(&secondRate)
//	fmt.Println("Введите третью ставку:")
//	fmt.Scan(&thirdRate)
//	if firstRate == secondRate && secondRate == thirdRate {
//		fmt.Println("Ставки равны")
//	} else if firstRate < secondRate && firstRate < thirdRate {
//		fmt.Print("Первая ставка самая выгодная ==> ", firstRate, "%")
//	} else if thirdRate < firstRate && thirdRate < secondRate {
//		fmt.Print("Третья ставка самая выгодная ==> ", thirdRate, "%")
//	} else {
//		fmt.Print("Вторая ставка самая выгодная ==> ", secondRate, "%")
//	}
//}
//func main() {
//	fmt.Println("Задание 6. Решение квадратного уравнения\n")
//	var a, b, c, d, x1, x2 float64
//	fmt.Println("Введите а")
//	fmt.Scan(&a)
//	fmt.Println("Введите б")
//	fmt.Scan(&b)
//	fmt.Println("Введите с")
//	fmt.Scan(&c)
//	d = math.Pow(b, 2) - 4*a*c
//	fmt.Println("D =", d)
//	if d > 0 {
//		fmt.Println("Уровнения имеет 2 квд. корня")
//		x1 = (-b - math.Sqrt(d)) / (2 * a)
//		fmt.Println(-b)
//		fmt.Println(2 * a)
//		fmt.Println(math.Sqrt(d))
//		x2 = (-b + math.Sqrt(d)) / (2 * a)
//		fmt.Println("Первый корень уровнения:", x1)
//		fmt.Println("Первый корень уровнения:", x2)
//	} else if d == 0 {
//		fmt.Println("Уровнения имеет 1 квд. корень")
//		x1 = -b / 2 * a
//		fmt.Println("Корень уровнения:", x1)
//	} else {
//		fmt.Println("Корней нет")
//	}
//}

//func main() {
//	fmt.Println("Задание 7. Счастливый билет\n")
//	var tickerUser int
//	fmt.Println("Введите ваше число:")
//	fmt.Scan(&tickerUser)
//	var isLucky bool = false
//	var isMirror bool = false
//	if tickerUser%100 == tickerUser/100 {
//		isLucky = true
//	}
//	if tickerUser/1000 == tickerUser%10 && (tickerUser/100)%10 == (tickerUser%100)/10 {
//		isMirror = true
//	}
//	if isLucky {
//		if isMirror {
//			fmt.Println("Ваше число счастливое и зеркальное! Ого")
//		} else {
//			fmt.Println("Ваше число счастливое!")
//		}
//	} else if isMirror {
//		fmt.Println("Ваше число зеркальное!")
//	} else {
//		fmt.Println("У вас обычное число:(")
//	}
//}
func main() {
	fmt.Println("Задание 8 (по желанию). Игра «Угадай число»\n")
	var count int
	var guess, startGuesing bool = false, true
	var userAnswer string
	var startGuess int = 5
	for startGuesing != false {
		fmt.Println("Ваше число:", startGuess)
		count += 1
		fmt.Println("Я угадал? Варианты ответа: (+, -, =)") // + значит , что загаданное число больше; - значит, что загаданное число менье
		fmt.Scan(&userAnswer)
		if userAnswer == "+" || userAnswer == "-" {
			if userAnswer == "+" {
				startGuess += 1
			} else if userAnswer == "-" {
				startGuess -= 1
			}
		} else if userAnswer == "=" {
			guess = true
			startGuesing = false
		} else {
			fmt.Println("Я не понял. Начинайте заново")
			startGuesing = false
		}
		if guess {
			fmt.Println("Легко! отгадал за", count, "попыток(-ки)")
		}
		if count == 4 {
			fmt.Println("Не смог отгадать ваше число за", count, "попытки :(")
			startGuesing = false
		}
	}
}
