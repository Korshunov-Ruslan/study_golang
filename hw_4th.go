package main

import "fmt"

//func main() {
//	fmt.Println("_____Задание №1_____")
//	var firstResult int
//	var secondResult int
//	var thirdResult int
//	var decision int
//	var neededResult int = 275 // Необходимое кол-во баллов для поступления
//	fmt.Println("Введите результаты первого экзамена:")
//	fmt.Scan(&firstResult)
//	fmt.Println("Введите результаты второго экзамена:")
//	fmt.Scan(&secondResult)
//	fmt.Println("Введите результаты третьего экзамена:")
//	fmt.Scan(&thirdResult)
//	decision = firstResult + secondResult + thirdResult
//	if decision >= neededResult {
//		fmt.Println("Кол-во ваших баллов:", decision)
//		fmt.Println("Поздравляем, вы прошли!")
//	} else {
//		fmt.Println("Кол-во ваших баллов:", decision)
//		fmt.Println("Вы не прошли")
//	}
//}

//func main() {
//	fmt.Println("Задание 2. Три числа\n")
//	var firstValue int
//	var secondValue int
//	var thirdValue int
//	fmt.Println("Введите первое число:")
//	fmt.Scan(&firstValue)
//	fmt.Println("Введите второе число:")
//	fmt.Scan(&secondValue)
//	fmt.Println("Введите третье число:")
//	fmt.Scan(&thirdValue)
//	if firstValue > 5 && secondValue > 5 && thirdValue > 5 {
//		fmt.Println("Среди ваших чисел есть число, больше 5")
//	} else {
//		fmt.Println("Все числа меньше или равно 5")
//	}
//}

//func main() {
//	fmt.Println("Задание 3. Банкомат\n")
//	var maxWithdraw int = 100000
//	var nominalValue int = 100
//	var userWithdraw int
//	fmt.Println("Введите сумму снятия со счёта:")
//	fmt.Scan(&userWithdraw)
//	if userWithdraw > maxWithdraw {
//		fmt.Println("Столько снять не получится. Лимит на снятие:", maxWithdraw, "рублей.")
//
//	} else {
//		if userWithdraw%nominalValue != 0 {
//			fmt.Println("Банкомат может выдавать сумму кратную", nominalValue, "рублей. Попробуйте ещё раз")
//		} else {
//			fmt.Println("Операция успешно выполнена. \nВы сняли", userWithdraw, "рублей.")
//		}
//	}
//}

//func main() {
//	fmt.Println("Задание 4. Три числа: ещё попытка\n")
//	var firstValue int
//	var secondValue int
//	var thirdValue int
//	fmt.Println("Введите первое число:")
//	fmt.Scan(&firstValue)
//	fmt.Println("Введите второе число:")
//	fmt.Scan(&secondValue)
//	fmt.Println("Введите третье число:")
//	fmt.Scan(&thirdValue)
//	var count = 0
//	if firstValue >= 5 {
//		count += 1
//		if secondValue >= 5 {
//			count += 1
//			if thirdValue >= 5 {
//				count += 1
//				fmt.Println("Среди введённых чисел <<", count, ">> числа больше или равны 5.")
//			}
//		}
//	} else if secondValue >= 5 {
//		count += 1
//		if thirdValue >= 5 {
//			count += 1
//			fmt.Println("Среди введённых чисел <<", count, ">> числа больше или равны 5.")
//		}
//	} else if thirdValue >= 5 {
//		count += 1
//		fmt.Println("Среди введённых чисел <<", count, ">> числа больше или равны 5.")
//	} else {
//		fmt.Println("Все числа менье 5.")
//	}
//
//}

//func main() {
//	fmt.Println("Задание 5. Ресторан\n")
//	var dayWeek, guestCount, totalIncome, discount, additionalMoney int
//
//	fmt.Println("Укажите день недели")
//	fmt.Scan(&dayWeek)
//	fmt.Println("Укажите число гостей")
//	fmt.Scan(&guestCount)
//	fmt.Println("Укажите сумму к оплате")
//	fmt.Scan(&totalIncome)
//	if dayWeek == 1 && guestCount > 5 {
//		additionalMoney = totalIncome / 10
//		discount = totalIncome / 10
//		fmt.Println("Так как сейчас понедельник, получайте скидку.")
//		fmt.Println("Скидка:", discount)
//		fmt.Println("Надбавка на обслуживание:", additionalMoney)
//		fmt.Println("Сумма к оплате:", totalIncome-discount+additionalMoney)
//	} else if dayWeek == 1 && guestCount <= 5 {
//		discount = totalIncome / 10
//		fmt.Println("Так как сейчас понедельник, получайте скидку.")
//		fmt.Println("Скидка:", discount)
//		fmt.Println("Сумма к оплате:", totalIncome-discount)
//	} else if dayWeek == 5 && totalIncome > 10000 {
//		if guestCount <= 5 {
//			discount = totalIncome / 20
//			fmt.Println("Сейчас пятница и ваш чек больше 10000, получайте скидку")
//			fmt.Println("Скидка:", discount)
//			fmt.Println("Сумма к оплате:", totalIncome-discount)
//		} else {
//			discount = totalIncome / 20
//			additionalMoney = totalIncome / 10
//			fmt.Println("Так как сейчас понедельник, получайте скидку.")
//			fmt.Println("Скидка:", discount)
//			fmt.Println("Надбавка на обслуживание:", additionalMoney)
//			fmt.Println("Сумма к оплате:", totalIncome-discount+additionalMoney)
//		}
//	} else if guestCount > 5 {
//		additionalMoney = totalIncome / 10
//		fmt.Println("Надбавка на обслуживание:", additionalMoney)
//		fmt.Println("Сумма к оплате:", totalIncome-discount+additionalMoney)
//	} else {
//		fmt.Println("Сумма к оплате:", totalIncome-discount+additionalMoney)
//	}
//}

func main() {
	var N int
	var K int
	var number int
	var count int = 1
	var extraStudents int = 0
	fmt.Println("Введите кол-во студентов")
	fmt.Scan(&N)
	fmt.Println("Введите кол-во групп")
	fmt.Scan(&K)
	fmt.Println("Введите порядковый номер студента")
	fmt.Scan(&number)
	var studentPerGroup int = N / K
	for i := 1; i <= K; i++ { // тут я перебираю каждую группы
		if number > K*i {
			count += 1 // Узнаю, в какой именно группе будет этот студент
		}
	}
	for o := 1; o <= N%K; o++ {
		extraStudents += 1
	}
	if extraStudents >= count {
		if count-1 <= 0 {
			fmt.Println("Студент с порядковым номером <<", number, ">> находится в <<", count, ">> группе")
		} else {
			fmt.Println("Студент с порядковым номером <<", number, ">> находится в <<", count-1, ">> группе")
		}
	} else if count > K {
		fmt.Println("Студент с порядковым номером <<", number, ">> находится в <<", count-1, ">> группе.------") // печатаю значение
	} else {
		fmt.Println("Студент с порядковым номером <<", number, ">> находится в <<", count, ">> группе")
	}
	if N%K >= 1 {
		fmt.Println("В первых", N%K, "группах", studentPerGroup+1, "человек", ",а в последних группах", studentPerGroup, "человек")
	} else {
		fmt.Println("В каждой группе по <<", studentPerGroup, ">> человек")
	}

}

//func main() {
//	var taxes int
//	var userIncome int
//	var poorTax int = 10000 / 100 * 13
//	fmt.Println("Задача 7")
//	fmt.Println("Введите свой доход")
//	fmt.Scan(&userIncome)
//
//	if userIncome > 10000 {
//		if userIncome > 50000 {
//			thirtyTax := (userIncome - 50000) / 100 * 30
//			twentyTax := (50000 - 10000) / 100 * 20
//			taxes = poorTax + thirtyTax + twentyTax
//			fmt.Println("Ваши налоги:", taxes)
//		} else if userIncome > 10000 && userIncome < 50000 {
//			twentyTax := (userIncome - 10000) / 100 * 20
//			taxes = poorTax + twentyTax
//			fmt.Println("Ваши налоги:", taxes)
//		}
//	} else {
//		fmt.Println("Вы без  налогов, гуляйте")
//	}
//}
