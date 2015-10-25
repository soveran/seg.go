// Copyright (c) 2015 Michel Martens
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package seg

import "strings"

type Seg struct {
	path string
	pos  int
}

func (s *Seg) Prev() string {
	return s.path[:s.pos-1]
}

func (s *Seg) Curr() string {
	return s.path[s.pos-1:]
}

func (s *Seg) IsRoot() bool {
	return s.pos >= len(s.path)
}

func (s *Seg) Consume(str string) bool {
	if s.IsRoot() {
		return false
	}

	found := s.find(len(str))

	if found == 0 || found == '/' {
		if s.subs(len(str)) == str {
			s.move(len(str))
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (s *Seg) Capture(key string, store map[string]string) bool {
	if s.IsRoot() {
		return false
	}

	index := strings.Index(s.path[s.pos:], "/")

	if index == -1 {
		index = len(s.path)
	} else {
		index += s.pos
	}

	store[key] = s.path[s.pos:index]

	s.move(index - s.pos)

	return true
}

func New(path string) *Seg {
	return &Seg { path: path, pos: 1 }
}

func (s *Seg) find(index int) byte {
	offset := s.pos + index

	if offset >= len(s.path) {
		return 0
	} else {
		return s.path[offset]
	}
}

func (s *Seg) subs(length int) string {
	offset := s.pos + length

	if offset > len(s.path) {
		return ""
	} else {
		return s.path[s.pos:offset]
	}
}

func (s *Seg) move(offset int) {
	s.pos += offset + 1
}
