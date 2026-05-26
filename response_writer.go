// Copyright 2013 Martini Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package macaron

import (
	"bufio"
	"net"
	"net/http"
)

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type ResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	http.Pusher
	// Status returns the status code of the response or 0 if the response has not been written.
	Status() int
	// Written returns whether or not the ResponseWriter has been written.
	Written() bool
	// Size returns the size of the response body.
	Size() int
	// Before allows for a function to be called before the ResponseWriter has been written to. This is
	// useful for setting headers or any other operations that must happen before a response has been written.
	Before(BeforeFunc)
}

// BeforeFunc is a function that is called before the ResponseWriter has been written to.
type BeforeFunc func(ResponseWriter)

// NewResponseWriter creates a ResponseWriter that wraps an http.ResponseWriter
func NewResponseWriter(method string, rw http.ResponseWriter) ResponseWriter {
	_ = "STUB: not implemented"
	return *new(ResponseWriter)
}

type responseWriter struct {
	method string
	http.ResponseWriter
	status      int
	size        int
	beforeFuncs []BeforeFunc
}

func (rw *responseWriter) WriteHeader(s int) { _ = "STUB: not implemented"; return }

func (rw *responseWriter) Write(b []byte) (size int, err error) {
	_ = "STUB: not implemented"

	// The status will be StatusOK if WriteHeader has not been called yet
	return 0, nil
}

func (rw *responseWriter) Status() int { _ = "STUB: not implemented"; return 0 }

func (rw *responseWriter) Size() int { _ = "STUB: not implemented"; return 0 }

func (rw *responseWriter) Written() bool { _ = "STUB: not implemented"; return false }

func (rw *responseWriter) Before(before BeforeFunc) { _ = "STUB: not implemented"; return }

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

// nolint
func (rw *responseWriter) CloseNotify() <-chan bool { _ = "STUB: not implemented"; return nil }

func (rw *responseWriter) callBefore() { _ = "STUB: not implemented"; return }

func (rw *responseWriter) Flush() { _ = "STUB: not implemented"; return }

func (rw *responseWriter) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}
