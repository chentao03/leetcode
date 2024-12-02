package main

import (
	"fmt"
	twosum "leetcode/solutions/0001_two_sum" // 使用你的模块名和包路径
)

func main() {
	nums := []int{2, 7, 2, 11, 15}
	target := 9
	fmt.Println(twosum.TwoSum(nums, target))
}
