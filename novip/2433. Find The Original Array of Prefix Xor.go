// 2433. Find The Original Array of Prefix Xor
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

type Fact struct {
	a int
	b int
}

var wmap = [5][]int{
	{1},
	{0, 2},
	{0, 1, 3, 4},
	{2, 4},
	{0},
}

func findArray(pref []int) []int {
	n := len(pref)
	arr := make([]int, n)
	arr[0] = pref[0]
	k := 0
	for i := 1; i < n; i++ {
		k ^= arr[i-1]
		arr[i] = pref[i] ^ k
	}
	return arr
}

type Case struct {
	pref []int
}

func main() {
	cases := []Case{{[]int{5, 2, 0, 3, 1}}, {[]int{13}}}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := findArray(c.pref)
		fmt.Println("result:", res)
	}
}
