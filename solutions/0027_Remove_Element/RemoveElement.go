package main

/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The order of the elements may be changed. Then return the number of elements in nums which are
not equal to val.
Consider the number of elements in nums which are not equal to val be k, to get accepted,
you need to do the following things:
Change the array nums such that the first k elements of nums contain the elements
which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.

*/

import "fmt"

// 暴力循环求解o(n^2),o(1)
func removeElement1(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if val == nums[i] {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i-- //移位一次，nums[i]覆盖成新的元素，需重新与val比较
			size--
		}
	}
	return size
}

// 快慢指针，fast指针一轮遍历，slow指针只有在val!=nums[fast]时移动
// o(n),o(1)
func removeElement(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if val != nums[fastIndex] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++ //不等于val时 快慢指针同步，否则slow指针不移位
		}
	}
	return slowIndex
}

// 技巧 o(n),o(1)
func removeElement2(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if val == nums[i] {
			nums[i] = nums[size-1]
			i-- //移位后重新检测i位置
			size--
		}
	}
	return size
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
