package main

type GameBoard []([]bool)

func PlayGameOfLife(initialBoard GameBoard, numGens int) []GameBoard {
	boards := []GameBoard{initialBoard}
	for i := 0; i < numGens; i++ {
		boards = append(boards, UpdateBoard(boards[i]))
	}
	return boards
}

func UpdateBoard(currentBoard GameBoard) GameBoard {
	numRows := CountRows(currentBoard)
	numCols := CountCols(currentBoard)
	newBoard := InitializeBoard(numRows, numCols)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			newBoard[r][c] = UpdateCell(currentBoard, r, c)
		}
	}

	return newBoard
}

func InitializeBoard(numRows, numCols int) GameBoard {
	return make(GameBoard, numRows, numCols)
}

func UpdateCell(currentBoard GameBoard, r, c int) bool {
	numNeighbors := CountLiveNeighbors(currentBoard, r, c)
	if (currentBoard[r][c] && (numNeighbors == 2 || numNeighbors == 3)) ||
		(!currentBoard[r][c] && numNeighbors == 3) {
		return true
	}
	return false
}

func CountLiveNeighbors(board GameBoard, r, c int) int {
	var live int
	for i := Max2(r-1, 0); i <= Min2(r+1, len(board)-1); i++ {
		for j := Max2(c-1, 0); j <= Min2(c+1, len(board[0])-1); j++ {
			if board[i][j] && !(i == r && j == c) {
				live++
			}
		}
	}
	return live
}

func InField(board GameBoard, i, j int) bool {
	if i >= 0 && i < CountRows(board) && j >= 0 && j < CountCols(board) {
		return true
	}
	return false
}

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
