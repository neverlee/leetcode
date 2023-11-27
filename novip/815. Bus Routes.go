// 815. Bus Routes

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

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	srmap := make(map[int][]int)
	for ri, route := range routes {
		for _, s := range route {
			srmap[s] = append(srmap[s], ri)
		}
	}

	rmat := make([]map[int]bool, len(routes))
	for i := range routes {
		rmat[i] = make(map[int]bool)
	}

	for _, rmap := range srmap {
		for _, sa := range rmap {
			for _, sb := range rmap {
				rmat[sa][sb] = true
				rmat[sb][sa] = true
			}
		}
	}

	targets := make(map[int]bool)
	for _, rt := range srmap[target] {
		targets[rt] = true
	}

	states := genVec(len(routes), -1)
	q1, q2 := make([]int, 0, 600), make([]int, 0, 600)
	for _, st := range srmap[source] {
		q1 = append(q1, st)
		states[st] = 1
	}

	for len(q1) > 0 {
		for _, cr := range q1 {
			for nr := range rmat[cr] {
				if states[nr] == -1 || states[nr] > states[cr]+1 {
					fmt.Println(">>> to to", cr, nr, states[cr], states[nr])
					states[nr] = states[cr] + 1
					q2 = append(q2, nr)
				}
			}
		}

		q1 = q1[:0]
		q1, q2 = q2, q1
	}

	// fmt.Println(">>>---", srmap)
	// fmt.Println(">>>---", rmat)
	// fmt.Println(" states", states)
	res := -1
	for t := range targets {
		if res == -1 || res > states[t] {
			res = states[t]
		}
	}

	return res
}

type Case struct {
	routes [][]int
	source int
	target int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{
			[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6,
		},
		{
			[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12,
		},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := numBusesToDestination(c.routes, c.source, c.target)
		fmt.Println(">>>>>>>>", c, r)
	}
}
