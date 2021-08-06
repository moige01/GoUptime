// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// This will define what a HandlerFunc must be. Every handler MUST implement the Dispatch
// method and manipulate the WaitGroup (mark as Done or Add-)

package main

import (
	"log"
	"sync"
)

type HandlerFunc interface {
	Dispatch(url string, status int, message string, wg *sync.WaitGroup)
}

type Handlers map[string]HandlerFunc

func (h Handlers) RegisterHandler(name string, handler HandlerFunc) {
	h[name] = handler
}

func (h Handlers) GetHandler(name string) HandlerFunc {
	logger, err := NewFileLogger()

	if err != nil {
		log.Fatal(err)
	}

	fn, ok := h[name]

	if !ok {
		logger.GetErrorLogger().Printf("Handler for: %s WAS NOT REGISTERED", name)

		return nil
	}

	return fn
}
