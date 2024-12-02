package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	funcs := []func([]int, int) []int{TwoSum1, TwoSum2, TwoSum}

	for _, testfunc := range funcs {
		if res := testfunc(nums, 9); !reflect.DeepEqual(res, []int{0, 1}) {
			t.Error("Failed,tow sum")
		}
		if res := testfunc(nums, 6); !reflect.DeepEqual(res, []int{}) {
			t.Error("Failed,tow sum")
		}
	}
}
