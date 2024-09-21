package main

func AffineAlignment(match, mismatch, gapOpen, gapExtension int, str1, str2 string) (maxScore int, align1 string, align2 string) {
	lower_mtx, middle_mtx, upper_mtx := AffineScoreTable(str1, str2, match, mismatch, gapOpen, gapExtension)
	r := len(str1)
	c := len(str2)

	align1 = ""
	align2 = ""

	lastAction := ""
	for r > 0 || c > 0 { // while not at the top left corner
		var goUp bool = (r > 0 && lower_mtx[r][c] == middle_mtx[r][c]) || (r > 0 && c == 0)                                                // top, with an extra case to check if at leftmost column
		var goLeft bool = (c > 0 && upper_mtx[r][c] == middle_mtx[r][c]) || (r == 0 && c > 0)                                              // left, with an extra case to catch if at topmost row
		var goDiagonal bool = (r > 0 && c > 0) && middle_mtx[r][c] == middle_mtx[r-1][c-1]+ComputeScore(str1, str2, r, c, match, mismatch) // diagonal

		counter := AddBools(goUp, goLeft, goDiagonal)
		var action string
		if counter == 1 {
			if goUp {
				action = "up"
			} else if goLeft {
				action = "left"
			} else if goDiagonal {
				action = "diag"
			}
		} else if counter == 2 {
			if goUp && goLeft {
				if lastAction != "diag" && lastAction != "" {
					action = lastAction
				} else {
					action = "up" // arbitrarily chosen
				}
			} else if goUp && goDiagonal {
				if lastAction != "left" && lastAction != "" {
					action = lastAction
				} else {
					action = "diag" // arbitrary
				}
			} else if goLeft && goDiagonal {
				if lastAction != "up" && lastAction != "" {
					action = lastAction
				} else {
					action = "left" // arbitrary
				}
			}
		} else if counter == 3 {
			if lastAction != "" {
				action = lastAction
			} else {
				action = "diag" // arbitrary
			}
		} else {
			panic("BAD MATRIX")
		}

		switch action {
		case "up":
			lastAction = "up"
			align1 = string(str1[r-1]) + align1
			align2 = "-" + align2
			r--
		case "diag":
			lastAction = "diag"
			align1 = string(str1[r-1]) + align1
			align2 = string(str2[c-1]) + align2
			r--
			c--
		case "left":
			lastAction = "left"
			align1 = "-" + align1
			align2 = string(str2[c-1]) + align2
			c--
		default:
			panic("AffineAlignment: Bad scoring matrix")
		}
	}

	// maxScore = MaximumElement(middle_mtx)
	maxScore = middle_mtx[len(str1)][len(str2)]
	return
}
