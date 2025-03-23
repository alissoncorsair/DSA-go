package main

import "fmt"

func main() {
	arr := []int{1, 22, -2, -3, 4, 2, 1}

	fmt.Println("unsorted", arr)

	quicksort(arr, 0, len(arr)-1)

	fmt.Println("sorted", arr)
}

func quicksort(arr []int, left, right int) {
	if left < right {
		pivot := partition(arr, left, right)

		quicksort(arr, left, pivot-1)
		quicksort(arr, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[right], arr[i+1] = arr[i+1], arr[right]

	return i + 1
}
