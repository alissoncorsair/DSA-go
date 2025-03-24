package main

import "fmt"

func main() {
	arr := []int{2, 8, 5, 3, 9, 4, 1, 7}
	fmt.Println("unsorted", arr)
	arr = mergeSort(arr)
	fmt.Println("sorted", arr)

}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	arrA := mergeSort(arr[:middle])
	arrB := mergeSort(arr[middle:])

	return merge(arrA, arrB)
}

func merge(leftArr, rightArr []int) []int {
	sorted := make([]int, 0)

	for len(leftArr) > 0 && len(rightArr) > 0 {
		if leftArr[0] > rightArr[0] {
			sorted = append(sorted, rightArr[0])
			rightArr = rightArr[1:]
		} else {
			sorted = append(sorted, leftArr[0])
			leftArr = leftArr[1:]
		}
	}

	for len(leftArr) > 0 {
		sorted = append(sorted, leftArr[0])
		leftArr = leftArr[1:]
	}

	for len(rightArr) > 0 {
		sorted = append(sorted, rightArr[0])
		rightArr = rightArr[1:]
	}

	return sorted
}
