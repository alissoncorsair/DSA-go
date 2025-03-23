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
	arrA := arr[:len(arr)/2]
	arrB := arr[len(arr)/2:]
	arrA = mergeSort(arrA)
	arrB = mergeSort(arrB)

	return merge(arrA, arrB)
}

func merge(arrA, arrB []int) []int {
	c := make([]int, 0)

	for len(arrA) > 0 && len(arrB) > 0 {

		if arrA[0] > arrB[0] {
			c = append(c, arrB[0])
			arrB = arrB[1:]
		} else {
			c = append(c, arrA[0])
			arrA = arrA[1:]
		}
	}

	for len(arrA) > 0 {
		c = append(c, arrA[0])
		arrA = arrA[1:]
	}

	for len(arrB) > 0 {
		c = append(c, arrB[0])
		arrB = arrB[1:]
	}

	return c
}
