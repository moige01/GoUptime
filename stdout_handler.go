// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"sync"
)

type StdoutHandler struct{}

func (s *StdoutHandler) Dispatch(page string, status int, message string, wg *sync.WaitGroup) {
	defer wg.Done()

	statusOK := status >= 200 && status < 300

	if !statusOK {
		fmt.Printf("%s: %d - %s\n", page, status, message)
	}
}
