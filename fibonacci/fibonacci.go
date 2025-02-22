package main

import "fmt"

func main() {
	fmt.Println(fibonacci(2)) //0, 1, 1, 2
	fmt.Println(fibonacci(5))
	fmt.Println(recursiveFibonacci(2)) //0, 1
	fmt.Println(recursiveFibonacci(5))
}

func fibonacci(n int) []int {

	if n <= 0 {
		return []int{}
	}

	fibo := make([]int, n)
	fibo[0] = 0

	if n > 1 {
		fibo[1] = 1
	}

	for i := 2; i < n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}

	return fibo
}

//big o -> o(n)

func recursiveFibonacci(n int) int {
	//n = 7
	//6
	//5
	//4
	//3
	//2
	//1
	if n < 2 {
		return n
	}

	//fn(2-1)[1] + fn0[0]
	//1 + 1
	//2 + 1
	//3 + 2
	//5 + 3
	//8 + 5
	//13
	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}
