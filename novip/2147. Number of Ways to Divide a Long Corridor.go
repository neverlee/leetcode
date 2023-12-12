// 2147. Number of Ways to Divide a Long Corridor

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

func numberOfWays(corridor string) int {
	const mod = 1e9 + 7

	cc := 0
	last := 0
	rets := 1
	for i, c := range corridor {
		if c == 'S' {
			cc++
			if cc%2 == 0 {
				last = i
			} else {
				if last != 0 {
					// fmt.Println("---", i-last)
					rets = rets * (i - last) % mod
				}
			}
		}
	}
	if cc%2 == 1 || cc == 0 {
		return 0
	}

	return rets
}

type Case struct {
	corridor string
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{
			"SSPPSPS",
		},
		{
			"PPSPSP",
		},
		{
			"S",
		},
		{
			"P",
		},
		{
			"SSPPSPSPSPPSP",
		},
		{
			"PPSPPSSSSSSSPSPS",
		},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := numberOfWays(c.corridor)
		fmt.Println(">>>>>>>>", c, r)
	}
}
