// 912. Sort an Array

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

func qsortRecursion(nums []int, low, high int) {
	if low >= high {
		return
	}
	t, i, k := nums[low], low+1, low
	for ; i <= high; i++ {
		if nums[i] < t {
			nums[k] = nums[i]
			nums[i] = nums[k+1]
			k++
		}
	}
	nums[k] = t
	qsortRecursion(nums, low, k-1)
	qsortRecursion(nums, k+1, high)
}

type Stack[T any] struct {
	data []T
}

func NewStack[T any](n int) *Stack[T] {
	s := &Stack[T]{
		data: make([]T, 0, n),
	}
	return s
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Push(a T) {
	s.data = append(s.data, a)
}

func (s *Stack[T]) Pop() T {
	n := len(s.data)
	v := s.data[n-1]
	s.data = s.data[:n-1]
	return v
}

func qsortNonRecursion(nums []int) {
	stack := NewStack[[2]int](200)
	stack.Push([2]int{0, len(nums) - 1})
	for stack.Len() > 0 {
		top := stack.Pop()
		low, high := top[0], top[1]
		if low >= high {
			continue
		}
		t, i, k := nums[low], low+1, low
		for ; i <= high; i++ {
			if nums[i] < t {
				nums[k] = nums[i]
				nums[i] = nums[k+1]
				k++
			}
		}
		nums[k] = t

		stack.Push([2]int{low, k - 1})
		stack.Push([2]int{k + 1, high})
	}
}

func sortArray(nums []int) []int {
	// qsortRecursion(nums, 0, len(nums)-1)
	qsortNonRecursion(nums)
	return nums
}

type Case struct {
	nums []int
}

func main() {
	fmt.Println("start")
	cases := []Case{
		{
			[]int{5, 1, 1, 2, 0, 0},
		},
		{
			[]int{3, 1, 4, 2, 5, 2, 7},
		},
	}
	for i, c := range cases {
		_ = c
		fmt.Println(i, "========================================")
		r := sortArray(c.nums)
		fmt.Println(">>>>>>>>", c, r)
	}
}
