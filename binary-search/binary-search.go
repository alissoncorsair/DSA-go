package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{-5, 2, 4, 6, 10}, 10))
	fmt.Println(binarySearch([]int{-5, 2, 4, 6, 10}, 6))
	fmt.Println(binarySearch([]int{-5, 2, 4, 6, 10}, 4))
	fmt.Println(binarySearch([]int{-5, 2, 4, 6, 10}, 20))
	fmt.Println("recursive")
	fmt.Println(recursiveBinarySearch([]int{-5, 2, 4, 6, 10}, 10))
	fmt.Println(recursiveBinarySearch([]int{-5, 2, 4, 6, 10}, 6))
	fmt.Println(recursiveBinarySearch([]int{-5, 2, 4, 6, 10}, 4))
	fmt.Println(recursiveBinarySearch([]int{-5, 2, 4, 6, 10}, 20))

}

func binarySearch(arr []int, n int) int {
	leftPosition := 0
	rightPosition := len(arr) - 1

	for leftPosition <= rightPosition {
		middlePosition := (leftPosition + rightPosition) / 2
		middleElement := arr[middlePosition]
		if middleElement == n {
			return middlePosition
		}

		if middleElement > n {
			rightPosition = middlePosition - 1
		}

		if middleElement < n {
			leftPosition = middlePosition + 1
		}
	}
	return -1
}

func recursiveBinarySearch(arr []int, n int) int {
	if len(arr) == 0 {
		return -1
	}

	leftPosition := 0
	rightPosition := len(arr) - 1
	middlePosition := (leftPosition + rightPosition) / 2

	if arr[middlePosition] == n {
		return middlePosition
	}

	if arr[middlePosition] > n {
		newArray := arr[:middlePosition]
		result := recursiveBinarySearch(newArray, n)
		if result == -1 {
			return -1
		}

		return middlePosition - 1 + result
	}

	if arr[middlePosition] < n {
		newArray := arr[middlePosition+1:]
		result := recursiveBinarySearch(newArray, n)

		if result == -1 {
			return -1
		}

		return middlePosition + 1 + result
	}

	return -1
}
