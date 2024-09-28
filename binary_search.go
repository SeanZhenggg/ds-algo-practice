package main

import (
	"log"
)

func main() {

	arr := []int{1, 4, 7, 10, 13}
	search := 13
	idx, found := binarySearch(arr, search)
	if found {
		log.Printf("found the index %d containing search value: %d, the corresponding value is: %d", idx, search, arr[idx])
	} else {
		log.Printf("value containing the search value not found")
	}
}

func binarySearch(arr []int, search int) (int, bool) {
	var left, right, mid int

	if len(arr) == 0 {
		return -1, false
	}

	left = 0
	right = len(arr) - 1

	if len(arr) == 1 {
		if arr[0] != search {
			return -1, false
		} else {
			return 0, true
		}
	}

	for left <= right {
		mid = (left + right) / 2
		log.Printf("left: %d, right: %d, current mid: %d", left, right, mid)

		if search < arr[mid] {
			right = mid - 1
		} else if search == arr[mid] {
			return mid, true
		} else {
			left = mid + 1
		}
	}

	log.Printf("current mid: %d", mid)
	if arr[mid] < search {
		return mid, true
	} else {
		return -1, false
	}
}
