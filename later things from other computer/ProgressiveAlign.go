package main

func ProgressiveAlign(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) Alignment {
	var a Alignment = make(Alignment, len(align1)+len(align2))

	if len(align1) == 1 && len(align2) == 1 {
		supergap = gap
	}

	mtx := ProgressiveAlignmentScoreTable(align1, align2, match, mismatch, gap, supergap)
	numRows := len(mtx)
	numCols := len(mtx[0])

	row := numRows - 1
	col := numCols - 1
	for row != 0 || col != 0 {
		if col > 0 && mtx[row][col] == mtx[row][col-1]-supergap { // left
			a = AddInsertion(a, align1, align2, row, col-1)
			col--
		} else if row > 0 && mtx[row][col] == mtx[row-1][col]-supergap { // up
			a = AddDeletion(a, align1, align2, row-1, col)
			row--
		} else if row > 0 && col > 0 { // diag
			a = AddMatch(a, align1, align2, row-1, col-1)
			row--
			col--
		}
	}

	return a
}

func AddInsertion(a, align1, align2 Alignment, row, col int) Alignment {
	for i := 0; i < len(align1); i++ {
		a[i] = "-" + a[i]
	}

	for j := len(align1); j < len(align1)+len(align2); j++ {
		a[j] = string(align2[j-len(align1)][col]) + a[j]
	}

	return a
}

func AddDeletion(a, align1, align2 Alignment, row, col int) Alignment {
	for i := 0; i < len(align1); i++ {
		a[i] = string(align1[i][row]) + a[i]
	}

	for j := len(align1); j < len(align1)+len(align2); j++ {
		a[j] = "-" + a[j]
	}

	return a
}

func AddMatch(a, align1, align2 Alignment, row, col int) Alignment {
	for i := 0; i < len(align1); i++ {
		a[i] = string(align1[i][row]) + a[i]
	}

	for j := len(align1); j < len(align1)+len(align2); j++ {
		a[j] = string(align2[j-len(align1)][col]) + a[j]
	}

	return a
}

func ProgressiveAlignmentScoreTable(align1 Alignment, align2 Alignment,
	match float64, mismatch float64, gap float64, supergap float64) [][]float64 {
	numRows := len(align1[0]) + 1
	numCols := len(align2[0]) + 1
	mtx := make([][]float64, numRows)
	for i := range mtx {
		mtx[i] = make([]float64, numCols)
	}

	for i := 1; i < numRows; i++ {
		mtx[i][0] = -(supergap * float64(i))
	}

	for i := 1; i < numCols; i++ {
		mtx[0][i] = -(supergap * float64(i))
	}

	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			up := mtx[i-1][j] - supergap
			left := mtx[i][j-1] - supergap
			diag := mtx[i-1][j-1] + SumPairsScore(align1, align2, i-1, j-1, match, mismatch, gap)
			mtx[i][j] = MaxFloat(up, left, diag)
		}
	}

	return mtx
}
