package main

import (
	"fmt"
	"os"
	"strconv"
)

func sumPrimesUpTo(n int) int {
	if n < 2 {
		return 0
	}

	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			sum += i
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return sum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(sumPrimesUpTo(n))
}
