// 1326. Minimum Number of Taps to Open to Water a Garden

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

func gen2darray[T any](c, r int, d ...T) [][]T {
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

func minTaps(n int, ranges []int) int {
	rareas := make([]int, len(ranges))
	for i, r := range ranges {
		left, right := i-r, i+r
		idx := max(left, 0)
		rareas[idx] = max(rareas[idx], right-idx)
	}

	need := 0
	cover := 0
	mc := 0
	for i := 0; i <= n; i++ {
		mc = max(mc, i+rareas[i])
		if i == cover {
			if mc <= cover {
				return -1
			}
			cover = mc
			need++
			if cover >= n {
				break
			}
		}
	}
	return need
}

type Case struct {
	n int
	r []int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{n: 5, r: []int{3, 4, 1, 1, 0, 0}},
		{n: 5, r: []int{3, 2, 3, 1, 0, 0}},
		{n: 3, r: []int{0, 0, 0, 0}},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := minTaps(c.n, c.r)
		fmt.Println(">>>>>>>>", c, r)
	}
}
