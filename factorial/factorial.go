package main

import "fmt"

func main() {
	fmt.Println(factorial1(0)) //4*3*2*1 = 1
	fmt.Println(factorial1(1)) //4*3*2*1 = 1
	fmt.Println(factorial1(5)) //5*4*3*2*1 = 120
	fmt.Println(factorial2(0)) //4*3*2*1 = 1
	fmt.Println(factorial2(1)) //5*4*3*2*1 = 1
	fmt.Println(factorial2(5)) //5*4*3*2*1 = 120
}

func factorial1(n int) int {
	//5
	//4
	//3
	//2
	if n == 0 {
		return 1
	}

	//2
	if n < 3 {
		return n
	}

	//5 * 24
	//4 * 6
	//3 * 2
	return n * factorial1(n-1)
}

func factorial2(n int) int {
	result := 1

	for i := n; i > 1; i-- {
		result = result * i
	}

	return result
}

//big o -> O(n)
