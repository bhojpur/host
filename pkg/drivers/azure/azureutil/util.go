package azureutil

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

import (
	"math/rand"
	"time"
)

/* Utilities */

// randomAzureStorageAccountName generates a valid storage account name prefixed
// with a predefined string. Availability of the name is not checked. Uses maximum
// length to maximise randomness.
func randomAzureStorageAccountName() string {
	const (
		maxLen = 24
		chars  = "0123456789abcdefghijklmnopqrstuvwxyz"
	)
	return storageAccountPrefix + randomString(maxLen-len(storageAccountPrefix), chars)
}

// randomString generates a random string of given length using specified alphabet.
func randomString(n int, alphabet string) string {
	r := timeSeed()
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[r.Intn(len(alphabet))]
	}
	return string(b)
}

func timeSeed() *rand.Rand { return rand.New(rand.NewSource(time.Now().UTC().UnixNano())) }
