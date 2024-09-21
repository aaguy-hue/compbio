package main

// LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for local alignment with these penalties.
func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	arr := make([][]float64, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		arr[i] = make([]float64, len(str2)+1)
	}

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

			arr[r][c] = Max(0, up, left, diag)
		}
	}

	return arr
}
