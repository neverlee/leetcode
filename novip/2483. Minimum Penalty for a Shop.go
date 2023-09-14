// 2483. Minimum Penalty for a Shop

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

func bestClosingTime(customers string) int {
	yc := make([]int, len(customers)+2)
	for i := 1; i <= len(customers); i++ {
		yc[i] = yc[i-1] + ifv(customers[i-1] == 'Y', 1, 0)
	}
	ymax := yc[len(customers)]
	midx, minpenalty := 0, ymax
	for i := 0; i <= len(customers); i++ {
		penalty := i - yc[i] + ymax - yc[i]
		if penalty < minpenalty {
			midx = i
			minpenalty = penalty
		}
	}

	return midx
}

type Case struct {
	n string
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{n: "abcabdddabcabc"},
		{n: "ababcababc"},
		{n: "ababab"},
		{n: "aabbaabbaabb"},
		{n: "abab"},
		{n: "aba"},
		{n: "abcabcabcabc"},
		{n: "abac"},
		{n: "abaababaababaab"},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := repeatedSubstringPattern(c.n)
		fmt.Println(">>>>>>>>", c, r)
	}
}
