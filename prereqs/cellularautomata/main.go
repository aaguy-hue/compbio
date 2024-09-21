package main

import "fmt"

func PlayAutomaton(initialBoard GameBoard, numGens int, neighborhoodType string, rules map[string]int) []GameBoard {
	var boards []GameBoard = make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1], neighborhoodType, rules)
	}

	return boards
}

func UpdateBoard(currentBoard GameBoard, neighborhoodType string, rules map[string]int) GameBoard {
	var numRows int = CountRows(currentBoard)
	var numCols int = CountCols(currentBoard)
	var newBoard GameBoard = InitializeBoard(numRows, numCols)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			newBoard[r][c] = UpdateCell(currentBoard, r, c, neighborhoodType, rules)
		}
	}

	return newBoard
}

func UpdateCell(currentBoard GameBoard, row, col int, neighborhoodType string, rules map[string]int) int {
	neighborhood := NeighborhoodToString(currentBoard, row, col, neighborhoodType)
	return rules[neighborhood]
}

func NeighborhoodToString(currentBoard GameBoard, r, c int, neighborhoodType string) string {
	neighborhood := fmt.Sprint(currentBoard[r][c])
	var offsets [][2]int

	if neighborhoodType == "Moore" {
		offsets = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	} else if neighborhoodType == "vonNewmann" {
		offsets = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	} else {
		panic("Invalid neighborhood type " + neighborhoodType + "!")
	}

	for i := 0; i < len(offsets); i++ {
		x := offsets[i][0]
		y := offsets[i][1]
		if InField(currentBoard, r+x, c+y) {
			neighborhood += fmt.Sprint(currentBoard[r+x][c+y])
		} else {
			neighborhood += "0"
		}
	}

	return neighborhood
}
