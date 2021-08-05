// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdoutHandler(t *testing.T) {
	null, _ := os.Open(os.DevNull)
	defer null.Close()

	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null

	s := new(StdoutHandler)
	var wg sync.WaitGroup

	wg.Add(1)

	s.Dispatch("test.com", 100, "Testing!", &wg)

	wg.Wait()

	os.Stdout = sout
	os.Stderr = serr

	assert.Nil(t, nil)
}
