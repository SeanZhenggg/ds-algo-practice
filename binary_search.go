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

	arr3 := []int{1, 4, 7, 7, 10, 13}
	idx, found = binarySearchWithDuplicateSearchValues(arr3, 14)
	log.Printf("found: %t, at index: %d", found, idx)

	idx, found = binarySearchWithDuplicateSearchValues2(arr3, 14)
	log.Printf("found: %t, at index: %d", found, idx)

	arr4 := []int{1, 4, 7, 7, 10, 13}
	idx, found = binarySearchWithDuplicateSearchValues(arr4, 0)
	log.Printf("found: %t, at index: %d", found, idx)

	idx, found = binarySearchWithDuplicateSearchValues2(arr4, 0)
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

	return l, false
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

	return r, false
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
	if l == -1 || l == len(arr) {
		return l, false
	}

	return l, arr[l] == search
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

	if r == -1 || r == len(arr) {
		return r, false
	}

	return r, arr[r] == search
}
