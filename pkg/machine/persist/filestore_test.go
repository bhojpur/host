package persist

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	"github.com/bhojpur/host/pkg/drivers/none"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/bhojpur/host/pkg/machine/hosttest"
)

func cleanup() {
	os.RemoveAll(os.Getenv("MACHINE_STORAGE_PATH"))
}

func getTestStore() Filestore {
	tmpDir, err := ioutil.TempDir("", "machine-test-")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mdirs.BaseDir = tmpDir

	return Filestore{
		Path:             tmpDir,
		CaCertPath:       filepath.Join(tmpDir, "certs", "ca-cert.pem"),
		CaPrivateKeyPath: filepath.Join(tmpDir, "certs", "ca-key.pem"),
	}
}

func TestStoreSave(t *testing.T) {
	defer cleanup()

	store := getTestStore()

	h, err := hosttest.GetDefaultTestHost()
	if err != nil {
		t.Fatal(err)
	}

	if err := store.Save(h); err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(store.GetMachinesDir(), h.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("Host path doesn't exist: %s", path)
	}

	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		r, err := regexp.Compile("config.json.tmp*")
		if err != nil {
			t.Fatalf("Failed to compile regexp string")
		}
		if r.MatchString(f.Name()) {
			t.Fatalf("Failed to remove temp filestore:%s", f.Name())
		}
	}
}

func TestStoreSaveOmitRawDriver(t *testing.T) {
	defer cleanup()

	store := getTestStore()

	h, err := hosttest.GetDefaultTestHost()
	if err != nil {
		t.Fatal(err)
	}

	if err := store.Save(h); err != nil {
		t.Fatal(err)
	}

	configJSONPath := filepath.Join(store.GetMachinesDir(), h.Name, "config.json")

	f, err := os.Open(configJSONPath)
	if err != nil {
		t.Fatal(err)
	}

	configData, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	fakeHost := make(map[string]interface{})

	if err := json.Unmarshal(configData, &fakeHost); err != nil {
		t.Fatal(err)
	}

	if rawDriver, ok := fakeHost["RawDriver"]; ok {
		t.Fatal("Should not have gotten a value for RawDriver reading host from disk but got one: ", rawDriver)
	}

}

func TestStoreRemove(t *testing.T) {
	defer cleanup()

	store := getTestStore()

	h, err := hosttest.GetDefaultTestHost()
	if err != nil {
		t.Fatal(err)
	}

	if err := store.Save(h); err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(store.GetMachinesDir(), h.Name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("Host path doesn't exist: %s", path)
	}

	err = store.Remove(h.Name)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(path); err == nil {
		t.Fatalf("Host path still exists after remove: %s", path)
	}
}

func TestStoreList(t *testing.T) {
	defer cleanup()

	store := getTestStore()

	h, err := hosttest.GetDefaultTestHost()
	if err != nil {
		t.Fatal(err)
	}

	if err := store.Save(h); err != nil {
		t.Fatal(err)
	}

	hosts, err := store.List()
	if len(hosts) != 1 {
		t.Fatalf("List returned %d items, expected 1", len(hosts))
	}

	if hosts[0] != h.Name {
		t.Fatalf("hosts[0] name is incorrect, got: %s", hosts[0])
	}
}

func TestStoreExists(t *testing.T) {
	defer cleanup()
	store := getTestStore()

	h, err := hosttest.GetDefaultTestHost()
	if err != nil {
		t.Fatal(err)
	}

	exists, err := store.Exists(h.Name)
	if exists {
		t.Fatal("Host should not exist before saving")
	}

	if err := store.Save(h); err != nil {
		t.Fatal(err)
	}

	exists, err = store.Exists(h.Name)
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("Host should exist after saving")
	}

	if err := store.Remove(h.Name); err != nil {
		t.Fatal(err)
	}

	exists, err = store.Exists(h.Name)
	if err != nil {
		t.Fatal(err)
	}

	if exists {
		t.Fatal("Host should not exist after removing")
	}
}

func TestStoreLoad(t *testing.T) {
	defer cleanup()

	expectedURL := "unix:///foo/baz"
	flags := hosttest.GetTestDriverFlags()
	flags.Data["url"] = expectedURL

	store := getTestStore()

	h, err := hosttest.GetDefaultTestHost()
	if err != nil {
		t.Fatal(err)
	}

	if err := h.Driver.SetConfigFromFlags(flags); err != nil {
		t.Fatal(err)
	}

	if err := store.Save(h); err != nil {
		t.Fatal(err)
	}

	h, err = store.Load(h.Name)
	if err != nil {
		t.Fatal(err)
	}

	rawDataDriver, ok := h.Driver.(*host.RawDataDriver)
	if !ok {
		t.Fatal("Expected driver loaded from store to be of type *host.RawDataDriver and it was not")
	}

	realDriver := none.NewDriver(h.Name, store.Path)

	if err := json.Unmarshal(rawDataDriver.Data, &realDriver); err != nil {
		t.Fatalf("Error unmarshaling rawDataDriver data into concrete 'none' driver: %s", err)
	}

	h.Driver = realDriver

	actualURL, err := h.URL()
	if err != nil {
		t.Fatal(err)
	}

	if actualURL != expectedURL {
		t.Fatalf("GetURL is not %q, got %q", expectedURL, actualURL)
	}
}