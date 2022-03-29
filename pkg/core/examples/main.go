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

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bhojpur/host/pkg/core"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/drivers/virtualbox"
)

func usage() {
	fmt.Println("Usage: go run main.go <example>\n" +
		"Available examples: create streaming.")
	os.Exit(1)
}

// Sample Virtualbox create independent of Machine CLI.
func create() {
	log.SetDebug(true)

	client := core.NewClient("/tmp/automatic", "/tmp/automatic/certs")
	defer client.Close()

	hostName := "myfunhost"

	// Set some options on the provider...
	driver := virtualbox.NewDriver(hostName, "/tmp/automatic")
	driver.CPU = 2
	driver.Memory = 2048

	data, err := json.Marshal(driver)
	if err != nil {
		log.Error(err)
		return
	}

	h, err := client.NewHost("virtualbox", data)
	if err != nil {
		log.Error(err)
		return
	}

	h.HostOptions.EngineOptions.StorageDriver = "overlay"

	if err := client.Create(h); err != nil {
		log.Error(err)
		return
	}

	out, err := h.RunSSHCommand("df -h")
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("Results of your disk space query:\n%s\n", out)

	fmt.Println("Powering down machine now...")
	if err := h.Stop(); err != nil {
		log.Error(err)
		return
	}
}

// Streaming the output of an SSH session in virtualbox.
func streaming() {
	log.SetDebug(true)

	client := core.NewClient("/tmp/automatic", "/tmp/automatic/certs")
	defer client.Close()

	hostName := "myfunhost"

	// Set some options on the provider...
	driver := virtualbox.NewDriver(hostName, "/tmp/automatic")
	data, err := json.Marshal(driver)
	if err != nil {
		log.Error(err)
		return
	}

	h, err := client.NewHost("virtualbox", data)
	if err != nil {
		log.Error(err)
		return
	}

	if err := client.Create(h); err != nil {
		log.Error(err)
		return
	}

	h.HostOptions.EngineOptions.StorageDriver = "overlay"

	sshClient, err := h.CreateSSHClient()
	if err != nil {
		log.Error(err)
		return
	}

	stdout, stderr, err := sshClient.Start("yes | head -n 10000")
	if err != nil {
		log.Error(err)
		return
	}
	defer func() {
		_ = stdout.Close()
		_ = stderr.Close()
	}()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Error(err)
	}
	if err := sshClient.Wait(); err != nil {
		log.Error(err)
	}

	fmt.Println("Powering down machine now...")
	if err := h.Stop(); err != nil {
		log.Error(err)
		return
	}
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	switch os.Args[1] {
	case "create":
		create()
	case "streaming":
		streaming()
	default:
		usage()
	}
}
