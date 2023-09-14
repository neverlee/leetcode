// 1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree

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

func kmpgen(s string) []int {
	next := make([]int, len(s))
	next[0] = 0
	now := 0
	for i := 1; i < len(s); {
		if s[i] == s[now] {
			i++
			now++
			next[i-1] = now
		} else if now > 0 {
			now = next[now-1]
		} else {
			next[i] = 0
			i++
		}

	}
	return next
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	next := kmpgen(s)
	k := n - next[n-1]

	fmt.Println("    ", []rune(s))
	fmt.Println("    ", next)
	if k > n/2 || n%k != 0 {
		return false
	}
	for i := 0; i+k < n; i++ {
		if s[i] != s[i+k] {
			return false
		}
	}

	return true
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
