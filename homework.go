package main

import (
	"fmt"
)

func main() {
	var lap int = 4
	var engine int = 254
	var wheels int = 93
	var steeringWheel int = 49
	var wind int = 21
	var rain int = 17
	var speed int = engine + wheels + steeringWheel - wind - rain

	fmt.Println("===================")
	fmt.Print("Супергонки. Круг ", lap, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speed, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колёса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: −", wind, "\n")
	fmt.Print("Дождь: −", rain, "\n")
	second()
}

func second() {
	productPrice := 6400
	delieveryPrice := 350
	discount := 700
	finalPrice := productPrice + delieveryPrice - discount

	fmt.Println("Стоимость товара:", productPrice)
	fmt.Println("Стоимость доставки:", delieveryPrice)
	fmt.Println("Размер скидки:", discount)
	fmt.Println("---------")
	fmt.Println("Итого:", finalPrice)
}

func third() {
	workingShiftMinutes := 600
	makeOrderMinutes := 10
	collectOrderMinutes := 5
	maintenanceTime := workingShiftMinutes / (makeOrderMinutes + collectOrderMinutes)
	fmt.Println("За смену в", workingShiftMinutes, "минут кассир сможет обслужить", maintenanceTime, "клиентов")
}

func fouth() {
	costEveryone := 1000000
	entranceCount := 1
	appartamentCount := 10
	costPerAppartament := costEveryone / (entranceCount + appartamentCount)
	fmt.Println("Стоимость ремонта для одной квартиры:", costPerAppartament, "рублей.")
}
