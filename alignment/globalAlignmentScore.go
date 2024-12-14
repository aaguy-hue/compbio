package main

// GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	// initialize the matrix
	arr := make([][]float64, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		arr[i] = make([]float64, len(str2)+1)
	}

	// assign the top row scores
	for i := 1; i < len(str2)+1; i++ {
		arr[0][i] = gap * float64(-i)
	}

	// assign the left column scores
	for i := 1; i < len(str1)+1; i++ {
		arr[i][0] = gap * float64(-i)
	}

	// assign scores for each node
	for r := 1; r < len(str1)+1; r++ {
		for c := 1; c < len(str2)+1; c++ {
			up := arr[r-1][c] - gap
			left := arr[r][c-1] - gap
			diag := arr[r-1][c-1]

			if str1[r-1] == str2[c-1] {
				diag += match
			} else {
				diag -= mismatch
			}

			arr[r][c] = Max(up, left, diag)
		}
	}

	return arr
}
