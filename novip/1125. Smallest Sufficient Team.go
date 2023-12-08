// 1125. Smallest Sufficient Team

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

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	skillMap := make(map[string]int, len(req_skills))
	for i, skill := range req_skills {
		skillMap[skill] = i
	}
	bpeople := make([]int, len(people))
	for i, pskills := range people {
		bp := 0
		for _, s := range pskills {
			bp = bp | (1 << skillMap[s])
		}
		bpeople[i] = bp
	}

	x2n := 1 << (len(req_skills))
	states := make([]struct {
		num  int
		last int
		pre  int
	}, x2n)

	queue := make([]int, 1, x2n*2)
	nback := 0
	queue[0] = 0
	states[0].num = 0
	states[0].last = -1
	states[0].pre = -1

	for i := 1; i < x2n; i++ {
		states[i].num = 100
	}

	for nback < len(queue) {
		cur := queue[nback]
		nback++

		for ip, bp := range bpeople {
			next := cur | bp
			if states[next].num > states[cur].num+1 {
				states[next].num = states[cur].num + 1
				states[next].last = ip
				states[next].pre = cur
				queue = append(queue, next)
			}
		}
	}

	var rets []int
	cur := x2n - 1
	for cur > 0 {
		rets = append(rets, states[cur].last)
		cur = states[cur].pre
	}
	return rets
}

type Case struct {
	req_skills []string
	people     [][]string
}

func main() {
	cases := []Case{
		{
			[]string{"java", "nodejs", "reactjs"},
			[][]string{{"java", "nodejs"}, {"nodejs", "reactjs"}},
		},
		{
			[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
			[][]string{
				{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"},
			}},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := smallestSufficientTeam(c.req_skills, c.people)
		fmt.Println("result:", res)
	}
}
