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
	"net/http"
	"sync"
)

var (
	// Known HTTP methods.
	_HTTP_METHODS = map[string]bool{
		"GET":     true,
		"POST":    true,
		"PUT":     true,
		"DELETE":  true,
		"PATCH":   true,
		"OPTIONS": true,
		"HEAD":    true,
	}
)

// routeMap represents a thread-safe map for route tree.
type routeMap struct {
	lock   sync.RWMutex
	routes map[string]map[string]*Leaf
}

// NewRouteMap initializes and returns a new routeMap.
func NewRouteMap() *routeMap { _ = "STUB: not implemented"; return nil }

// getLeaf returns Leaf object if a route has been registered.
func (rm *routeMap) getLeaf(method, pattern string) *Leaf { _ = "STUB: not implemented"; return nil }

// add adds new route to route tree map.
func (rm *routeMap) add(method, pattern string, leaf *Leaf) { _ = "STUB: not implemented"; return }

type group struct {
	pattern  string
	handlers []Handler
}

// Router represents a Macaron router layer.
type Router struct {
	m        *Macaron
	autoHead bool
	routers  map[string]*Tree
	*routeMap
	namedRoutes map[string]*Leaf

	groups              []group
	notFound            http.HandlerFunc
	internalServerError func(*Context, error)

	// handlerWrapper is used to wrap arbitrary function from Handler to inject.FastInvoker.
	handlerWrapper func(Handler) Handler
}

func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

// SetAutoHead sets the value who determines whether add HEAD method automatically
// when GET method is added.
func (r *Router) SetAutoHead(v bool) { _ = "STUB: not implemented"; return }

type Params map[string]string

// Handle is a function that can be registered to a route to handle HTTP requests.
// Like http.HandlerFunc, but has a third parameter for the values of wildcards (variables).
type Handle func(http.ResponseWriter, *http.Request, Params)

// Route represents a wrapper of leaf route and upper level router.
type Route struct {
	router *Router
	leaf   *Leaf
}

// Name sets name of route.
func (r *Route) Name(name string) { _ = "STUB: not implemented"; return }

// handle adds new route to the router tree.
func (r *Router) handle(method, pattern string, handle Handle) *Route {
	_ = "STUB: not implemented"
	return nil
}

// Prevent duplicate routes.

// Validate HTTP methods.

// Generate methods need register.

// Add to router tree.

// Handle registers a new request handle with the given pattern, method and handlers.
func (r *Router) Handle(method string, pattern string, handlers []Handler) *Route {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Group(pattern string, fn func(), h ...Handler) { _ = "STUB: not implemented"; return }

// Get is a shortcut for r.Handle("GET", pattern, handlers)
func (r *Router) Get(pattern string, h ...Handler) (leaf *Route) {
	_ = "STUB: not implemented"
	return nil
}

// Patch is a shortcut for r.Handle("PATCH", pattern, handlers)
func (r *Router) Patch(pattern string, h ...Handler) *Route { _ = "STUB: not implemented"; return nil }

// Post is a shortcut for r.Handle("POST", pattern, handlers)
func (r *Router) Post(pattern string, h ...Handler) *Route { _ = "STUB: not implemented"; return nil }

// Put is a shortcut for r.Handle("PUT", pattern, handlers)
func (r *Router) Put(pattern string, h ...Handler) *Route { _ = "STUB: not implemented"; return nil }

// Delete is a shortcut for r.Handle("DELETE", pattern, handlers)
func (r *Router) Delete(pattern string, h ...Handler) *Route { _ = "STUB: not implemented"; return nil }

// Options is a shortcut for r.Handle("OPTIONS", pattern, handlers)
func (r *Router) Options(pattern string, h ...Handler) *Route {
	_ = "STUB: not implemented"
	return nil
}

// Head is a shortcut for r.Handle("HEAD", pattern, handlers)
func (r *Router) Head(pattern string, h ...Handler) *Route { _ = "STUB: not implemented"; return nil }

// Any is a shortcut for r.Handle("*", pattern, handlers)
func (r *Router) Any(pattern string, h ...Handler) *Route { _ = "STUB: not implemented"; return nil }

// Route is a shortcut for same handlers but different HTTP methods.
//
// Example:
//
//	m.Route("/", "GET,POST", h)
func (r *Router) Route(pattern, methods string, h ...Handler) (route *Route) {
	_ = "STUB: not implemented"
	return nil
}

// Combo returns a combo router.
func (r *Router) Combo(pattern string, h ...Handler) *ComboRouter {
	_ = "STUB: not implemented"
	return nil
}

// NotFound configurates http.HandlerFunc which is called when no matching route is
// found. If it is not set, http.NotFound is used.
// Be sure to set 404 response code in your handler.
func (r *Router) NotFound(handlers ...Handler) { _ = "STUB: not implemented"; return }

// InternalServerError configurates handler which is called when route handler returns
// error. If it is not set, default handler is used.
// Be sure to set 500 response code in your handler.
func (r *Router) InternalServerError(handlers ...Handler) { _ = "STUB: not implemented"; return }

// SetHandlerWrapper sets handlerWrapper for the router.
func (r *Router) SetHandlerWrapper(f func(Handler) Handler) { _ = "STUB: not implemented"; return }

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Fast match for static routes

// Easy name.

// URLFor builds path part of URL by given pair values.
func (r *Router) URLFor(name string, pairs ...string) string { _ = "STUB: not implemented"; return "" }

// ComboRouter represents a combo router.
type ComboRouter struct {
	router   *Router
	pattern  string
	handlers []Handler
	methods  map[string]bool // Registered methods.

	lastRoute *Route
}

func (cr *ComboRouter) checkMethod(name string) { _ = "STUB: not implemented"; return }

func (cr *ComboRouter) route(fn func(string, ...Handler) *Route, method string, h ...Handler) *ComboRouter {
	_ = "STUB: not implemented"
	return nil
}

func (cr *ComboRouter) Get(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

func (cr *ComboRouter) Patch(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

func (cr *ComboRouter) Post(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

func (cr *ComboRouter) Put(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

func (cr *ComboRouter) Delete(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

func (cr *ComboRouter) Options(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

func (cr *ComboRouter) Head(h ...Handler) *ComboRouter { _ = "STUB: not implemented"; return nil }

// Name sets name of ComboRouter route.
func (cr *ComboRouter) Name(name string) { _ = "STUB: not implemented"; return }
