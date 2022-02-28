package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Задание 1. Определение количества слов, начинающихся с большой буквы\n")
	var count int
	text := "евг А Г"
	//var text string = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	r := []rune(text)
	for len(r) > 0 {
		a := string(r[:1])
		if a == " " || a == "." || a == "," {
			r = r[1:]
			continue
		}
		if string(r[:1]) == strings.ToUpper(string(r[:1])) {
			fmt.Println(string(r[:1]))
			count++
		}
		r = r[1:]
	}
	fmt.Printf("Строка содержит %v слов с большой буквы.", count)
}

//func main() {
//	text := "a10 10 20b 20 30c30 30 dd"
//	for strings.Contains(text, " ") {
//		spaceIndex := strings.Index(text, " ")
//		a := text[:spaceIndex]
//		val, _ := strconv.ParseInt(a, 10, 0)
//		if val/10 >= 1 {
//			fmt.Println(val)
//		}
//		text = text[spaceIndex+1:]
//	}
//}
