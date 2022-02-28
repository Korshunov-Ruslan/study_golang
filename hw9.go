package main

//func main() {
//	fmt.Println("Задание 1. Переполнение\n")
//	var value8 uint8 = 0
//	var value16 uint16 = 0
//	var value32 uint32 = 0
//	var count8, count16 int
//	for i := 0; i != 1<<32-1; i++ {
//		value8++
//		value16++
//		value32++
//		switch value8 {
//		case 1<<8 - 1:
//			count8++
//		}
//		switch value16 {
//		case 1<<16 - 1:
//			count16++
//
//		}
//	}
//	fmt.Printf("Тип %T переполнился %v раз\nТип %T переполнился %v раз", value8, count8, value16, count16)
//}

//func main() {
//	fmt.Println("Задание 2. Минимальный тип данных\n")
//	var firstUserInput, secondUserInput int16
//	fmt.Scan(&firstUserInput, &secondUserInput)
//	result := int32(firstUserInput) * int32(secondUserInput)
//	fmt.Println(result)
//	switch {
//	case result >= math.MinInt8 && result <= math.MaxInt8:
//		fmt.Println("Это int8")
//	case result <= math.MaxUint8 && result >= 0:
//		fmt.Println("Это uint8")
//	case result >= math.MinInt16 && result <= math.MaxInt16:
//		fmt.Println("Это int16")
//	case result <= math.MaxUint16 && result >= 0:
//		fmt.Println("Это uint16")
//	case result >= math.MinInt32 && result <= math.MaxInt32:
//		fmt.Println("Это int32")
//	case uint32(result) <= math.MaxUint32 && result >= 0:
//		fmt.Println("Это uint32")
//	}
//}
