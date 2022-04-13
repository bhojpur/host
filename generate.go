//go:generate go run k8s.io/gengo/examples/deepcopy-gen --go-header-file ./scripts/boilerplate.go.txt --input-dirs ./pkg/engine/types --input-dirs ./pkg/engine/types/kdm --output-file-base zz_generated_deepcopy
//go:generate go run pkg/codegen/generator/cleanup/main.go
//go:generate go run ./pkg/codegen/main.go
//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o ./pkg/data/bindata.go -ignore bindata.go -pkg data -modtime 1557785965 -mode 0644 ./pkg/data/

package main

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
