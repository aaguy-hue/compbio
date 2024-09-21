package main

// LongestCommonSubsequence takes two strings as input.
// It returns a longest common subsequence of the two strings.
func LongestCommonSubsequence(str1, str2 string) string {
	finalString := ""

	mtx := LCSScoreMatrix(str1, str2)
	lcsLength := mtx[len(str1)][len(str2)]

	row := len(str1)
	col := len(str2)

	// either you loop until you reach the top left, until you reach the first row/col, the length of the string is the expected length of the lcs (the bottom right corner), or you reach a value of 0 (no more matches)
	for len(finalString) < lcsLength {
		values := [3][2]int{{row, col - 1}, {row - 1, col - 1}, {row - 1, col}}

		for i := 0; i < 3; i++ {
			if inRange(values[i][0], values[i][1], mtx) {
				prow := values[i][0]
				pcol := values[i][1]

				//DIAGONAL
				if i == 1 {
					isMatch := match(str1, str2, row, col)
					if mtx[prow][pcol]+isMatch == mtx[row][col] {
						row = prow
						col = pcol
						if isMatch == 1 {
							finalString = string(str1[prow]) + finalString
						}

						break
					}
				}

				//NOT DIAGONAL
				if i != 1 && mtx[prow][pcol] == mtx[row][col] {
					row = prow
					col = pcol
					break
				}
			}
		}
	}

	return finalString
}

func match(str1, str2 string, val1, val2 int) int {
	if str1[val1-1] == str2[val2-1] {
		return 1
	}
	return 0
}

func inRange(row, col int, mtx [][]int) bool {
	if 0 <= row && row < len(mtx) && 0 <= col && col < len(mtx[0]) {
		return true
	}
	return false
}

func LCSScoreMatrix(str1, str2 string) [][]int {
	arr := make([][]int, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		arr[i] = make([]int, len(str2)+1)
	}

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				arr[i][j] = Max(arr[i][j-1], arr[i-1][j], arr[i-1][j-1]+1)
			} else {
				arr[i][j] = Max(arr[i][j-1], arr[i-1][j], arr[i-1][j-1])
			}
		}
	}
	return arr
}

// func main() {
// 	var str1 string
// 	var str2 string
// 	fmt.Scanln(&str1)
// 	fmt.Scanln(&str2)
// 	fmt.Println(LongestCommonSubsequence(str1, str2))
// }
