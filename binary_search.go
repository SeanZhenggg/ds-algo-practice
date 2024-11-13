package main

import (
	"log"
)

func main() {

	arr := []int{1, 4, 7, 10, 13}
	search := 13
	idx, found := binarySearch(arr, search)
	if found {
		var val int
		if idx != -1 {
			val = arr[idx]
		}
		log.Printf("search value: %d, index: %d, the corresponding value is: %d", search, idx, val)
	} else {
		log.Printf("value containing the search value not found")
	}
}

func binarySearch(arr []int, search int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	var left, right, mid int
	left = 0
	right = len(arr) - 1

	for left <= right {
		mid = (left + right) / 2
		if search < arr[mid] {
			right = mid - 1
		} else if search == arr[mid] {
			return mid, true
		} else {
			left = mid + 1
		}
	}

	return -1, false
}
