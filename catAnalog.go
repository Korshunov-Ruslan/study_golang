package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		The program receives the names of two files as input and concatenate their contents.
	*/
	fmt.Println(len(os.Args), "1")
	FileName := "first.txt" // Write File name
	firstFile, secondFile := "", ""
	flag.StringVar(&firstFile, "FirstFile", "", "give FirstFile")
	flag.StringVar(&secondFile, "SecondFile", "", "give SecondFile")
	flag.Parse()
	//if firstFile == "" || secondFile == ""
	if len(os.Args) < 5 {
		countFiles(FileName)
	} else {
		makeFile(contactContent(firstFile, secondFile))
	}

}
func countFiles(filename ...string) {
	if len(filename) <= 2 {
		showResult(len(filename), filename)
	} else if len(filename) == 0 {
		fmt.Println("No files")
		return
	} else if len(filename) > 2 {
		fmt.Println("Two files max")
		return
	}
}

func showResult(len int, filename ...[]string) {
	if len == 1 {
		result := getFileContent(filename[0][0])
		fmt.Println(result)
	} else {
		firstResult := getFileContent(filename[0][0])
		secondFileResult := getFileContent(filename[0][1])
		fmt.Println(firstResult + "\n" + secondFileResult)
	}
}

func getFileContent(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	return contents
}

func contactContent(firstFile, secondFile string) string {
	dataOfFirstFile, dataOfSecondFile := getFileContent(firstFile), getFileContent(secondFile)
	strArray := make([]string, 0, 2)
	strArray = append(strArray, dataOfFirstFile)
	strArray = append(strArray, dataOfSecondFile)
	return strings.Join(strArray, "-")
}

func makeFile(data string) {
	f, err := os.Create("cat-analog-result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err2 := f.WriteString(data)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Done")
}
