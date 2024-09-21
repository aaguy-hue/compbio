package main

import "fmt"

type Alignment [2]string

type Number interface {
	~int | ~float64
}

func InitializeMatrix[ElementType any](rows, cols int) [][]ElementType {
	arr := make([][]ElementType, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]ElementType, cols)
	}
	return arr
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

func Min[T Number](nums ...T) T {
	if len(nums) == 0 {
		panic("Error: no inputs given to Min.")
	}
	m := nums[0]
	// nums gets converted to an array
	for _, val := range nums {
		if val < m {
			// update m
			m = val
		}
	}
	return m
}

func MaximumElementIndex(mtx [][]float64) (int, int) {
	max := mtx[0][0]
	row := 0
	col := 0

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if mtx[i][j] > max {
				max = mtx[i][j]
				row = i
				col = j
			}
		}
	}

	return row, col
}

func MaximumElement(mtx [][]int) int {
	max := mtx[0][0]

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if mtx[i][j] > max {
				max = mtx[i][j]
			}
		}
	}

	return max
}

func PrintMatrix[T any](mtx [][]T) {
	for _, row := range mtx {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Print("\n")
	}
}

func AddBools(theBools ...bool) int {
	sum := 0
	for _, thing := range theBools {
		if thing {
			sum++
		}
	}
	return sum
}
