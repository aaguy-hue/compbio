package main

import "fmt"

func PrintBoard(board GameBoard) {
	for r := 0; r < CountRows(board); r++ {
		for c := 0; c < CountCols(board); c++ {
			PrintCell(board[r][c])
		}
		fmt.Print("\n")
	}
}

func PrintCell(cell bool) {
	if cell {
		fmt.Print("⬛")
	} else {
		fmt.Print("⬜")
	}
}
