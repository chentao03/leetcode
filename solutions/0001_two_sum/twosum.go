package twosum

//

/*
brute force v1
time complexity o(n^2)
space complexity o(1)
*/
func TwoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			//if i == 1 {
			//	fmt.Println(i, j)
			//}
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// brute force v2
func TwoSum2(nums []int, target int) []int {
	for index, value := range nums {
		for k := index + 1; k < len(nums); k++ {
			if nums[k]+value == target && k != index {
				return []int{index, k}
			}
		}
	}
	return []int{}
}

/*
map func
time complexity 0(n)
space comlexity O(n)
*/
func TwoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for index := 0; index < len(nums); index++ {
		if res, ok := record[target-nums[index]]; ok {
			return []int{res, index}
		} else {
			record[nums[index]] = index
		}
	}
	return []int{}
}
