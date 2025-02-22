package main

import "fmt"

func main() {
	fmt.Println(powerOfTwo(1))
	fmt.Println(powerOfTwo(2))
	fmt.Println(powerOfTwo(8))
}

func powerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	result := n
	for result > 1 {
		fmt.Println("result", result)
		remaining := result % 2

		if remaining != 0 {
			return false
		}

		result = result / 2

		if result == 1 {
			return true
		}

	}
	return true
}

// big o => o(logn)
