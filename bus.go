package main

import "fmt"

func main() {
	var userValue int
	var robotValue int = 3
	var answer string
	fmt.Println("Введите своё число (от 1 до 5)")
	fmt.Scan(&userValue)
	fmt.Println("Предполагаю, что ваше число =", robotValue)
	if userValue != robotValue {
		fmt.Println("Больше или меньше?")
		fmt.Scan(&answer)
		if answer == "больше" {
			robotValue += 1
			fmt.Println("Предполагаю, что ваше число =", robotValue)
			if robotValue != userValue {
				robotValue += 1
				fmt.Println("Ошибся, бля. Ваше число:", robotValue)
			}
		} else {
			robotValue -= 1
			fmt.Println("Предполагаю, что ваше число =", robotValue)
			if robotValue != userValue {
				robotValue -= 1
				fmt.Println("Ошибся бл ваше число =", robotValue)
			}
		}
	} else {
		fmt.Println("Изи, отгадал с первой попытки. Ваше число:", robotValue)
	}
}
