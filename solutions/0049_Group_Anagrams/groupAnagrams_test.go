package GroupAnagrams

//package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	funcs := []func([]string) [][]string{GroupAnagrams, GroupAnagrams1}
	solutions := [][]string{{"bat"}, {"eat", "tea", "ate"}, {"tan", "nat"}} // 修正为正确的二维切片语法

	for _, testfunc := range funcs {
		res := testfunc(strs)                   // 修正参数，移除 solutions
		if !reflect.DeepEqual(res, solutions) { // 修正比较的预期结果
			t.Error("Failed, group anagrams")
		}
	}
}
