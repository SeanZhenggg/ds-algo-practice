package main

import "log"

func main() {
	arr := findFirstSmallerValueBeforeCurrent([]int{9, 2, 4, 5, 7, 3})
	log.Println(arr)
	arr2 := findMaximumValueBeforeCurrent([]int{9, 2, 4, 5, 7, 3})
	log.Println(arr2)
	arr3 := findMinimumValueBeforeCurrent([]int{9, 2, 4, 5, 7, 3})
	log.Println(arr3)
}

func findFirstSmallerValueBeforeCurrent(nums []int) []int {
	var stack = make([]int, 0, len(nums))
	var ret = make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ret = append(ret, stack[len(stack)-1])
		} else {
			ret = append(ret, -1)
		}
		stack = append(stack, nums[i])
	}
	return ret
}

func findMaximumValueBeforeCurrent(nums []int) []int {
	var stack = make([]int, 0, len(nums))
	var ret = make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		}
		ret = append(ret, stack[len(stack)-1])
	}
	return ret
}

func findMinimumValueBeforeCurrent(nums []int) []int {
	var stack = make([]int, 0, len(nums))
	var ret = make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		}
		ret = append(ret, stack[len(stack)-1])
	}
	return ret
}
