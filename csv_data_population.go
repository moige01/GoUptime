// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type CSVDataPopulation struct{}

func (c *CSVDataPopulation) ReadFromCSV(h Handlers) []*Node {
	logger, err := NewFileLogger()

	if err != nil {
		log.Fatal(err)
	}

	in, err := os.OpenFile("csv/pages.csv", os.O_RDONLY, 0666)

	defer in.Close()

	if err != nil {
		logger.GetErrorLogger().Printf("Error trying to load CSV pages data: %s\n", err.Error())
	}

	r := csv.NewReader(in)
	var nodes []*Node

	i := 0
	for record, err := r.Read(); err != io.EOF; record, err = r.Read() {
		if err != nil {
			logger.GetErrorLogger().Println(err)
		}

		// Skip headers
		if i == 0 {
			i++
			continue
		}

		priority, err := strconv.Atoi(record[1])

		if err != nil {
			logger.GetErrorLogger().Println(err)
		}

		node := &Node{
			value: &Page{
				url:     record[0],
				handler: h.getHandler(record[2]),
			},
			index:    i,
			priority: priority,
		}

		nodes = append(nodes, node)

		i++
	}

	return nodes
}

func (c *CSVDataPopulation) Populate(h Handlers) *PriorityQueue {
	nodes := c.ReadFromCSV(h)
	pq := make(PriorityQueue, len(nodes))

	for i, node := range nodes {
		pq[i] = node
	}

	return &pq
}
