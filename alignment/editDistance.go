package main

// EditDistance takes two strings as input. It returns the Levenshtein distance
// between the two strings; that is, the minimum number of substitutions, insertions, and deletions
// needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		panic("Error: empty string passed into function EditDistance")
	}

	// initialize the matrix
	arr := make([][]int, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		arr[i] = make([]int, len(str2)+1)
	}

	for i := 0; i < len(str1)+1; i++ {
		arr[i][0] = i
	}

	for i := 0; i < len(str2)+1; i++ {
		arr[0][i] = i
	}

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				arr[i][j] = Min(arr[i][j-1]+1, arr[i-1][j]+1, arr[i-1][j-1])
			} else {
				arr[i][j] = Min(arr[i][j-1]+1, arr[i-1][j]+1, arr[i-1][j-1]+1)
			}
		}
	}

	//fmt.Println(arr)

	return arr[len(str1)][len(str2)]
}

// // EditMatrix takes two strings as input. It returns a matrix of values corresponding
// // to the dynamic programming table of the two strings for the edit distance.
// func EditMatrix(str1, str2 string) [][]int {
// 	if len(str1) == 0 || len(str2) == 0 {
// 		panic("Blah")
// 	}
// 	numRows := len(str1) + 1
// 	numCols := len(str2) + 1

// 	//declare scoring matrix
// 	scoringMatrix := make([][]int, numRows)
// 	// make the rows each have length # columns
// 	for i := range scoringMatrix {
// 		scoringMatrix[i] = make([]int, numCols)
// 	}

// 	// we know what the distance is in the 0-th row and column -- the number of symbols encountered thus far.
// 	for j := range scoringMatrix[0] {
// 		scoringMatrix[0][j] = j
// 	}
// 	for i := range scoringMatrix {
// 		scoringMatrix[i][0] = i
// 	}

// 	// now we fill the interior of the matrix based on our recurrence relation
// 	for row := 1; row < numRows; row++ {
// 		for col := 1; col < numCols; col++ {
// 			// apply recurrence relation
// 			up := scoringMatrix[row-1][col] + 1
// 			left := scoringMatrix[row][col-1] + 1
// 			diag := scoringMatrix[row-1][col-1]
// 			// if we see a mismatch, update diag by adding 1.
// 			if str1[row-1] != str2[col-1] {
// 				diag++
// 			}
// 			scoringMatrix[row][col] = Min(up, left, diag)
// 		}
// 	}

// 	return scoringMatrix
// }

// func main() {
// 	var str1 string
// 	var str2 string
// 	fmt.Scanln(&str1)
// 	fmt.Scanln(&str2)
// 	fmt.Println(EditDistance(str1, str2))
// }
