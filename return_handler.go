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
	"reflect"
)

// ReturnHandler is a service that Martini provides that is called
// when a route handler returns something. The ReturnHandler is
// responsible for writing to the ResponseWriter based on the values
// that are passed into this function.
type ReturnHandler func(*Context, []reflect.Value)

func canDeref(val reflect.Value) bool { _ = "STUB: not implemented"; return false }

func isError(val reflect.Value) bool { _ = "STUB: not implemented"; return false }

func isByteSlice(val reflect.Value) bool { _ = "STUB: not implemented"; return false }

func defaultReturnHandler() ReturnHandler { _ = "STUB: not implemented"; return *new(ReturnHandler) }

// Ignore nil error
