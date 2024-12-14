package main

type DistanceMatrix [][]float64

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your AddRowCol() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func AddRowCol(row, col, clusterSize1, clusterSize2 int, mtx DistanceMatrix) DistanceMatrix {
	var newRow []float64 = make([]float64, len(mtx[0])+1)
	for r := 0; r < len(newRow)-1; r++ {
		if r == row || r == col {
			newRow[r] = 0
			continue
		}
		newRow[r] = (float64(clusterSize1)*mtx[row][r] + float64(clusterSize2)*mtx[r][col]) / float64(clusterSize1+clusterSize2)
	}
	mtx = append(mtx, newRow)
	for c := 0; c < len(newRow)-1; c++ {
		if c == col || c == row {
			mtx[c] = append(mtx[c], 0)
			continue
		}
		mtx[c] = append(mtx[c], newRow[c])
	}
	return mtx
}
