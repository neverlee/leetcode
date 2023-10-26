// 515. Find Largest Value in Each Tree Row
package main

import "fmt"

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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	l1, l2 := make([]*TreeNode, 0, 1000), make([]*TreeNode, 0, 1000)
	l1 = append(l1, root)
	rets := []int{}
	for len(l1) > 0 {
		max := -(0xffffffff)
		for _, node := range l1 {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				l2 = append(l2, node.Left)
			}
			if node.Right != nil {
				l2 = append(l2, node.Right)
			}
		}

		l1, l2 = l2, l1
		l2 = l2[:0]
		rets = append(rets, max)
	}
	return rets
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
