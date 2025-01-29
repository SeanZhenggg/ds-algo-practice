package main

import "log"

func MergeSort(nums []int, l int, r int) []int {
	// 拆分左右兩邊，繼續遞迴處理左半邊跟右半邊

	if l == r {
		return []int{nums[l]}
	}

	mid := l + (r-l)/2

	left := MergeSort(nums, l, mid)
	right := MergeSort(nums, mid+1, r)

	merge(nums, left, right, l)
	arr := make([]int, r-l+1)
	copy(arr, nums[l:r+1])
	return arr
}

func merge(nums []int, left []int, right []int, l int) {
	var ll = 0
	var rr = 0

	for ll < len(left) && rr < len(right) {
		if left[ll] <= right[rr] {
			nums[l] = left[ll]
			ll++
		} else {
			nums[l] = right[rr]
			rr++
		}
		l++
	}

	if ll < len(left) {
		for i := ll; i < len(left); i++ {
			nums[l] = left[i]
			l++
		}
	}
	if rr < len(right) {
		for i := rr; i < len(right); i++ {
			nums[l] = right[i]
			l++
		}
	}
}

func main() {
	nums := []int{3, 1, 5, 2, 4, 11, 6, 5, 7, 4}
	MergeSort(nums, 0, len(nums)-1)
	log.Printf("nums: %v", nums)
}
