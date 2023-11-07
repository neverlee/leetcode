// 2328. Number of Increasing Paths in a Grid

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

type Pos struct {
	x int
	y int
}

var dirs = []Pos{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func countPaths(grid [][]int) int {
	const mod = 1e9 + 7

	xl, yl := len(grid), len(grid[0])
	pos := make([]Pos, xl*yl)
	for x := 0; x < xl; x++ {
		for y := 0; y < yl; y++ {
			pos[x*yl+y] = Pos{x, y}
		}
	}
	sort.Slice(pos, func(i, j int) bool {
		pi, pj := pos[i], pos[j]
		return grid[pi.x][pi.y] > grid[pj.x][pj.y]
	})

	fm := matrix(xl, yl, 1)
	for _, p := range pos {
		for _, d := range dirs {
			nx, ny := p.x+d.x, p.y+d.y
			if nx >= 0 && nx < xl && ny >= 0 && ny < yl && grid[p.x][p.y] > grid[nx][ny] {
				fm[nx][ny] = (fm[nx][ny] + fm[p.x][p.y]) % mod
			}
		}
	}
	res := 0
	for _, p := range pos {
		res = (res + fm[p.x][p.y]) % mod
	}

	return res
}

type Case struct {
	grid [][]int
}

func main() {
	cases := []Case{
		{[][]int{{1, 1}, {2, 3}}},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := countPaths(c.grid)
		fmt.Println("result:", res)
	}
}
