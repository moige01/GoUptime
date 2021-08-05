// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestData() PriorityQueue {
	items := map[string]int{
		"www.google.com":   3,
		"www.facebook.com": 10,
		"www.alibaba.com":  202,
		"linkedin.com":     -1,
	}

	pq := make(PriorityQueue, len(items))

	i := 0
	for url, priority := range items {
		pq[i] = &Node{
			value: &Page{
				url: url,
			},
			index:    i,
			priority: priority,
		}

		i++
	}

	return pq
}

func ExamplePriorityQueue() {
	pq := getTestData()
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Node)
		page := item.value.(*Page)

		fmt.Printf("%s => %d\n", page.url, item.priority)
	}

	// Output:
	// www.alibaba.com => 202
	// www.facebook.com => 10
	// www.google.com => 3
	// linkedin.com => -1
}

func TestPriorityQueue(t *testing.T) {
	assert := assert.New(t)

	pq := getTestData()

	heap.Init(&pq)

	assert.Greater(pq.Len(), 0)
	assert.Equal(pq.Len(), 4)

	item := heap.Pop(&pq).(*Node)
	page := item.value.(*Page)
	assert.Equal(page.url, "www.alibaba.com")
	assert.Equal(item.priority, 202)

	item = heap.Pop(&pq).(*Node)
	page = item.value.(*Page)
	assert.Equal(page.url, "www.facebook.com")
	assert.Equal(item.priority, 10)

	assert.Equal(pq.Len(), 2)

	node := &Node{
		value: &Page{
			url: "www.abc.es",
		},
		priority: -100,
	}

	heap.Push(&pq, node)
	pq.update(node, 4)

	assert.Equal(pq.Len(), 3)

	item = heap.Pop(&pq).(*Node)
	page = item.value.(*Page)
	assert.Equal(page.url, "www.abc.es")
	assert.Equal(item.priority, 4)
}
