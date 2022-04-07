package plugin

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
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"

	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/drivers/plugin/localbinary"
	rpcdriver "github.com/bhojpur/host/pkg/machine/drivers/rpc"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/version"
)

var (
	heartbeatTimeout = 10 * time.Second
)

func RegisterDriver(d drivers.Driver) {
	if os.Getenv(localbinary.PluginEnvKey) != localbinary.PluginEnvVal {
		fmt.Fprintf(os.Stderr, `This is a Bhojpur Host machine plugin binary.
Plugin binaries are not intended to be invoked directly.
Please use this plugin through the main 'bhojpur-machine' binary.
(API version: %d)
`, version.APIVersion)
		os.Exit(1)
	}

	log.SetDebug(true)
	os.Setenv("MACHINE_DEBUG", "1")

	rpcd := rpcdriver.NewRPCServerDriver(d)
	rpc.RegisterName(rpcdriver.RPCServiceNameV0, rpcd)
	rpc.RegisterName(rpcdriver.RPCServiceNameV1, rpcd)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading RPC server: %s\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println(listener.Addr())

	go http.Serve(listener, nil)

	for {
		select {
		case <-rpcd.CloseCh:
			log.Debug("Closing plugin on server side")
			os.Exit(0)
		case <-rpcd.HeartbeatCh:
			continue
		case <-time.After(heartbeatTimeout):
			// TODO: Add heartbeat retry logic
			os.Exit(1)
		}
	}
}
