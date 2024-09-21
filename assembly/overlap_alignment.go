package main

//ALL PENALTIES POSITIVE

// ScoreOverlapAlignment takes two strings along with match, mismatch, and gap penalties.
// It returns the maximum score of an overlap alignment in which str1 is overlapped with str2.
// Assume we are overlapping a suffix of str1 with a prefix of str2.
func ScoreOverlapAlignment(str1, str2 string, match, mismatch, gap float64) float64 {
	if len(str1) == 0 || len(str2) == 0 {
		panic("Empty string passed into ScoreOverlapAlignment")
	}

	arr := make([][]float64, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		arr[i] = make([]float64, len(str2)+1)
	}

	for i := 1; i < len(str2)+1; i++ {
		arr[0][i] = gap * float64(-i)
	}

	for i := 1; i < len(str1)+1; i++ {
		arr[i][0] = 0
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

			arr[r][c] = Max(up, left, diag)
		}
	}

	// find maximum value of last row
	maxVal := 0.0
	for _, val := range arr[len(str1)] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

type Number interface {
	~int | ~float64
}

func Max[T Number](nums ...T) T {
	if len(nums) == 0 {
		panic("Error: no inputs given to Max.")
	}
	m := nums[0]
	// nums gets converted to an array
	for _, val := range nums {
		if val > m {
			// update m
			m = val
		}
	}
	return m
}
