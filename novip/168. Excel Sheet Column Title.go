// 168. Excel Sheet Column Title

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

func convertToTitle(columnNumber int) string {
	m, n := columnNumber, 0
	for k := 1; m >= k; k *= 26 {
		m -= k
		n++
	}

	bs := make([]byte, n)
	for k := n - 1; k >= 0; k-- {
		bs[k] = byte(m%26) + 'A'
		m /= 26
	}

	return string(bs[:n])
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
