// 987. Vertical Order Traversal of a Binary Tree

package main

import (
	"fmt"
	"sort"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	lv  int
	x   int
	val int
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := make([]*Node, 0, 1002)

	var recurr func(root *TreeNode, lv, x int)
	recurr = func(root *TreeNode, lv, x int) {
		if root == nil {
			return
		}

		nodes = append(nodes, &Node{lv, x, root.Val})

		recurr(root.Left, lv+1, x-1)
		recurr(root.Right, lv+1, x+1)
	}
	recurr(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].x == nodes[j].x {
			if nodes[i].lv == nodes[j].lv {
				return nodes[i].val < nodes[j].val
			} else if nodes[i].lv < nodes[j].lv {
				return true
			}
			return false
		} else if nodes[i].x < nodes[j].x {
			return true
		}
		return false
	})

	n := len(nodes)
	nodes = append(nodes, &Node{-1, -2000, 0})
	var results [][]int
	vals := make([]int, len(nodes))
	cur := 0
	for i := 0; i < n; i++ {
		vals[i] = nodes[i].val

		k := i + 1
		if nodes[cur].x != nodes[k].x {
			results = append(results, vals[cur:k])
			cur = k
		}
	}

	return results
}

type Case struct {
}

func main() {
	fmt.Println("start")
	cases := []Case{}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		// r := verticalTraversal(c.root)
		// fmt.Println(">>>>>>>>", c, r)
	}
}
