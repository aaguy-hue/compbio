package main

// LocalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score local alignment of the strings corresponding to these penalties.
func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	var a Alignment
	mtx := LocalScoreTable(str1, str2, match, mismatch, gap)

	maxrow, maxcol := MaximumElementIndex(mtx)
	r, c := maxrow, maxcol

	align1 := ""
	align2 := ""

	for r > 0 || c > 0 {
		if mtx[r][c] == 0 { // if we reached the start of the "free driving"
			break
		}

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

	start1 := r
	start2 := c
	end1 := maxrow
	end2 := maxcol

	a[0] = align1
	a[1] = align2
	return a, start1, end1, start2, end2
}

// func main() {
// 	table := LocalScoreTable("GAGA", "GAT", 1, 1, 2)
// 	fmt.Println(table)
// 	fmt.Println(MaximumIndex(table))
// 	fmt.Println(LocalAlignment("GAGA", "GAT", 1, 1, 2))
// }
