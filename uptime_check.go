// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT
// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"container/heap"
	"net/http"
	"sync"
)

type UptimeChecker struct {
	pages *PriorityQueue
	wg    sync.WaitGroup
}

func getStatusAndMessageOfUrl(url string, sch chan int, mch chan string) {
	res, err := http.Get(url)

	if err != nil {
		sch <- 999
		mch <- err.Error()

		return
	}

	defer res.Body.Close()

	sch <- res.StatusCode
	mch <- res.Status
}

func (p *Page) GetStatus(wg *sync.WaitGroup) {
	sch := make(chan int)
	mch := make(chan string)

	go getStatusAndMessageOfUrl(p.url, sch, mch)

	status, message := <-sch, <-mch

	if p.handler == nil {
		WarningLogger.Printf(
			"Handler for %s doesn't exists. Seems like it was not properly registered in Init step",
			p.url,
		)

		wg.Done()
		return
	}

	go p.handler.Dispatch(p.url, status, message, wg)
}

func (uc *UptimeChecker) VerifyStatus() {
	heap.Init(uc.pages)

	uc.wg.Add(uc.pages.Len())

	for uc.pages.Len() > 0 {
		node := heap.Pop(uc.pages).(*Node)
		page := node.value.(*Page)

		go page.GetStatus(&uc.wg)
	}

	uc.wg.Wait()
}

func (uc *UptimeChecker) Init(p DataPopulation, h Handlers) {
	uc.pages = p.Populate(h)
}
