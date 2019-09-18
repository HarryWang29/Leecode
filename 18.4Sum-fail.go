package main

import "fmt"

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func combine(arr []int, start int, result *[]int, index, n int, ret *[][]int) {
	for ct := start; ct < len(arr)-index+1; ct++ {
		(*result)[index-1] = ct
		if index-1 == 0 {
			r := make([]int, 0)
			for j := n - 1; j >= 0; j-- {
				r = append(r, arr[(*result)[j]])
			}
			*ret = append(*ret, r)
		} else {
			combine(arr, ct+1, result, index-1, n, ret)
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	result := make([]int, 4)
	ret := make([][]int, 0)
	ret1 := make([][]int, 0)
	n := 4
	index := 4
	ct := 0
	combine(nums, ct, &result, index, n, &ret)
	fmt.Println(ret)
	for _, value := range ret {
		r := 0
		for _, vv := range value {
			r += vv
		}
		if r == target {
			ret1 = append(ret1, value)
		}
	}
	return ret1
}

func main() {
	//fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
