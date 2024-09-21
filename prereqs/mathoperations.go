package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

func Factorial(n int) int {
    ans := 1
    i := 1
    
    for i <= n {
        ans *= i
        i++
    }
    
    return ans
}

func AnotherFactorial(n int) int {
    ans := 1
    for i := 1; i <= n; i++ {
        ans *= i
    }
    
    return ans;
}

func FactorialArray(n int) []int {
    
    arr := make([]int, n+1)
    
    arr[0] = 1
    for i:=1; i<=n; i++ {
        arr[i] = arr[i-1] * i
    }
    
    return arr
}

func TrivialGCD(a, b int) int {
    gcd := 1
    min := Min2(a, b)
    
    for i := 1; i<= min; i++ {
        if a % i == 0 && b % i == 0 {
            gcd = i
        }
    }
    
    return gcd
}

func EuclidGCD(a, b int) int {
    for a != b {
        if a > b {
            a = a - b
        } else {
            b = b - a
        }
    }
    return a
}

func GCDArray(a []int) int {
    
    gcd := 1
    min := MinIntegerArray(a)
    
    for i := 1; i <= min; i++ {
        if DividesAll(a, i) {
            gcd = i
        }
    }
    
    return gcd
}

func SumIntegers(numbers ...int) int {
    var ans int = 0
    for _, num := range numbers {
        ans += num
    }
    return ans
}

func SumFirstNIntegers(n int) int {
    i := 1
    ans := 0
    
    for i <= n {
        ans += i
        i++
    }
    
    return ans
}

func SumEven(k int) int {
    ans := 0
    for i := 0; i <= k; i = i+2 {
        ans += i
    }
    return ans
}

func TrivialPrimeFinder(n int) []bool {
    isPrimeArr := make([]bool, n+1)
    
    for p := 2; p <= n; p++ {
        isPrimeArr[p] = IsPrime(p)
    }
    
    return isPrimeArr
}

func CrossOffMultiples(primeBooleans []bool, p int) []bool {
    for k := 2*p; k < len(primeBooleans); k+=p {
        primeBooleans[k] = false
    }
    return primeBooleans
}

func SieveOfEratosthenes(n int) []bool {
    primeBools := make([]bool, n+1)
    
    for p := 2; p <= n; p++ {
        primeBools[p] = true
    }
    
    for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
        if IsPrime(p) {
            primeBools = CrossOffMultiples(primeBools, p)
        }
    }
    
    return primeBools
}

func ListPrimes(n int) []int {
    numPrimes := 0
    sieve := SieveOfEratosthenes(n);
    for _, val := range sieve {
        if val {
            numPrimes++
        }
    }
    
    primeList := make([]int, numPrimes);
    lastAppend := 0;
    for i, val := range sieve {
        if val {
            primeList[lastAppend] = i
            lastAppend++
        }
    }
    
    return primeList
}

func Permutation(n, k int) int {
    ans := 1
    
    for i := n; i > (n-k); i-- {
        ans *= i
    }
    return ans
}

func Combination(n, k int) int {
    ans := 1
    
    for i := n; i > (n-k); i-- {
        ans *= i
    }
    
    kfac := 1
    for i := 1; i <= k; i++ {
        kfac *= i
    }
    
    return ans/kfac
}

func Power(a, b int) int {
    ans := 1
    for i := 1; i <= b; i++ {
        ans *= a
    }
    return ans
}

func SumProperDivisors(n int) int {
    ans := 0
    for i := 1; i < n; i++ {
        if n%i == 0 {
            ans += i
        }
    }
    return ans
}