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
	size := len(os.Args[1:]) // Getting len of flags
	firstFile, secondFile, resultFile := "", "", ""
	flag.StringVar(&firstFile, "firstFile", "", "give FirstFile")
	flag.StringVar(&secondFile, "secondFile", "", "give SecondFile")
	flag.StringVar(&resultFile, "resultFile", "", "give name for result file")

	flag.Parse()
	fmt.Println(size, firstFile, secondFile, resultFile)
	if firstFile == "" && secondFile == "" {
		return
	} else {
		countFiles(size, firstFile, secondFile, resultFile)
	}
}
func countFiles(size int, firstFile, secondFile, resultFile string) {
	if size == 6 {
		data, err := contactContent(firstFile, secondFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		makeFile(data, resultFile)
	} else {
		showResult(size, firstFile, secondFile)
	}
}

func showResult(size int, filename ...string) {
	size /= 2
	fmt.Println(filename[0])
	fmt.Println("1")
	if size == 1 {
		result, err := getFileContent(filename[0])
		if err != nil {
			return
		}
		fmt.Println(result)
	} else {
		firstResult, err := getFileContent(filename[0])
		secondFileResult, err2 := getFileContent(filename[1])
		if err != nil || err2 != nil {
			return
		}
		fmt.Println(firstResult + "\n" + secondFileResult)
	}
}

func getFileContent(filename string) (string, error) {
	/*
		Getting data from requested file
	*/
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	return contents, nil
}

func contactContent(firstFile, secondFile string) (string, error) {
	/*
		Getting data from files and connectig it
	*/
	dataOfFirstFile, err := getFileContent(firstFile)
	dataOfSecondFile, err2 := getFileContent(secondFile)
	if err != nil || err2 != nil {
		return "", err
	}
	strArray := make([]string, 0, 2)
	strArray = append(strArray, dataOfFirstFile)
	strArray = append(strArray, dataOfSecondFile)
	return strings.Join(strArray, "-"), nil
}

func makeFile(data string, resultFile string) {
	/*
		Creating file with data of two files
	*/
	f, err := os.Create(resultFile)
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
