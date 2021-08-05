// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// There is a implementation of a Priority Queue using the heap implementation
// built-in in Goglang.
//
// This will Pop us an element based on their proprity in and efficent way.

package main

import "container/heap"

type Node struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	Node := x.(*Node)
	Node.index = n
	*pq = append(*pq, Node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]

	return item
}

func (pq *PriorityQueue) update(item *Node, priority int) {
	item.priority = priority

	heap.Fix(pq, item.index)
}
