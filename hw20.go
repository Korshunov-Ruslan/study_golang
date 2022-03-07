package main

import "fmt"

const rows = 3
const cols = 3

const oRows = 3
const oCols = 5
const nCols = 4

//func getDetermine(input [rows][cols]int) int {
//	first := input[1][1]*input[2][2] - input[1][2]*input[2][1]
//	second := input[1][0]*input[2][2] - input[1][2]*input[2][0]
//	third := input[1][0]*input[2][1] - input[1][1]*input[2][0]
//	result := input[0][0]*first - input[0][1]*second + input[0][2]*third
//
//	return result
//}
//func main() {
//	/*
//		Get matrix determinant (3x3)
//	*/
//	matrix := [rows][cols]int{
//		{1, 4, -2},
//		{1, 2, 1},
//		{3, 2, 3},
//	}
//	fmt.Println(getDetermine(matrix))
//}

func main() {
	/*
		Matrix multiplication
	*/
	firstMatrix := [oRows][oCols]int{
		{1, -3, 5, 7, 100},
		{11, 0, 4, 6, 7},
		{3, 1, -1, -2, -1},
	}
	secondMatrix := [oCols][nCols]int{
		{1, 1, 2, -1},
		{100, 100, 100, 100},
		{2, -1, -5, -5},
		{1, 0, 0, 1},
		{-3, 1, 5, 3},
	}
	fmt.Println(matrixMulti(firstMatrix, secondMatrix))
}
func matrixMulti(firstMat [oRows][oCols]int, secondMat [oCols][nCols]int) [oRows][nCols]int {
	var newMatrix [oRows][nCols]int
	for i := 0; i < oRows; i++ {
		fmt.Println(firstMat[i])
		for j := 0; j < nCols; j++ {
			for k := 0; k < oCols; k++ {
				newMatrix[i][j] = newMatrix[i][j] + firstMat[i][k]*secondMat[k][i]
			}
		}
	}
	return newMatrix
}
