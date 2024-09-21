package main

import (
	"math"
)

// LCSLength takes two strings as input. It returns the length of a longest common
// subsequence of the two strings.
func LCSLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 2 {
		panic("Error: empty string passed into function LCSLength")
	}

	// initialize the matrix
	arr := make([][]int, len(str1)+1)
	for i := 0; i < len(str1)+1; i++ {
		arr[i] = make([]int, len(str2)+1)
	}

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				arr[i][j] = Max3(arr[i][j-1], arr[i-1][j], arr[i-1][j-1]+1)
			} else {
				arr[i][j] = Max3(arr[i][j-1], arr[i-1][j], arr[i-1][j-1])
			}
		}
	}

	//fmt.Println(arr)

	return arr[len(str1)][len(str2)]
}

func Max3(a, b, c int) int {
	return int(math.Max(float64(a), math.Max(float64(b), float64(c))))
}

// func main() {
// 	var str1 string
// 	var str2 string
// 	fmt.Scanln(&str1)
// 	fmt.Scanln(&str2)
// 	fmt.Println(LCSLength(str1, str2))
// }
