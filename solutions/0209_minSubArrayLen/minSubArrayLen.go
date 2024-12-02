package main

import "fmt"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。找出该数组中满足其总和大于等于target的长度最小的
子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/

// 暴力求解答，时间超限
func minSubArrayLen1(target int, nums []int) int {
	result := 1<<31 - 1
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				sublength := j - i + 1
				if result > sublength {
					result = sublength
					break
				}
			}
		}
	}
	if result == 1<<31-1 {
		result = 0
	}
	return result
}

// 双指针滑动窗口，降低一次循环. 窗口尾部j一轮遍历，i根据target比较情况移动
// o(n)
func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	sum := 0
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		//for循环
		for sum >= target {
			sublength := j - i + 1
			if result > sublength {
				result = sublength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == len(nums)+1 {
		result = 0
	}
	return result
}

func main() {
	nums := []int{1, 4, 4}
	target := 4
	fmt.Println(minSubArrayLen(target, nums))

}
