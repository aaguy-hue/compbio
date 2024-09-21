package main

import (
	"math"
)

func CountNumDigits(x int) int {
	if x < 0 {
		x = -x
	}

	digits := 1
	for {
		if x < 10 {
			return digits
		}
		x = x / 10
		digits++
	}
}

func SquareMiddle(x, numDigits int) int {
	if (numDigits%2 != 0) || (numDigits <= 0) || (x < 0) || (CountNumDigits(x) > numDigits) {
		return -1
	}

	x2 := x * x
	x2 /= int(math.Pow10(numDigits / 2))
	x2 %= int(math.Pow10(numDigits))
	return x2
}

func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	var retval []int
	retval = append(retval, seed)

	for !HasRepeat(retval) {
		seed = SquareMiddle(seed, numDigits)
		retval = append(retval, seed)
	}
	return retval
}

func ComputePeriodLength(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	for i, a := range arr[0 : len(arr)-1] {
		for j, b := range arr[i+1:] {
			if a == b {
				return j + 1
			}
		}
	}

	return 0
}

// func main() {
// 	count := 0
// 	for seed := 1; seed <= 9999; seed++ {
// 		seq := GenerateMiddleSquareSequence(seed, 4)
// 		if ComputePeriodLength(seq) < 10 {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }
