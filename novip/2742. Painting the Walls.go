// 2742. Painting the Walls
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

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	xl, yl := n/2, n
	const ulimit = 0x7fffffff
	res := ulimit
	f := matrix(xl+1, yl+1, -1) // f[n][time] -> cost
	f[0][0] = 0
	for i := 0; i < n; i++ {
		c, t := cost[i], time[i]
		for x := xl; x >= 0; x-- {
			for y := yl; y >= x; y-- {
				if f[x][y] != -1 {
					nx, ny := x+1, y+t
					ncost := f[x][y] + c
					if nx <= xl && ny <= yl {
						if f[nx][ny] == -1 || f[nx][ny] > ncost {
							f[nx][ny] = ncost
						}
					}
					if n-nx <= ny {
						res = min(res, ncost)
					}
				}
			}
		}
	}
	// for i := 0; i <= xl; i++ { fmt.Println(">>>", f[i]) }

	return res
}

type Case struct {
	cost []int
	time []int
}

func main() {
	cases := []Case{
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 2}},
		{[]int{2, 3, 4, 2}, []int{1, 1, 1, 1}},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := paintWalls(c.cost, c.time)
		fmt.Println("result:", res)
	}
}
