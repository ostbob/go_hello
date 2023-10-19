package main

import "fmt"

func main() {
	var scores = []int{90, 200, 70, 50, 80, 60, 85, 120}
	var length = len(scores)

	bubble_sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func bubble_sort(arrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arrays[j] > arrays[j+1] {
				var temp = arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = temp
			}
		}
	}
}
