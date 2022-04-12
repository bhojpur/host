package data

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

func MergeMaps(base, overlay map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range base {
		result[k] = v
	}
	for k, v := range overlay {
		if baseMap, overlayMap, bothMaps := bothMaps(result[k], v); bothMaps {
			v = MergeMaps(baseMap, overlayMap)
		}
		result[k] = v
	}
	return result
}

func bothMaps(left, right interface{}) (map[string]interface{}, map[string]interface{}, bool) {
	leftMap, ok := left.(map[string]interface{})
	if !ok {
		return nil, nil, false
	}
	rightMap, ok := right.(map[string]interface{})
	return leftMap, rightMap, ok
}

func bothSlices(left, right interface{}) ([]interface{}, []interface{}, bool) {
	leftSlice, ok := left.([]interface{})
	if !ok {
		return nil, nil, false
	}
	rightSlice, ok := right.([]interface{})
	return leftSlice, rightSlice, ok
}

func MergeMapsConcatSlice(base, overlay map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range base {
		result[k] = v
	}
	for k, v := range overlay {
		if baseMap, overlayMap, bothMaps := bothMaps(result[k], v); bothMaps {
			v = MergeMaps(baseMap, overlayMap)
		} else if baseSlice, overlaySlice, bothSlices := bothSlices(result[k], v); bothSlices {
			s := make([]interface{}, 0, len(baseSlice)+len(overlaySlice))
			s = append(s, baseSlice...)
			s = append(s, overlaySlice...)
			v = s
		}
		result[k] = v
	}
	return result

}
