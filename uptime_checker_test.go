// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "testing"

func testData() PriorityQueue {
	items := map[string]int{
		"https://www.google.com":             3,
		"https://www.facebook.com":           10,
		"https://www.alibaba.com":            202,
		"linkedin.com":                       1,
		"http://thisdomaindoesnotexists.com": 2,
	}

	pq := make(PriorityQueue, len(items))
	handler := new(StdoutHandler)

	i := 0
	for url, priority := range items {
		pq[i] = &Node{
			value: &Page{
				url:     url,
				handler: handler,
			},
			index:    i,
			priority: priority,
		}

		i++
	}

	return pq
}

func TestUptimeChecker(t *testing.T) {
	pq := testData()

	ut := new(UptimeChecker)
	ut.pages = &pq

	ut.VerifyStatus()
}
