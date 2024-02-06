// 332. Reconstruct Itinerary

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

func maxs(as ...int) int {
	r := as[0]
	for _, n := range as {
		r = max(n, r)
	}
	return r
}

type Edge struct {
	to    int32
	count int32
}

func findItinerary(tickets [][]string) []string {
	const start = "JFK"
	// nticket := len(tickets)

	ticketsmap := make(map[[2]string]int)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		key := [2]string{from, to}
		ticketsmap[key] = ticketsmap[key] + 1
	}

	tids := make(map[string]int32, len(tickets))
	for ticket := range ticketsmap {
		tids[ticket[0]] = 0
		tids[ticket[1]] = 0
	}

	tnames := make([]string, len(tids))
	i := 0
	for t := range tids {
		tnames[i] = t
		i++
	}
	sort.Slice(tnames, func(i, j int) bool { return tnames[i] < tnames[j] })

	for i, t := range tnames {
		tids[t] = int32(i)
	}

	graph := make([][]*Edge, len(tnames))
	for ticket, count := range ticketsmap {
		from, to := tids[ticket[0]], tids[ticket[1]]
		graph[from] = append(graph[from], &Edge{to, int32(count)})
	}
	for _, tos := range graph {
		sort.Slice(tos, func(i, j int) bool { return tos[i].to < tos[j].to })
	}
	// for from, tos := range graph {
	// 	sort.Slice(tos, func(i, j int) bool { return tos[i].to < tos[j].to })
	// 	fmt.Print("---", tnames[from], "  -->> ")
	// 	for _, edge := range tos {
	// 		fmt.Print(edge.to, ":", tnames[edge.to], "  ")
	// 	}
	// 	fmt.Println()
	// }

	resids := make([]int32, len(tickets)+1)
	var recurr func(cur int32, nedges int) bool
	recurr = func(cur int32, nedges int) bool {
		resids[nedges] = cur
		if nedges == len(tickets) {
			return true
		}

		for _, edge := range graph[cur] {
			if edge.count > 0 {
				edge.count--
				if recurr(edge.to, nedges+1) {
					return true
				}
				edge.count++
			}
		}
		return false
	}

	recurr(tids[start], 0)
	res := make([]string, len(tickets)+1)
	for i := 0; i <= len(tickets); i++ {
		res[i] = tnames[resids[i]]
	}

	return res
}

type Case struct {
	tickets [][]string
}

func main() {
	cases := []Case{
		{
			tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
		},
		{
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
		},
		{
			tickets: [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}},
		},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := findItinerary(c.tickets)
		fmt.Println("result:", res)
	}
}
