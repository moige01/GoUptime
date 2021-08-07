// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
)

var (
	NotAFile = errors.New("Target is not a file")
)

type CSVDataSource struct{}

func (c *CSVDataSource) verifyMime(path string) bool {
	mtype, err := mimetype.DetectFile(path)

	if err != nil {
		ErrorLogger.Fatalln("Fail to get MIME type of given path")
	}

	return mtype.Is("text/csv")
}

func (c *CSVDataSource) verifyFile(path string) error {
	fi, err := os.Stat(path)

	if err != nil {
		return err
	}

	if !fi.Mode().IsRegular() {
		return NotAFile
	}

	return nil
}

func (c *CSVDataSource) verifyPath() string {
	path, ok := os.LookupEnv("CSV_PATH")

	if !ok {
		ErrorLogger.Fatalln("CSV_PATH MUST be filled ")
	}

	ok = filepath.IsAbs(path)

	if !ok {
		ErrorLogger.Fatalln("CSV_PATH MUST be an absolute path.")
	}

	return path
}

func (c *CSVDataSource) readFromCSV(h Handlers) []*Node {
	path := c.verifyPath()
	err := c.verifyFile(path)

	if err != nil {
		ErrorLogger.Fatalln(err)
	}

	ok := c.verifyMime(path)

	if !ok {
		ErrorLogger.Fatalf("%s seems to not be a valid CSV file\n", path)
	}

	in, err := os.OpenFile(path, os.O_RDONLY, 0666)

	defer in.Close()

	if err != nil {
		ErrorLogger.Printf("Error trying to load CSV pages data: %s\n", err.Error())
	}

	r := csv.NewReader(in)
	var nodes []*Node

	i := 0
	for record, err := r.Read(); err != io.EOF; record, err = r.Read() {
		if err != nil {
			ErrorLogger.Println(err)
		}

		// Skip headers
		if i == 0 {
			i++
			continue
		}

		priority, err := strconv.Atoi(record[1])

		if err != nil {
			ErrorLogger.Println(err)
		}

		node := &Node{
			value: &Page{
				url:     record[0],
				handler: h.GetHandler(record[2]),
			},
			index:    i,
			priority: priority,
		}

		nodes = append(nodes, node)

		i++
	}

	return nodes
}

func (c *CSVDataSource) Populate(h Handlers) *PriorityQueue {
	nodes := c.readFromCSV(h)
	pq := make(PriorityQueue, len(nodes))

	for i, node := range nodes {
		pq[i] = node
	}

	return &pq
}
