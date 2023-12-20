// 1359. Count All Valid Pickup and Delivery Options

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

func countOrders(n int) int {
	const mod = 1e9 + 7
	ret := 1
	for i := 1; i < n; i++ {
		ln := (2*i + 1)
		ln = (ln * (ln + 1)) / 2 % mod
		ret = (ret * ln) % mod
	}

	return ret
}

type Case struct {
	n int
}

func main() {
	cases := []Case{
		{
			n: 1,
		},
		{
			n: 2,
		},
		{
			n: 3,
		},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := countOrders(c.n)
		fmt.Println("result:", res)
	}
}
