package main

//func main() {
//	fmt.Println("Задание 1. Работа с файлами\n")
//	t := time.Now()
//	neededTimeFormat := t.Format("2006-01-02 15:04:05")
//	file, _ := os.Create("1stex")
//	defer file.Close()
//	var userInput string
//	var count int
//	for {
//		fmt.Scan(&userInput)
//		fmt.Println(userInput)
//		if userInput == "exit" {
//			return
//		} else {
//			count++
//			str := strconv.Itoa(count) + " " + neededTimeFormat + " " + userInput + "\n"
//			file.WriteString(str)
//		}
//	}
//}

//func main() {
//	file, err := os.Open("1stex")
//	if err != nil {
//		fmt.Println("Cant open the file", err)
//		return
//	}
//	defer file.Close()
//	fs, _ := file.Stat()
//	if fs.Size() == 0 {
//		fmt.Println("Requested file is empty")
//		return
//	}
//	buf := make([]byte, fs.Size())
//	if _, err := io.ReadFull(file, buf); err != nil {
//		fmt.Println("Cant read bytes", err)
//		return
//	}
//	fmt.Printf("%s\n", buf) }

//func main() {
//	fmt.Println("Уровни доступа")
//	// Trying to create file with 0 access level (only for reading)
//	filename := "0444_accesslvl.txt"
//	file, _ := os.Create(filename)
//	if err := os.Chmod(filename, 0444); err != nil {
//		fmt.Println(err)
//	}
//	file.WriteString("avc")
//	file.WriteString("asss")
//	defer file.Close()
//}

//func main() {
//	fmt.Println("Пакет ioutil, создание")
//	var userInput string
//	filename := "ioutill.txt"
//	ioutil.WriteFile(filename, []byte(filename), 0644)
//	sum := ""
//	for {
//		fmt.Scan(&userInput)
//		if userInput == "exit" {
//			break
//		} else {
//			sum += userInput
//			sum += "\n"
//		}
//		text := []byte(sum)
//		err := ioutil.WriteFile(filename, text, 0644)
//		if err != nil {
//			fmt.Println("Cant create file", err)
//			break
//		}
//	}
//	file, _ := os.Open(filename)
//	esultBytes, _ := ioutil.ReadAll(file)
//	fmt.Println(string(esultBytes))
//	defer file.Close()
//}
