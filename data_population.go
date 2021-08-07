// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

type DataSource interface {
	Populate(h Handlers) *PriorityQueue
}
