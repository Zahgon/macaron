// Copyright 2013 Martini Authors
// Copyright 2014 The Macaron Authors
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
	"log"
	"net/http"
	"sync"
)

// StaticOptions is a struct for specifying configuration options for the macaron.Static middleware.
type StaticOptions struct {
	// Prefix is the optional prefix used to serve the static directory content
	Prefix string
	// SkipLogging will disable [Static] log messages when a static file is served.
	SkipLogging bool
	// IndexFile defines which file to serve as index if it exists.
	IndexFile string
	// Expires defines which user-defined function to use for producing a HTTP Expires Header
	// https://developers.google.com/speed/docs/insights/LeverageBrowserCaching
	Expires func() string
	// ETag defines if we should add an ETag header
	// https://developers.google.com/web/fundamentals/performance/optimizing-content-efficiency/http-caching#validating-cached-responses-with-etags
	ETag bool
	// FileSystem is the interface for supporting any implmentation of file system.
	FileSystem http.FileSystem
}

// FIXME: to be deleted.
type staticMap struct {
	lock sync.RWMutex
	data map[string]*http.Dir
}

func (sm *staticMap) Set(dir *http.Dir) { _ = "STUB: not implemented"; return }

func (sm *staticMap) Get(name string) *http.Dir { _ = "STUB: not implemented"; return nil }

func (sm *staticMap) Delete(name string) { _ = "STUB: not implemented"; return }

var statics = staticMap{sync.RWMutex{}, map[string]*http.Dir{}}

// staticFileSystem implements http.FileSystem interface.
type staticFileSystem struct {
	dir *http.Dir
}

func newStaticFileSystem(directory string) staticFileSystem {
	_ = "STUB: not implemented"
	return *new(staticFileSystem)
}

func (fs staticFileSystem) Open(name string) (http.File, error) {
	_ = "STUB: not implemented"
	return *new(http.File), nil
}

func prepareStaticOption(dir string, opt StaticOptions) StaticOptions {
	_ = "STUB: not implemented"
	// Defaults
	return *new(StaticOptions)
}

// Normalize the prefix if provided

// Ensure we have a leading '/'

// Remove any trailing '/'

func prepareStaticOptions(dir string, options []StaticOptions) StaticOptions {
	_ = "STUB: not implemented"
	return *new(StaticOptions)
}

func staticHandler(ctx *Context, log *log.Logger, opt StaticOptions) bool {
	_ = "STUB: not implemented"
	return false
}

// if we have a prefix, filter requests by stripping the prefix

// File exists but fail to open.

// Try to serve index file

// path.Clean removes the trailing slash, so we need to add it back when
// the original path has it.

// Redirect if missing trailing slash.

// Discard error.

// Add an Expires header to the static content

// GenerateETag generates an ETag based on size, filename and file modification time
func GenerateETag(fileSize, fileName, modTime string) string { _ = "STUB: not implemented"; return "" }

// Static returns a middleware handler that serves static files in the given directory.
func Static(directory string, staticOpt ...StaticOptions) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}

// Statics registers multiple static middleware handlers all at once.
func Statics(opt StaticOptions, dirs ...string) Handler {
	_ = "STUB: not implemented"
	return *new(Handler)
}
