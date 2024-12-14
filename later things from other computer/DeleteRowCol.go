package main

type DistanceMatrix [][]float64

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your DeleteRowCol() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func DeleteRowCol(mtx DistanceMatrix, row, col int) DistanceMatrix {
	// assume col > row
	for i := range mtx {
		mtx[i] = append(mtx[i][:col], mtx[i][col+1:]...)
		mtx[i] = append(mtx[i][:row], mtx[i][row+1:]...)
	}
	mtx = append(mtx[:col], mtx[col+1:]...)
	mtx = append(mtx[:row], mtx[row+1:]...)
	return mtx
}
