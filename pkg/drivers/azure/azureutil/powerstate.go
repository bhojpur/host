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
	"strings"

	"github.com/bhojpur/host/pkg/machine/log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

// VMPowerState indicates the power state of the VM or is Unknown
type VMPowerState string

const (
	// Unknown is returned when Azure does not provide a PowerState (happens
	// when VM is just deployed or started transitioning to another state) or
	// obtained PowerState is not one of the following.
	Unknown VMPowerState = ""

	// Stopped indicates that VM is allocated and in powered off state or the VM
	// has been just deployed for the first time. In this state, VM can be powered
	// on or
	Stopped VMPowerState = "stopped"

	// Stopping indicates that VM is about to go into powered off state.
	Stopping VMPowerState = "stopping"

	// Starting indicates that VM is being created or powered on.
	Starting VMPowerState = "starting"

	// Running indicates that VM is either powered on or being rebooted. VM
	// stays in this state during the reboot operation. In this state VM can be
	// stopped, restarted or deallocated.
	Running VMPowerState = "running"

	// Deallocating indicates that the VM is being terminated.
	Deallocating VMPowerState = "deallocating"

	// Deallocated indicates that the VM is being terminated. In this state, VM
	// can be powered on or powered off.
	Deallocated VMPowerState = "deallocated"
)

const (
	powerStateCodePrefix = "PowerState/"
)

// powerStateFromInstanceView reads the instance view response and extracts the
// power state status (if exists) from there. If no status is found or an
// unknown status has occurred, returns Unknown.
func powerStateFromInstanceView(instanceView *compute.VirtualMachineInstanceView) VMPowerState {
	if instanceView == nil {
		log.Debug("Retrieved nil instance view.")
		return Unknown
	}
	if instanceView.Statuses == nil || len(*instanceView.Statuses) == 0 {
		log.Debug("Retrieved nil or empty instanceView.statuses.")
		return Unknown
	}
	statuses := *instanceView.Statuses

	// Filter statuses whose "code" starts with "PowerState/"
	var s *compute.InstanceViewStatus
	for _, v := range statuses {
		log.Debugf("Matching pattern for code=%q", to.String(v.Code))
		if strings.HasPrefix(to.String(v.Code), powerStateCodePrefix) {
			log.Debug("Power state found.")
			s = &v
			break
		}
	}
	if s == nil {
		log.Debug("No PowerState found in the instance view statuses.")
		return Unknown
	}
	code := strings.TrimPrefix(to.String(s.Code), powerStateCodePrefix)
	switch code {
	case "stopped":
		return Stopped
	case "stopping":
		return Stopping
	case "starting":
		return Starting
	case "running":
		return Running
	case "deallocated":
		return Deallocated
	case "deallocating":
		return Deallocating
	default:
		log.Warn("Encountered unknown PowerState for virtual machine: %q", code)
		return Unknown
	}
}
