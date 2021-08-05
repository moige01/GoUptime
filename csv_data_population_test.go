// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvDataPopulation(t *testing.T) {
	c := new(CSVDataPopulation)
	h := make(Handlers)

	h.registerHandler("STDOUT", new(StdoutHandler))

	nodes := c.ReadFromCSV(h)

	assert.Equal(t, 4, len(nodes))
}