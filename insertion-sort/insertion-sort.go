package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{-6, 20, 8, -2, 4}))
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}

	return arr
}

//big(o) = o(n^2)
