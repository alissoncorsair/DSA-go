package main

import "fmt"

func main() {
	fmt.Println(linearSearch([]int{-5, 2, 10, 4, 6}, 10))
}

func linearSearch(arr []int, n int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}

//big o -> o(n)
