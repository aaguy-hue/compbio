package main

// BinarizeMatrix takes a matrix of values and a threshold.
// It binarizes the matrix according to the threshold.
// If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
	numRows := len(mtx)
	numCols := len(mtx[0])
	binarizedMtx := InitializeMatrix[int](numRows, numCols)

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if mtx[i][j] >= threshold {
				if mtx[i][j] > mtx[j][i] {
					binarizedMtx[i][j] = 1
				} else if mtx[i][j] == mtx[j][i] && i < j { // if they're equal, choose the one in the top right corner
					binarizedMtx[i][j] = 1
				}
			} else {
				binarizedMtx[i][j] = 0
			}
		}
	}
	return binarizedMtx
}

func InitializeMatrix[ElementType any](rows, cols int) [][]ElementType {
	arr := make([][]ElementType, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]ElementType, cols)
	}
	return arr
}
