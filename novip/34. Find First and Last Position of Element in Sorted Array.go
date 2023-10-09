// 34. Find First and Last Position of Element in Sorted Array
package main

import "fmt"

type Case struct {
	nums   []int
	target int
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	i, j, m := 0, len(nums)-1, 0
	for i < j {
		m = (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m
		}
	}

	start, end := -1, -1
	if i < len(nums) && nums[i] == target {
		start = i
		j = len(nums) - 1
		for i < j {
			m = (i + j + 1) / 2
			if nums[m] > target {
				j = m - 1
			} else {
				i = m
			}
		}
		end = j
	}

	return []int{start, end}
}

func main() {
	cases := []Case{
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 6},
		{nums: []int{}, target: 0},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := searchRange(c.nums, c.target)
		fmt.Println("result:", res)
	}
}
