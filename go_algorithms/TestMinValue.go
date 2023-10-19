package main

import "fmt"

func min(arrays []int, length int) int {
	var minIndex = 0
	for i := 1; i < length; i++ {
		if arrays[minIndex] > arrays[i] {
			minIndex = i
		}
	}
	var minValue = arrays[minIndex]
	return minValue
}

func main() {
	var scores = []int{60, 80, 1, 95, 50, 70}
	var length = len(scores)
	var minValue = min(scores, length)
	fmt.Printf("Min Value = %d\n", minValue)
}
