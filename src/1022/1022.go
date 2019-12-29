package main

import (
	"fmt"
)

func indexMapFactory(r1, c1 int) (func(int) int, func(int) int) {
	return func(r int) int {
			return r - r1
		},
		func(c int) int {
			return c - c1
		}
}

func isValidIndexFacotry(rowSize, colSize int) func(int, int) bool {
	return func(r, c int) bool {
		if (r >= 0 && r < rowSize) && (c >= 0 && c < colSize) {
			return true
		} else {
			return false
		}
	}
}

func printFormatingNumberFactory(maxNum int) func(int) {
	if maxNum >= 1000000000 {
		return func(num int) {
			fmt.Printf("%10d", num)
		}
	} else if maxNum >= 100000000 {
		return func(num int) {
			fmt.Printf("%9d", num)
		}
	} else if maxNum >= 10000000 {
		return func(num int) {
			fmt.Printf("%8d", num)
		}
	} else if maxNum >= 1000000 {
		return func(num int) {
			fmt.Printf("%7d", num)
		}
	} else if maxNum >= 100000 {
		return func(num int) {
			fmt.Printf("%6d", num)
		}
	} else if maxNum >= 10000 {
		return func(num int) {
			fmt.Printf("%5d", num)
		}
	} else if maxNum >= 1000 {
		return func(num int) {
			fmt.Printf("%4d", num)
		}
	} else if maxNum >= 100 {
		return func(num int) {
			fmt.Printf("%3d", num)
		}
	} else if maxNum >= 10 {
		return func(num int) {
			fmt.Printf("%2d", num)
		}
	} else {
		return func(num int) {
			fmt.Printf("%1d", num)
		}
	}
}

func printMatrix(matrix [][]int, maxNum int) {
	printFormatingNumber := printFormatingNumberFactory(maxNum)
	for _, vec := range matrix {
		for i, ele := range vec {
			printFormatingNumber(ele)
			if i < len(vec) {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	var r1, c1, r2, c2 int

	fmt.Scanln(&r1, &c1, &r2, &c2)

	rowSize := r2 - r1 + 1
	colSize := c2 - c1 + 1

	// Make Matrix
	matrix := make([][]int, rowSize)
	for i := 0; rowSize > i; i++ {
		matrix[i] = make([]int, colSize)
	}

	rowMap, colMap := indexMapFactory(r1, c1)
	isValidIndex := isValidIndexFacotry(rowSize, colSize)

	rowIndex, colIndex := rowMap(0), colMap(0)

	prevVerticalMove := 0
	prevHorizontalMove := 0

	isVerticalMove := true
	upFlag := true
	leftFlag := true
	matrixSize := rowSize * colSize

	// 여기 darwCount 검사 안해서 1, 1사이즈일때 뻘짓 개많이함
	drawCount := 0
	count := 1
	if isValidIndex(rowIndex, colIndex) {
		matrix[rowIndex][colIndex] = count
		drawCount++
	}
	if drawCount >= matrixSize {
		printMatrix(matrix, count)
		return
	}
	prevVerticalMove++
	prevHorizontalMove++
	count++

	colIndex++
	if isValidIndex(rowIndex, colIndex) {
		matrix[rowIndex][colIndex] = count
		drawCount++
	}
	if drawCount >= matrixSize {
		printMatrix(matrix, count)
		return
	}
	prevHorizontalMove++
	count++

ALL:
	for {
		if isVerticalMove {

			for i := 0; i < prevVerticalMove; i++ {
				if upFlag {
					rowIndex--
				} else {
					rowIndex++
				}

				if isValidIndex(rowIndex, colIndex) {
					matrix[rowIndex][colIndex] = count
					drawCount++
					if drawCount >= matrixSize {
						break ALL
					}
				}

				count++
			}
			prevVerticalMove++
			upFlag = !upFlag
		} else {

			for i := 0; i < prevHorizontalMove; i++ {
				if leftFlag {
					colIndex--
				} else {
					colIndex++
				}

				if isValidIndex(rowIndex, colIndex) {
					matrix[rowIndex][colIndex] = count
					drawCount++
					if drawCount >= matrixSize {
						break ALL
					}
				}
				count++
			}
			prevHorizontalMove++
			leftFlag = !leftFlag
		}
		isVerticalMove = !isVerticalMove

	}
	printMatrix(matrix, count)
}

/*
순간이동 할 경우 3가지
1. 안쪽에서 밖으로 벗어날 때 (남은 개수 전부 이동)
2. 밖쪽에서 안쪽으로 들어갈 떄 (column까지 이동)
3. 전혀 상관 없는 수직 방면 (남은 개수 전부 이동)
*/
