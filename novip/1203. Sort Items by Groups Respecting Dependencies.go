// 2147. Number of Ways to Divide a Long Corridor

package main

import (
	"fmt"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func genVec[T any](l int, d ...T) []T {
	ret := make([]T, l)
	if len(d) > 0 {
		for i := 0; i < l; i++ {
			ret[i] = d[0]
		}
	}
	return ret
}

func genMat[T any](c, r int, d ...T) [][]T {
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

type Node struct {
	gnID    int
	isGroup bool
}

func toposort(counts map[int]int, graph map[int][]int) []int {
	cur, topos := 0, make([]int, len(counts))

	for len(counts) > 0 {
		have := 0
		for ng, c := range counts {
			if c != 0 {
				continue
			}

			have++

			topos[cur] = ng
			cur++
			delete(counts, ng)

			ngts := graph[ng]
			for _, ngt := range ngts {
				if ngtc, ok := counts[ngt]; ok {
					ngtc--
					if counts[ngt] == 0 {
						topos[cur] = ngtc
						cur++
						delete(counts, ngt)
					}
					counts[ngt] = ngtc
				}
			}
		}

		if have == 0 && len(counts) != 0 {
			return nil
		}
	}

	return topos
}

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	const gbase = 1e6
	getGNID := func(nid int) int {
		if gnid := group[nid]; gnid != -1 {
			return gnid + gbase
		}
		return nid
	}

	ngraph := make(map[int][]int, len(beforeItems))
	ggraph := make(map[int][]int)
	for i, bis := range beforeItems {
		gna := getGNID(i)
		for _, item := range bis {
			gnb := getGNID(item)
			if gna != gnb {
				ggraph[gnb] = append(ggraph[gnb], gna)
			} else {
				ngraph[item] = append(ngraph[item], i)
			}
		}
	}

	gcounts := make(map[int]int)
	for i := range group {
		gcounts[getGNID(i)] = 0
	}
	for _, el := range ggraph {
		for _, t := range el {
			gcounts[t] = gcounts[t] + 1
		}
	}

	groupset := make(map[int][]int)
	for n := range group {
		if g := getGNID(n); g >= gbase {
			groupset[g] = append(groupset[g], n)
		}
	}

	ncounts := make(map[int]int)
	for _, el := range ngraph {
		for _, t := range el {
			ncounts[t] = ncounts[t] + 1
		}
	}

	topoGroup := toposort(gcounts, ggraph)
	if topoGroup == nil {
		return nil
	}

	rets := make([]int, 0, len(group))
	for _, ng := range topoGroup {
		if ng < gbase {
			rets = append(rets, ng)
		} else {
			counts := make(map[int]int)
			for _, n := range groupset[ng] {
				counts[n] = ncounts[n]
			}
			topoNode := toposort(counts, ngraph)
			// fmt.Println(">>>>", counts, topoNode)
			if topoNode == nil {
				return nil
			}
			rets = append(rets, topoNode...)
		}
	}

	// fmt.Println("ggraph", ggraph)
	// fmt.Println("ngraph", ngraph)
	// fmt.Println("counts", gcounts)
	// fmt.Println("groupset", groupset)
	// fmt.Println("topose", topoGroup)

	return rets
}

type Case struct {
	n           int
	m           int
	group       []int
	beforeItems [][]int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{
			8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1},
			[][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}},
		},
		{
			8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1},
			[][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}},
		},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := sortItems(c.n, c.m, c.group, c.beforeItems)
		fmt.Println(">>>>>>>>", c, r)
	}
}
