// 332. Reconstruct Itinerary

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

func maxs(as ...int) int {
	r := as[0]
	for _, n := range as {
		r = max(n, r)
	}
	return r
}

type Queue struct {
	data  []int
	front int
	back  int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		data:  make([]int, capacity+1),
		front: 0,
		back:  0,
	}
}

func (q *Queue) Size() int {
	return ((q.back + len(q.data)) - q.front) % len(q.data)
}

func (q *Queue) PushBack(x int) bool {
	if q.Size()+1 < len(q.data) {
		q.data[q.back] = x
		q.back = (q.back + 1) % len(q.data)
		return true
	}

	return false
}

func (q *Queue) Back() (int, bool) {
	if q.Size() > 0 {
		bi := (q.back + len(q.data) - 1) % len(q.data)
		return q.data[bi], true
	}
	return 0, false
}

func (q *Queue) PopBack() (int, bool) {
	if q.Size() > 0 {
		bi := (q.back + len(q.data) - 1) % len(q.data)
		v := q.data[bi]
		q.back = bi
		return v, true
	}
	return 0, false
}

func (q *Queue) Front() (int, bool) {
	if q.Size() > 0 {
		return q.data[q.front], true
	}
	return 0, false
}

func (q *Queue) PopFront() (int, bool) {
	if q.Size() > 0 {
		v := q.data[q.front]
		q.front = (q.front + 1) % len(q.data)
		return v, true
	}
	return 0, false
}

func (q *Queue) SmallPush(cur int) {
	for {
		v, ok := q.Back()
		if ok && v > cur {
			q.PopBack()
		} else {
			break
		}
	}
	q.PushBack(cur)
}

func (q *Queue) BigPush(cur int) {
	for {
		v, ok := q.Back()
		if ok && v < cur {
			q.PopBack()
		} else {
			break
		}
	}
	q.PushBack(cur)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	bucket := make(map[int]int)
	mod := valueDiff + 1
	for i := 0; i < len(nums); i++ {
		fmt.Println(">>>", i, nums[i], bucket)
		if len(bucket) > indexDiff {
			last := nums[i-indexDiff-1]
			delete(bucket, last/mod)
			fmt.Println("last ", last, bucket)
		}

		num := nums[i]
		key := num / mod
		if num < 0 {
			key -= 1
		}

		pre, preok := bucket[key-1]
		_, curok := bucket[key]
		nxt, nxtok := bucket[key+1]
		if curok || (preok && abs(pre-num) <= valueDiff) || (nxtok && abs(nxt-num) <= valueDiff) {
			fmt.Println(key, num, "----", pre, preok, "---", curok, "----", nxt, nxtok)
			return true
		}
		bucket[key] = num
	}

	return false
}

type Case struct {
	nums      []int
	indexDiff int
	valueDiff int
}

func main() {
	cases := []Case{
		{
			[]int{1, 2, 3, 1}, 3, 0,
		},
		{
			[]int{1, 5, 9, 1, 5, 9}, 2, 3,
		},
		{
			[]int{-3, 3}, 2, 4,
		},
	}

	for i, c := range cases {
		fmt.Printf("start case_%d: %+v\n", i, c)
		res := containsNearbyAlmostDuplicate(c.nums, c.indexDiff, c.valueDiff)
		fmt.Println("result:", res)
	}

}
