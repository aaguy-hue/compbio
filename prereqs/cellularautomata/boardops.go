package main

type GameBoard []([]int)

func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}

func InitializeBoard(numRows, numCols int) GameBoard {
	hm := make(GameBoard, numRows)
	for i := 0; i < numRows; i++ {
		hm[i] = make([]int, numCols)
	}
	return hm
}

func InField(board GameBoard, row, col int) bool {
	if row >= 0 && row < CountRows(board) && col >= 0 && col < CountCols(board) {
		return true
	}
	return false
}
