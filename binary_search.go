package main

import (
	"log"
)

func main() {

	arr := []int{1, 4, 7, 10, 13}
	idx, found := binarySearch(arr, 13)
	log.Printf("found: %t, at index: %d", found, idx)

	idx, found = binarySearchWithFindGreaterMinIndexWhenNotFound(arr, 8)
	log.Printf("found: %t, at index: %d", found, idx)

	idx, found = binarySearchWithFindLesserMaxIndexWhenNotFound(arr, 8)
	log.Printf("found: %t, at index: %d", found, idx)

	arr2 := []int{1, 4, 7, 7, 10, 13}
	idx, found = binarySearchWithDuplicateSearchValues(arr2, 7)
	log.Printf("found: %t, at index: %d", found, idx)

	idx, found = binarySearchWithDuplicateSearchValues2(arr2, 7)
	log.Printf("found: %t, at index: %d", found, idx)
}

func binarySearch(arr []int, search int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if search < arr[mid] {
			r = mid - 1
		} else if search == arr[mid] {
			return mid, true
		} else {
			l = mid + 1
		}
	}

	return -1, false
}

func binarySearchWithFindGreaterMinIndexWhenNotFound(arr []int, search int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if search < arr[mid] {
			r = mid - 1
		} else if search == arr[mid] {
			return mid, true
		} else {
			l = mid + 1
		}
	}

	return l, true
}

func binarySearchWithFindLesserMaxIndexWhenNotFound(arr []int, search int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if search < arr[mid] {
			r = mid - 1
		} else if search == arr[mid] {
			return mid, true
		} else {
			l = mid + 1
		}
	}

	return r, true
}

func binarySearchWithDuplicateSearchValues(arr []int, search int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if search <= arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l, true
}

func binarySearchWithDuplicateSearchValues2(arr []int, search int) (int, bool) {
	if len(arr) == 0 {
		return -1, false
	}

	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if search < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return r, true
}
