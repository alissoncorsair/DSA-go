package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{8, 20, -2, 4, -6}))
}

func bubbleSort(arr []int) []int {
	swap := true
	ordered := arr
	for swap {
		swapped := false
		for i := range len(ordered) - 1 {
			if ordered[i] > ordered[i+1] {
				ordered[i], ordered[i+1] = ordered[i+1], ordered[i]
				swapped = true
			}
		}

		if !swapped {
			swap = false
		}
	}
	return ordered
}

//o(n2)
