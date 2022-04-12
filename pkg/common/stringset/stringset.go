package stringset

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

var empty struct{}

// Set is an exceptionally simple `set` implementation for strings.
// It is not threadsafe, but can be used in place of a simple `map[string]struct{}`
// as long as you don't want to do too much with it.
type Set struct {
	m map[string]struct{}
}

func (s *Set) Add(ss ...string) {
	if s.m == nil {
		s.m = make(map[string]struct{}, len(ss))
	}
	for _, k := range ss {
		s.m[k] = empty
	}
}

func (s *Set) Delete(ss ...string) {
	if s.m == nil {
		return
	}
	for _, k := range ss {
		delete(s.m, k)
	}
}

func (s *Set) Has(ss string) bool {
	if s.m == nil {
		return false
	}
	_, ok := s.m[ss]
	return ok
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Values() []string {
	i := 0
	keys := make([]string, len(s.m))
	for key := range s.m {
		keys[i] = key
		i++
	}

	return keys
}
