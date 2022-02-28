package main

import "fmt"

func main() {
	fmt.Println("Введите название планеты:")
	var planetName string
	fmt.Scan(&planetName)
	fmt.Println("Введите название звёздной системы:")
	var starSytemName string
	fmt.Scan(&starSytemName)
	fmt.Println("Введите имя рейнджера:")
	var rangerName string
	fmt.Scan(&rangerName)
	fmt.Println("Введите вознаграждение:")
	var money int
	fmt.Scan(&money)
	fmt.Println(money)

	fmt.Print("Как вам, ", rangerName, ", известно, мы раса мирная, поэтому на наших военных кораблях используются наёмники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты ", planetName, " системы ", starSytemName, ".", "\nНо случилась неприятность: в связи с большими потерями в последнее время престиж профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар планеты ", planetName, ", впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!\n", rangerName, ", вы должны прибыть на планету ", planetName, " системы ", starSytemName, " и помочь выполнить план призыва. Мы готовы выплатить вам премию в ", money, " кредитов за эту маленькую услугу.")
}

func main() {
	fmt.Println("Программа Симулятор маршрутки.")
	var busTicket int = 20 // Стоимость билета
	var busTicketsSold int // Общее кол-во проданных билетов
	var passagersNow int   // текущее кол-во пассажиров в маршрутке
	var passagersIn int    // Кол-во пассажиров, которое вошло
	var passagersOut int   // Кол-во пассажиров, которое вышло
	var busStopCount int   // Всего остановок
	fmt.Println("Введите кол-во остановок:")
	fmt.Scan(&busStopCount)
	for i := 0; i < busStopCount; i++ {
		fmt.Println("Остановка №", i)
		fmt.Println("В салоне ", passagersNow, "пассажиров")
		fmt.Println("Сколько пассажиров зашло?")
		fmt.Scan(&passagersIn)
		fmt.Println("Сколько пассажиров вышло?")
		fmt.Scan(&passagersOut)
		passagersNow = passagersNow + passagersIn - passagersOut
		busTicketsSold += busTicket * passagersIn
		fmt.Println("---Едем---")
	}
	var salaryCost int = busTicketsSold / 4 // Траты на зарплату
	var fuelCost int = busTicketsSold / 5   // Траты на бензин
	var taxesCost int = busTicketsSold / 5  // Траты на налоги
	var carRepair int = busTicketsSold / 5  // Траты на ремонт
	var totalIncome int = busTicketsSold - (salaryCost + fuelCost + taxesCost + carRepair)
	fmt.Println("Всего заработали:", busTicketsSold)
	fmt.Println("Зарплата водителя:", salaryCost)
	fmt.Println("Расходы на топливо:", fuelCost)
	fmt.Println("Расходы на ремонт:", carRepair)
	fmt.Println("Расходы на налоги:", taxesCost)
	fmt.Println("Доход после вычета:", totalIncome)
}

func main() {
	fmt.Println("Программа меняет местами значения")
	a := 50
	b := 70
	fmt.Println("Сейчас a=", a, ", a б=", b)
	a = b
	b = a
	fmt.Println("А сейчас a=", a, ", a б=", b)
}

func main() {
	fmt.Println("Задание 4. Задачи 1 и 2")
	startHeight := 100
	dayGrowth := 50
	nightLoses := 20
	twoDaysLoses := nightLoses * 2
	threeDaysGrowth := dayGrowth*2 + dayGrowth/2
	thirdDayHeight := threeDaysGrowth - twoDaysLoses
	fullDayChanges := dayGrowth - nightLoses
	sellTreeDaysCount := (300 - startHeight) / fullDayChanges
	fmt.Println("Высота бамбука в середине третьего дня:", thirdDayHeight)
	fmt.Println("Полных дней, чтобы бамбук можно было продать:", sellTreeDaysCount)

}

func main() {
	fmt.Println("Задание 4. Задача 3")
	var startHeight int
	var dayGrowth int
	var nightLoses int
	var dayCount int
	fmt.Println("Введите стартовую высоту дерева:")
	fmt.Scan(&startHeight)
	fmt.Println("Введите рост за день:")
	fmt.Scan(&dayGrowth)
	fmt.Println("Введите скорость поедания гусенец:")
	fmt.Scan(&nightLoses)
	fmt.Println("Введите кол-во дней, за которое созреет бамбук")
	fmt.Scan(&dayCount)
	fullDayChanges := dayGrowth - nightLoses
	treeHeightFinalDay := fullDayChanges*dayCount + 100
	fmt.Println("Высота бамбука на", dayCount, "день составит:", treeHeightFinalDay)
}

func main() {
	fmt.Println("Задание 4. Задача 4")
	var startHeight int
	var dayGrowth int
	var nightLoses int
	var neededHeight int
	fmt.Println("Введите стартовую высоту дерева:")
	fmt.Scan(&startHeight)
	fmt.Println("Введите рост за день:")
	fmt.Scan(&dayGrowth)
	fmt.Println("Введите скорость поедания гусенец:")
	fmt.Scan(&nightLoses)
	fmt.Println("Введите целевую высоту бамбука")
	fmt.Scan(&neededHeight)
	fullDayChanges := dayGrowth - nightLoses
	dayForNeededHeight := (neededHeight - startHeight) / fullDayChanges
	fmt.Print("Нужная высота бамбука <", neededHeight, "> см будет на ", dayForNeededHeight, " день.")
}
