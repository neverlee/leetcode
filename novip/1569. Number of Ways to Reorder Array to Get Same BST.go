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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BstOutput(bn *TreeNode, lv int, cs string) {
	if bn == nil {
		return
	}
	for i := 0; i < lv; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("%d %s\n", bn.Val, cs)
	BstOutput(bn.Left, lv+1, "l")
	BstOutput(bn.Right, lv+1, "r")
}

func BstInsert(bn **TreeNode, val int) bool {
	cur := *bn
	if cur == nil {
		*bn = &TreeNode{Val: val, Left: nil, Right: nil}
		return true
	}

	var cb bool
	if val < cur.Val {
		cb = BstInsert(&(cur.Left), val)
	} else if cur.Val < val {
		cb = BstInsert(&(cur.Right), val)
	}
	return cb
}

func numOfWays(nums []int) int {
	const mod = 1e9 + 7
	var bst *TreeNode
	for _, n := range nums {
		BstInsert(&bst, n)
	}

	n := len(nums)
	cc := matrix(n+1, n+1, 0)
	cc[0][0] = 1
	for i := 1; i <= n; i++ {
		cc[i][0] = 1
		for j := 1; j <= i; j++ {
			cc[i][j] = (cc[i-1][j] + cc[i-1][j-1]) % mod
		}
	}

	var drs func(bn *TreeNode) (int, int)
	drs = func(bn *TreeNode) (int, int) {
		if bn == nil {
			return 1, 0
		}
		ls, ln := drs(bn.Left)
		rs, rn := drs(bn.Right)

		return (ls * rs % mod) * cc[ln+rn][ln] % mod, ln + rn + 1
	}

	s, _ := drs(bst)
	return (s + mod - 1) % mod
}

type Case struct {
	nums []int
}

func main() {
	cases := []Case{{[]int{2, 1, 3}}, {[]int{3, 4, 5, 1, 2}}, {[]int{1, 2, 3}}}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := numOfWays(c.nums)
		fmt.Println("result:", res)
	}
}
