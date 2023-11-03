// 1095. Find in Mountain Array

package main

import (
	"fmt"
	"math/rand"
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

func numWays(steps int, arrLen int) int {
	slen := min(steps/2+1, arrLen) + 2
	const mod = 1e9 + 7
	v1, v2 := make([]int, slen), make([]int, slen)
	v1[1] = 1
	for k := 0; k < steps; k++ {
		v1[0] = 0
		for i := 0; i < slen; i++ {
			v2[i] = 0
		}
		for i := 1; i < slen-1; i++ {
			v2[i] = (v2[i] + v1[i]) % mod
			v2[i-1] = (v2[i-1] + v1[i]) % mod
			v2[i+1] = (v2[i+1] + v1[i]) % mod
		}
		v1, v2 = v2, v1
	}

	return v1[1]
}

type MountainArray struct {
	data  []int
	count int
}

func (this *MountainArray) get(index int) int {
	v := this.data[index]
	this.count++
	return v
}
func (this *MountainArray) length() int {
	return len(this.data)
}

func Newm(n int) *MountainArray {
	m := MountainArray{
		data:  make([]int, n),
		count: 0,
	}
	hi := rand.Intn(n)
	m.data[hi] = 10000
	for i := hi - 1; i >= 0; i-- {
		m.data[i] = m.data[i+1] - 1 - rand.Intn(10)
	}
	for i := hi + 1; i < n; i++ {
		m.data[i] = m.data[i-1] - 1 - rand.Intn(10)
	}

	return &m
}

func Newv(n []int) *MountainArray {
	return &MountainArray{
		data: n,
	}
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	i, j, m := 0, n-1, 0
	li, lj := 0, n-1
	for i < j {
		m = (i + j) / 2
		c1 := mountainArr.get(m)
		c2 := mountainArr.get(m + 1)
		if c1 < c2 {
			if c1 < target {
				i = m + 1
			} else {
				j = m
			}
		} else {
			j = m
		}
	}
	if target == mountainArr.get(i) {
		return i
	}

	i, j = li, lj
	for i < j {
		m = (i + j + 1) / 2
		c1 := mountainArr.get(m)
		c2 := mountainArr.get(m - 1)
		if c1 < c2 {
			if c1 < target {
				j = m - 1
			} else {
				i = m
			}
		} else {
			i = m
		}
	}

	if target == mountainArr.get(i) {
		return i
	}
	return -1
}

type Case struct {
	array  []int
	target int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{[]int{1, 2, 3, 4, 5, 3, 1}, 3},
		{[]int{0, 1, 2, 4, 2, 1}, 3},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := findInMountainArray(c.target, Newv(c.array))
		fmt.Println(">>>>>>>>", c, r)
	}
	for i := 0; i < 1000; i++ {
		m := Newm(4000)
		t := rand.Intn(20000)
		r := findInMountainArray(t, m)
		// if m.count >= 100 {
		fmt.Println("n case ", i, t, r, m.count)
		// }

	}
}
