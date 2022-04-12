package mapper

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
	"reflect"
	"testing"
)

func Test_OsInfo(t *testing.T) {
	mapper := OSInfo{}

	tests := []struct {
		internal map[string]interface{}
		wantInfo map[string]interface{}
	}{
		{
			internal: map[string]interface{}{
				"capacity": map[string]interface{}{
					"cpu":    "2",
					"memory": "123456Ki",
				},
			},
			wantInfo: map[string]interface{}{
				"cpu": map[string]interface{}{
					"count": int64(2),
				},
				"memory": map[string]interface{}{
					"memTotalKiB": int64(123456),
				},
				"os": map[string]interface{}{
					"dockerVersion":   "",
					"kernelVersion":   nil,
					"operatingSystem": nil,
				},
				"kubernetes": map[string]interface{}{
					"kubeletVersion":   nil,
					"kubeProxyVersion": nil,
				},
			},
		},
		{
			internal: map[string]interface{}{
				"capacity": map[string]interface{}{
					"cpu":    "1M",
					"memory": "123456Ti",
				},
			},
			wantInfo: map[string]interface{}{
				"cpu": map[string]interface{}{
					"count": int64(1000000),
				},
				"memory": map[string]interface{}{
					"memTotalKiB": int64(132559870623744),
				},
				"os": map[string]interface{}{
					"dockerVersion":   "",
					"kernelVersion":   nil,
					"operatingSystem": nil,
				},
				"kubernetes": map[string]interface{}{
					"kubeletVersion":   nil,
					"kubeProxyVersion": nil,
				},
			},
		},
	}

	for _, tt := range tests {
		mapper.FromInternal(tt.internal)
		if !reflect.DeepEqual(tt.wantInfo, tt.internal["info"]) {
			t.Fatal("os info does not match after mapping")
		}
	}
}
