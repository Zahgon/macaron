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
	"io"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/go-macaron/inject"
)

// Locale reprents a localization interface.
type Locale interface {
	Language() string
	Tr(string, ...interface{}) string
}

// RequestBody represents a request body.
type RequestBody struct {
	reader io.ReadCloser
}

// Bytes reads and returns content of request body in bytes.
func (rb *RequestBody) Bytes() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// String reads and returns content of request body in string.
		nil
}

func (rb *RequestBody) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadCloser returns a ReadCloser for request body.
func (rb *RequestBody) ReadCloser() io.ReadCloser {
	_ = "STUB: not implemented"

	// Request represents an HTTP request received by a server or to be sent by a client.
	return *new(io.ReadCloser)
}

type Request struct {
	*http.Request
}

// Body returns a RequestBody for the request
func (r *Request) Body() *RequestBody { _ = "STUB: not implemented"; return nil }

// ContextInvoker is an inject.FastInvoker wrapper of func(ctx *Context).
type ContextInvoker func(ctx *Context)

// Invoke implements inject.FastInvoker which simplifies calls of `func(ctx *Context)` function.
func (invoke ContextInvoker) Invoke(params []interface{}) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Context represents the runtime context of current request of Macaron instance.
// It is the integration of most frequently used middlewares and helper methods.
type Context struct {
	inject.Injector
	handlers []Handler
	action   Handler
	index    int

	*Router
	Req    Request
	Resp   ResponseWriter
	params Params
	Render
	Locale
	Data map[string]interface{}
}

func (ctx *Context) handler() Handler { _ = "STUB: not implemented"; return *new(Handler) }

// Next runs the next handler in the context chain
func (ctx *Context) Next() { _ = "STUB: not implemented"; return }

// Written returns whether the context response has been written to
func (ctx *Context) Written() bool { _ = "STUB: not implemented"; return false }

func (ctx *Context) run() { _ = "STUB: not implemented"; return }

// if the handler returned something, write it to the http response

// RemoteAddr returns more real IP address.
func (ctx *Context) RemoteAddr() string { _ = "STUB: not implemented"; return "" }

func (ctx *Context) renderHTML(status int, setName, tplName string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// HTML renders the HTML with default template set.
func (ctx *Context) HTML(status int, name string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// HTMLSet renders the HTML with given template set name.
func (ctx *Context) HTMLSet(status int, setName, tplName string, data ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Redirect sends a redirect response
func (ctx *Context) Redirect(location string, status ...int) { _ = "STUB: not implemented"; return }

// MaxMemory is the maximum amount of memory to use when parsing a multipart form.
// Set this to whatever value you prefer; default is 10 MB.
var MaxMemory = int64(1024 * 1024 * 10)

func (ctx *Context) parseForm() { _ = "STUB: not implemented"; return }

// Query querys form parameter.
func (ctx *Context) Query(name string) string { _ = "STUB: not implemented"; return "" }

// QueryTrim querys and trims spaces form parameter.
func (ctx *Context) QueryTrim(name string) string { _ = "STUB: not implemented"; return "" }

// QueryStrings returns a list of results by given query name.
func (ctx *Context) QueryStrings(name string) []string { _ = "STUB: not implemented"; return nil }

// QueryEscape returns escapred query result.
func (ctx *Context) QueryEscape(name string) string { _ = "STUB: not implemented"; return "" }

// QueryBool returns query result in bool type.
func (ctx *Context) QueryBool(name string) bool { _ = "STUB: not implemented"; return false }

// QueryInt returns query result in int type.
func (ctx *Context) QueryInt(name string) int { _ = "STUB: not implemented"; return 0 }

// QueryInt64 returns query result in int64 type.
func (ctx *Context) QueryInt64(name string) int64 { _ = "STUB: not implemented"; return 0 }

// QueryFloat64 returns query result in float64 type.
func (ctx *Context) QueryFloat64(name string) float64 { _ = "STUB: not implemented"; return 0 }

// Params returns value of given param name.
// e.g. ctx.Params(":uid") or ctx.Params("uid")
func (ctx *Context) Params(name string) string { _ = "STUB: not implemented"; return "" }

// AllParams returns all params.
func (ctx *Context) AllParams() Params {
	_ = "STUB: not implemented"

	// SetParams sets value of param with given name.
	return *new(Params)
}

func (ctx *Context) SetParams(name, val string) { _ = "STUB: not implemented"; return }

// ReplaceAllParams replace all current params with given params
func (ctx *Context) ReplaceAllParams(params Params) { _ = "STUB: not implemented"; return }

// ParamsEscape returns escapred params result.
// e.g. ctx.ParamsEscape(":uname")
func (ctx *Context) ParamsEscape(name string) string { _ = "STUB: not implemented"; return "" }

// ParamsInt returns params result in int type.
// e.g. ctx.ParamsInt(":uid")
func (ctx *Context) ParamsInt(name string) int { _ = "STUB: not implemented"; return 0 }

// ParamsInt64 returns params result in int64 type.
// e.g. ctx.ParamsInt64(":uid")
func (ctx *Context) ParamsInt64(name string) int64 { _ = "STUB: not implemented"; return 0 }

// ParamsFloat64 returns params result in int64 type.
// e.g. ctx.ParamsFloat64(":uid")
func (ctx *Context) ParamsFloat64(name string) float64 { _ = "STUB: not implemented"; return 0 }

// GetFile returns information about user upload file by given form field name.
func (ctx *Context) GetFile(name string) (multipart.File, *multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return *new(multipart.File), nil, nil
}

// SaveToFile reads a file from request by field name and saves to given path.
func (ctx *Context) SaveToFile(name, savePath string) error { _ = "STUB: not implemented"; return nil }

// SetCookie sets given cookie value to response header.
// FIXME: IE support? http://golanghome.com/post/620#reply2
func (ctx *Context) SetCookie(name string, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetCookie returns given cookie value from request header.
func (ctx *Context) GetCookie(name string) string { _ = "STUB: not implemented"; return "" }

// GetCookieInt returns cookie result in int type.
func (ctx *Context) GetCookieInt(name string) int { _ = "STUB: not implemented"; return 0 }

// GetCookieInt64 returns cookie result in int64 type.
func (ctx *Context) GetCookieInt64(name string) int64 { _ = "STUB: not implemented"; return 0 }

// GetCookieFloat64 returns cookie result in float64 type.
func (ctx *Context) GetCookieFloat64(name string) float64 { _ = "STUB: not implemented"; return 0 }

var defaultCookieSecret string

// SetDefaultCookieSecret sets global default secure cookie secret.
func (m *Macaron) SetDefaultCookieSecret(secret string) { _ = "STUB: not implemented"; return }

// SetSecureCookie sets given cookie value to response header with default secret string.
func (ctx *Context) SetSecureCookie(name, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetSecureCookie returns given cookie value from request header with default secret string.
func (ctx *Context) GetSecureCookie(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// SetSuperSecureCookie sets given cookie value to response header with secret string.
func (ctx *Context) SetSuperSecureCookie(secret, name, value string, others ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// GetSuperSecureCookie returns given cookie value from request header with secret string.
func (ctx *Context) GetSuperSecureCookie(secret, name string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ctx *Context) setRawContentHeader() { _ = "STUB: not implemented"; return }

// ServeContent serves given content to response.
func (ctx *Context) ServeContent(name string, r io.ReadSeeker, params ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// ServeFileContent serves given file as content to response.
func (ctx *Context) ServeFileContent(file string, names ...string) {
	_ = "STUB: not implemented"
	return
}

// ServeFile serves given file to response.
func (ctx *Context) ServeFile(file string, names ...string) { _ = "STUB: not implemented"; return }

// ChangeStaticPath changes static path from old to new one.
func (ctx *Context) ChangeStaticPath(oldPath, newPath string) { _ = "STUB: not implemented"; return }
