// Copyright 2015 The Macaron Authors
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
	"regexp"
)

type patternType int8

const (
	_PATTERN_STATIC    patternType = iota // /home
	_PATTERN_REGEXP                       // /:id([0-9]+)
	_PATTERN_PATH_EXT                     // /*.*
	_PATTERN_HOLDER                       // /:user
	_PATTERN_MATCH_ALL                    // /*
)

// Leaf represents a leaf route information.
type Leaf struct {
	parent *Tree

	typ        patternType
	pattern    string
	rawPattern string // Contains wildcard instead of regexp
	wildcards  []string
	reg        *regexp.Regexp
	optional   bool

	handle Handle
}

var wildcardPattern = regexp.MustCompile(`:[a-zA-Z0-9]+`)

func isSpecialRegexp(pattern, regStr string, pos []int) bool {
	_ = "STUB: not implemented"
	return false
}

// getNextWildcard tries to find next wildcard and update pattern with corresponding regexp.
func getNextWildcard(pattern string) (wildcard, _ string) { _ = "STUB: not implemented"; return "", "" }

// Reach last character or no regexp is given.

// Cut out placeholder directly.

func getWildcards(pattern string) (string, []string) { _ = "STUB: not implemented"; return "", nil }

// Keep getting next wildcard until nothing is left.

// getRawPattern removes all regexp but keeps wildcards for building URL path.
func getRawPattern(rawPattern string) string { _ = "STUB: not implemented"; return "" }

func checkPattern(pattern string) (typ patternType, rawPattern string, wildcards []string, reg *regexp.Regexp) {
	_ = "STUB: not implemented"
	return *new(patternType), "", nil, nil
}

func NewLeaf(parent *Tree, pattern string, handle Handle) *Leaf {
	_ = "STUB: not implemented"
	return nil
}

// URLPath build path part of URL by given pair values.
func (l *Leaf) URLPath(pairs ...string) string { _ = "STUB: not implemented"; return "" }

// Tree represents a router tree in Macaron.
type Tree struct {
	parent *Tree

	typ        patternType
	pattern    string
	rawPattern string
	wildcards  []string
	reg        *regexp.Regexp

	subtrees []*Tree
	leaves   []*Leaf
}

func NewSubtree(parent *Tree, pattern string) *Tree { _ = "STUB: not implemented"; return nil }

func NewTree() *Tree { _ = "STUB: not implemented"; return nil }

func (t *Tree) addLeaf(pattern string, handle Handle) *Leaf { _ = "STUB: not implemented"; return nil }

// Add exact same leaf to grandparent/parent level without optional.

// Root tree can add as empty pattern.

func (t *Tree) addSubtree(segment, pattern string, handle Handle) *Leaf {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tree) addNextSegment(pattern string, handle Handle) *Leaf {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tree) Add(pattern string, handle Handle) *Leaf { _ = "STUB: not implemented"; return nil }

func (t *Tree) matchLeaf(globLevel int, url string, params Params) (Handle, bool) {
	_ = "STUB: not implemented"
	return *new(Handle), false
}

// Number of results and wildcasrd should be exact same.

func (t *Tree) matchSubtree(globLevel int, segment, url string, params Params) (Handle, bool) {
	_ = "STUB: not implemented"
	return *new(Handle), false
}

func (t *Tree) matchNextSegment(globLevel int, url string, params Params) (Handle, bool) {
	_ = "STUB: not implemented"
	return *new(Handle), false
}

func (t *Tree) Match(url string) (Handle, Params, bool) {
	_ = "STUB: not implemented"
	return *new(Handle), *new(Params), false
}

// MatchTest returns true if given URL is matched by given pattern.
func MatchTest(pattern, url string) bool { _ = "STUB: not implemented"; return false }
