// 1220. Count Vowels Permutation
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

func countVowelPermutation(n int) int {
	const mod = 1e9 + 7
	v1, v2 := [5]int{}, [5]int{}
	for i := 0; i < 5; i++ {
		v1[i] = 1
	}
	for k := 1; k < n; k++ {
		for i := 0; i < 5; i++ {
			v2[i] = 0
		}
		for i, ws := range wmap {
			for _, w := range ws {
				v2[w] = (v2[w] + v1[i]) % mod
			}
		}
		v1, v2 = v2, v1
	}
	sum := 0
	for i := 0; i < 5; i++ {
		sum = (sum + v1[i]) % mod
	}

	return sum
}

type Case struct {
	n int
}

func main() {
	cases := []Case{{1}, {2}, {5}, {100}, {144}, {1000}}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := countVowelPermutation(c.n)
		fmt.Println("result:", res)
	}
}
