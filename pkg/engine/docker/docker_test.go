package docker

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
	"testing"

	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/stretchr/testify/assert"
)

const (
	basicRepoUname = "basicUser"
	basicRepoPass  = "basicPass"
	basicImage     = "repo.com/bhojpur/bke-tools:v1"
	repoUname      = "user"
	repoPass       = "pass"
	image          = "repo.com/foo/bar/bhojpur/bke-tools:v1"
)

func TestPrivateRegistry(t *testing.T) {
	privateRegistries := map[string]v3.PrivateRegistry{}
	pr1 := v3.PrivateRegistry{
		URL:      "repo.com",
		User:     basicRepoUname,
		Password: basicRepoPass,
	}
	a1, err := getRegistryAuth(pr1)
	assert.Nil(t, err)
	privateRegistries[pr1.URL] = pr1

	pr2 := v3.PrivateRegistry{
		URL:      "repo.com/foo/bar",
		User:     repoUname,
		Password: repoPass,
	}
	a2, err := getRegistryAuth(pr2)
	assert.Nil(t, err)
	privateRegistries[pr2.URL] = pr2

	a, _, err := GetImageRegistryConfig(basicImage, privateRegistries)
	assert.Nil(t, err)
	assert.Equal(t, a, a1)

	a, _, err = GetImageRegistryConfig(image, privateRegistries)
	assert.Nil(t, err)
	assert.Equal(t, a, a2)

}

func TestGetKubeletDockerConfig(t *testing.T) {
	e := "{\"auths\":{\"https://registry.example.com\":{\"auth\":\"dXNlcjE6cGFzc3d+cmQ=\"}}}"
	c, err := GetKubeletDockerConfig(map[string]v3.PrivateRegistry{
		"https://registry.example.com": v3.PrivateRegistry{
			User:     "user1",
			Password: "passw~rd",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, c, e)
}
