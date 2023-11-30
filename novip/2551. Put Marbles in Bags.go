// 2551. Put Marbles in Bags

package main

import (
	"fmt"
	"sort"
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

func putMarbles(weights []int, k int) int64 {
	w := make([]int, len(weights)-1)
	for i := 0; i < len(w); i++ {
		w[i] = weights[i] + weights[i+1]
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})

	res := 0
	for i := 1; i < k; i++ {
		res += w[len(w)-i] - w[i-1]
	}
	return int64(res)
}

type Case struct {
	weights []int
	k       int
}

func main() {
	cases := []Case{
		{[]int{1, 3, 5, 1}, 2},
		{[]int{1, 3}, 2},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := putMarbles(c.weights, c.k)
		fmt.Println("result:", res)
	}
}
