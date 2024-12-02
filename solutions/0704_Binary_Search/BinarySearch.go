package main

import (
	"fmt"
)

// brust cycle
func search1(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}
	return -1
}

// binary search，因为要输出原数据下标，不适合用递归
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start <= end {
		//mid := start + (end-start)/2
		mid := start + (end-start)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

/*
func binarysearch(nums []int, target int) int {
	mid := int(len(nums))
	if nums[mid] == target {
		return target
	}
	if nums[mid] > target {
		binarysearch(nums[0:mid], target)
	} else {
		binarysearch(nums[mid:len(nums)-1], target)
	}
}
*/

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))

}
