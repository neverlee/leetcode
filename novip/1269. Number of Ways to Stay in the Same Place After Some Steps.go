// 1269. Number of Ways to Stay in the Same Place After Some Steps

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

func numWays(steps int, arrLen int) int {
	slen := min(steps/2+1, arrLen) + 2
	const mod = 1e9 + 7
	v1, v2 := make([]int, slen), make([]int, slen)
	v1[1] = 1
	for k := 0; k < steps; k++ {
		v1[0] = 0
		for i := 0; i < slen; i++ {
			v2[i] = 0
		}
		for i := 1; i < slen-1; i++ {
			v2[i] = (v2[i] + v1[i]) % mod
			v2[i-1] = (v2[i-1] + v1[i]) % mod
			v2[i+1] = (v2[i+1] + v1[i]) % mod
		}
		v1, v2 = v2, v1
	}

	return v1[1]
}

type Case struct {
	steps  int
	arrLen int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{3, 2},
		{2, 4},
		{4, 2},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := numWays(c.steps, c.arrLen)
		fmt.Println(">>>>>>>>", c, r)
	}
}
