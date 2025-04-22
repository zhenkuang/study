package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}
type MyHeap []*ListNode

func (m MyHeap) Len() int {
	return len(m)
}

func (m MyHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m MyHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MyHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *MyHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var h = new(MyHeap)
	heap.Init(h)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	var ans = new(ListNode)
	var cur = ans
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return ans.Next
}
