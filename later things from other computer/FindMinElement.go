package main

type DistanceMatrix [][]float64

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your FindMinElement() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func FindMinElement(mtx DistanceMatrix) (int, int, float64) {
	minVal := mtx[0][1]
	minRow := 0
	minCol := 1

	for i := 0; i < len(mtx); i++ {
		for j := i + 1; j < len(mtx[0]); j++ {
			if mtx[i][j] < minVal && i != j {
				minVal = mtx[i][j]
				minRow = i
				minCol = j
			}
		}
	}

	return minRow, minCol, minVal
}
