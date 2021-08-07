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

	h.RegisterHandler("STDOUT", new(StdoutHandler))

	nodes := c.readFromCSV(h)

	assert.Equal(t, 4, len(nodes))
}
