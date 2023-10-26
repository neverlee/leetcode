// 823. Binary Trees With Factors
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

type Fact struct {
	a int
	b int
}

func numFactoredBinaryTrees(arr []int) int {
	n := len(arr)
	const mod = 1e9 + 7
	ntree := make(map[int]int, n)
	facts := make(map[int][]Fact, n)
	for _, a := range arr {
		facts[a] = nil
	}

	for i, a := range arr {
		for j := i; j < n; j++ {
			b := arr[j]
			ab := a * b
			if farr, ok := facts[ab]; ok {
				facts[ab] = append(farr, Fact{a, b})
			}
		}
	}
	fmt.Println(facts)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	res := 0
	for _, a := range arr {
		cur := 1
		farr := facts[a]
		for _, fact := range farr {
			k := ntree[fact.a] * ntree[fact.b] % mod
			if fact.a != fact.b {
				k = k * 2 % mod
			}
			cur = (cur + k) % mod
		}
		ntree[a] = cur
		res = (res + cur) % mod
	}

	return res
}

type Case struct {
	arr []int
}

func main() {
	cases := []Case{{[]int{2, 4}}, {[]int{2, 4, 5, 10}}}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := numFactoredBinaryTrees(c.arr)
		fmt.Println("result:", res)
	}
}
