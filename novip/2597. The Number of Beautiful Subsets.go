// 2597. The Number of Beautiful Subsets

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(ay []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, n := range ay {
		p.Next = &ListNode{Val: n}
		p = p.Next
	}
	return head.Next
}

func PrintList(l *ListNode) {
	for p := l; p != nil; p = p.Next {
		fmt.Print(p.Val, ", ")
	}
	fmt.Println("")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func genVec[T any](l int, d ...T) []T {
	ret := make([]T, l)
	if len(d) > 0 {
		for i := 0; i < l; i++ {
			ret[i] = d[0]
		}
	}
	return ret
}

func genMat[T any](c, r int, d ...T) [][]T {
	ret := make([][]T, c)
	for i := 0; i < c; i++ {
		ret[i] = make([]T, r)
	}

	if len(d) > 0 {
		for i := 0; i < c; i++ {
			for j := 0; j < r; j++ {
				ret[i][j] = d[0]
			}
		}
	}
	return ret
}

func ifv[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}

func beautifulSubsets(nums []int, k int) int {
	bads := make([]int, 0, len(nums)*len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				k := (1 << i) | (1 << j)
				bads = append(bads, k)
			}
		}
	}

	total := 1 << len(nums)
	ret := 0
	for i := 1; i < total; i++ {
		isBad := false
		for _, b := range bads {
			if i&b == b {
				isBad = true
				break
			}
		}
		if !isBad {
			ret++
		}
	}

	return ret
}

type Case struct {
	nums []int
	k    int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{
			[]int{2, 4, 6}, 2,
		},
		{
			[]int{1}, 1,
		},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := beautifulSubsets(c.nums, c.k)
		fmt.Println(">>>>>>>>", c, r)
	}
}
