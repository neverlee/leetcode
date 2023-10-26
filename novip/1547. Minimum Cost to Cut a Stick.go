// 1547. Minimum Cost to Cut a Stick
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

type Range struct {
	Start int
	End   int
}

func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i] < cuts[j]
	})
	clen := len(cuts)

	f := make(map[Range]int, len(cuts)*len(cuts))
	for i := 1; i < len(cuts); i++ {
		f[Range{i - 1, i}] = 0
	}
	for l := 2; l < clen; l++ {
		for s := 0; s+l < clen; s++ {
			e := s + l
			min := 0x7fffffff
			for m := s + 1; m < e; m++ {
				nf := f[Range{s, m}] + f[Range{m, e}] + cuts[e] - cuts[s]
				if nf < min {
					min = nf
				}
			}
			f[Range{s, e}] = min
		}
	}

	return f[Range{0, len(cuts) - 1}]
}

type Case struct {
}

func main() {
	cases := []Case{}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		// res := largestValues()
		// fmt.Println("result:", res)
	}
}
