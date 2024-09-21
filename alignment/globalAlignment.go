package main

// GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var a Alignment
	mtx := GlobalScoreTable(str1, str2, match, mismatch, gap)

	r := len(str1)
	c := len(str2)

	align1 := ""
	align2 := ""

	for r > 0 || c > 0 { // while not at the top left corner
		if r > 0 && mtx[r-1][c]-gap == mtx[r][c] { // top
			align1 = string(str1[r-1]) + align1
			align2 = "-" + align2
			r--
		} else if c > 0 && mtx[r][c-1]-gap == mtx[r][c] { // left
			align1 = "-" + align1
			align2 = string(str2[c-1]) + align2
			c--
		} else if r > 0 && c > 0 { // diagonal
			align1 = string(str1[r-1]) + align1
			align2 = string(str2[c-1]) + align2

			if str1[r-1] == str2[c-1] && mtx[r-1][c-1]+match == mtx[r][c] {
				r--
				c--
			} else if str1[r-1] != str2[c-1] && mtx[r-1][c-1]-mismatch == mtx[r][c] {
				r--
				c--
			}
		}
	}

	a[0] = align1
	a[1] = align2
	return a
}
