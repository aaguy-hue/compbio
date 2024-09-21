package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
    "math"
)

func Min2(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func MinIntegerArray(list []int) int {
    min := list[0]
    for i:=0;i<len(list);i++ {
        if list[i] < min {
            min = list[i]
        }
    }
    return min
}

func MinIntegers(numbers ...int) int {
    min := numbers[0]
    for _, val := range numbers {
        if val < min {
            min = val
        }
    }
    return min

}

func MaxIntegerArray(list []int) int {
    var ans int = list[0]
    for _, val := range list[1:] {
        if val > ans {
            ans = val
        }
    }
    return ans
}

func MaxIntegers(numbers ...int) int {
    var ans int = numbers[0]
    for _, val := range numbers[1:] {
        if val > ans {
            ans = val
        }
    }
    return ans
}

func WhichIsGreater(x, y int) int {
    if x == y {
        return 0
    } else if x > y {
        return 1
    } else {
        return -1
    }
}

func SameSign(x, y int) bool {
    if x >= 0 && y >= 0 || x <= 0 && y <= 0 {
        return true
    } else {
        return false
    }
}

func PositiveDifference(a, b int) int {
    diff := a-b
    if diff > 0 {
        return diff
    } else {
        return -diff
    }
}

func IsPrime(p int) bool {
    if p == 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(p))); i++ {
        if p % i == 0 {
            return false
        }
    }
    return true
}

func DividesAll(a []int, d int) bool {
    if d == 0 {
        return false
    }
    
    for _, val := range a {
        if val % d != 0 {
            return false
        }
    }
    return true
}