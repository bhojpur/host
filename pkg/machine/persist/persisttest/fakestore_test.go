package persisttest

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
	"fmt"
	"reflect"
	"testing"

	"github.com/bhojpur/host/pkg/machine/host"
)

func TestExists(t *testing.T) {
	store := FakeStore{
		Hosts: []*host.Host{{Name: "my-host"}},
	}
	exist, err := store.Exists("my-host")
	if err != nil {
		t.Fatal(err)
	}
	if exist == false {
		t.Fatal("Expected host 'my-host' to exist")
	}
	exist, err = store.Exists("not-found")
	if err != nil {
		t.Fatal(err)
	}
	if exist == true {
		t.Fatal("Expected host 'not-found' to no exist")
	}
	store.ExistsErr = fmt.Errorf("error checking host")
	exist, err = store.Exists("my-host")
	if err != store.ExistsErr {
		t.Fatalf("Expected err %s.", store.ExistsErr)
	}
}

func TestList(t *testing.T) {
	store := FakeStore{
		Hosts: []*host.Host{{Name: "my-host"}, {Name: "my-host-2"}},
	}
	list, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	expected := []string{"my-host", "my-host-2"}
	if !reflect.DeepEqual(list, expected) {
		t.Fatalf("Expected hosts to be %s. Got %s.", expected, list)
	}
	store.ListErr = fmt.Errorf("error listing")
	list, err = store.List()
	if err != store.ListErr {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	expectedHost := &host.Host{Name: "my-host"}
	store := FakeStore{
		Hosts: []*host.Host{expectedHost},
	}
	h, err := store.Load("my-host")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expectedHost, h) {
		t.Fatalf("Wrong host. Expected %v. Got %v.", expectedHost, h)
	}
	h, err = store.Load("not-found")
	if err != nil {
		t.Fatal(err)
	}
	if h != nil {
		t.Fatalf("Expected nil host. Got %v.", h)
	}
	store.LoadErr = fmt.Errorf("error loading")
	h, err = store.Load("my-host")
	if err != store.LoadErr {
		t.Fatalf("Wrong error. Expected %s. Got %s.", store.LoadErr, err)
	}
	if h != nil {
		t.Fatalf("Expected nil host. Got %v.", h)
	}
}

func TestRemove(t *testing.T) {
	store := FakeStore{
		Hosts: []*host.Host{{Name: "my-host"}},
	}
	err := store.Remove("not-found")
	if err != nil {
		t.Fatal(err)
	}
	err = store.Remove("my-host")
	if err != nil {
		t.Fatal(err)
	}
	if len(store.Hosts) != 0 {
		t.Fatalf("Expected hosts length to be zero. Got %d", len(store.Hosts))
	}
	store.RemoveErr = fmt.Errorf("error removing")
	err = store.Remove("my-host")
	if err != store.RemoveErr {
		t.Fatal(err)
	}
}

func TestSave(t *testing.T) {
	store := FakeStore{}
	err := store.Save(&host.Host{Name: "my-host"})
	if err != nil {
		t.Fatal(err)
	}
	expectedHosts := []*host.Host{{Name: "my-host"}}
	if !reflect.DeepEqual(store.Hosts, expectedHosts) {
		t.Fatalf("Expected hosts to be %v. Got %v.", expectedHosts, store.Hosts)
	}
	store.SaveErr = fmt.Errorf("error saving")
	err = store.Save(&host.Host{Name: "new-host"})
	if err != store.SaveErr {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(store.Hosts, expectedHosts) {
		t.Fatalf("Expected hosts to be %v. Got %v.", expectedHosts, store.Hosts)
	}
}
