// 1458. Max Dot Product of Two Subsequences

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(ay []int) *ListNode {
	var ln *ListNode
	p := &ln
	for _, n := range ay {
		*p = &ListNode{Val: n, Next: nil}
		p = &((*p).Next)
	}
	return ln
}

func (ln *ListNode) Add(node *ListNode) {
	node.Next = ln.Next
	ln.Next = node
}

func (ln *ListNode) Print() {
	p := ln
	fmt.Print("List: ")
	for p != nil {
		fmt.Print(p.Val, ", ")
		p = p.Next
	}
	fmt.Println("")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func matrix(r, c, v int) [][]int {
	m := make([][]int, r)
	for i := 0; i < r; i++ {
		m[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m[i][j] = v
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxs(as ...int) int {
	r := as[0]
	for _, n := range as {
		r = max(n, r)
	}
	return r
}

func maxDotProduct(nums1 []int, nums2 []int) int {
	const minlmt = -600 * 1000 * 1000
	f := matrix(len(nums1)+1, len(nums2)+1, minlmt)
	ret := minlmt
	for i, n1 := range nums1 {
		for j, n2 := range nums2 {
			f[i+1][j+1] = maxs(f[i][j+1], f[i+1][j], f[i][j]+n1*n2, n1*n2)
			ret = max(ret, f[i+1][j+1])
		}
	}

	return ret
}

type Case struct {
	nums1 []int
	nums2 []int
}

func main() {
	cases := []Case{
		{
			nums1: []int{2, 1, -2, 5}, nums2: []int{3, 0, -6},
		},
		{
			nums1: []int{3, -2}, nums2: []int{2, -6, 7},
		},
		{
			nums1: []int{-1, -1}, nums2: []int{1, 1},
		},
		{
			nums1: []int{-3, -8, 3, -10, 1, 3, 9}, nums2: []int{9, 2, 3, 7, -9, 1, -8, 5, -1, -1},
		},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := maxDotProduct(c.nums1, c.nums2)
		fmt.Println("result:", res)
	}
}
