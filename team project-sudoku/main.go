package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

func findEmptyCell(board [][]byte) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j, false
			}
		}
	}
	return 0, 0, true
}

var solutionCount int

func solveSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				temp := board[i][j]
				board[i][j] = '.'
				if !isValid(board, i, j, temp) {
					return false
				}
				board[i][j] = temp
			}
		}
	}

	row, col, isEmpty := findEmptyCell(board)

	if isEmpty {
		solutionCount++
		return solutionCount == 1
	}

	for num := byte('1'); num <= '9'; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if !solveSudoku(board) {
				board[row][col] = '.'
				continue
			}
			if solutionCount > 1 {
				return false
			}
		}
	}
	return solutionCount == 1
}

func printError() {
	strErr := "Error"
	for _, r := range strErr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	if len(args) != 9 {
		printError()
		return
	}

	board := make([][]byte, 9)
	for i, line := range args {
		if len(line) != 9 {
			printError()
			return
		}
		board[i] = []byte(line)
	}

	if !solveSudoku(board) {
		printError()
		return
	}

	for _, row := range board {
		for i, cell := range row {
			z01.PrintRune(rune(cell))
			if i != len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
