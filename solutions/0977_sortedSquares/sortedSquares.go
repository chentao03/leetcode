package main

import (
	"fmt"
	"sort"
)

/*

 */

// 暴力求解o(n+nlogn)
func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// 技巧，原始数据有序，平方后一定是两端大中间小，则首尾双指针移位找大的数
func sortedSquares(nums []int) []int {
	k := len(nums) - 1
	result := make([]int, len(nums)) //创建切片而不是数组
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			result[k] = nums[j] * nums[j]
			j--
		} else {
			result[k] = nums[i] * nums[i]
			i++
		}
		k--
	}
	return result
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
