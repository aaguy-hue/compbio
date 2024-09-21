package main

func AffineScoreTable(str1, str2 string, match, mismatch, gapOpen, gapExtension int) (lower_mtx, middle_mtx, upper_mtx [][]int) {
	numRows := len(str1) + 1
	numCols := len(str2) + 1
	lower_mtx = InitializeMatrix[int](numRows, numCols)
	middle_mtx = InitializeMatrix[int](numRows, numCols)
	upper_mtx = InitializeMatrix[int](numRows, numCols)

	for i := 0; i < numCols; i++ {
		// lower_mtx[0][i] = -1e10
		middle_mtx[0][i] = -gapOpen - gapExtension*(i-1)
	}

	for i := 0; i < numRows; i++ {
		// upper_mtx[i][0] = -1e10
		middle_mtx[i][0] = -gapOpen - gapExtension*(i-1)
	}

	middle_mtx[0][0] = 0

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if r == 1 {
				lower_mtx[r][c] = middle_mtx[r-1][c] - gapOpen
			} else if r > 1 {
				lower_mtx[r][c] = Max(
					lower_mtx[r-1][c]-gapExtension,
					middle_mtx[r-1][c]-gapOpen,
				)
			}

			if c == 1 {
				upper_mtx[r][c] = middle_mtx[r][c-1] - gapOpen
			} else if c > 1 {
				upper_mtx[r][c] = Max(
					upper_mtx[r][c-1]-gapExtension,
					middle_mtx[r][c-1]-gapOpen,
				)
			}

			if r >= 1 && c >= 1 {
				score := ComputeScore(str1, str2, r, c, match, mismatch)
				middle_mtx[r][c] = Max(
					lower_mtx[r][c],
					upper_mtx[r][c],
					middle_mtx[r-1][c-1]+score,
				)
			}
		}
	}

	return // this will return the 3 matrices
}

func ComputeScore[T Number](str1, str2 string, r, c int, match, mismatch T) T {
	if str1[r-1] == str2[c-1] {
		return match
	} else {
		return -mismatch
	}
}
