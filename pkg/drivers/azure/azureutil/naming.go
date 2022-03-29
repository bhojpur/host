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
	"fmt"
)

const (
	fmtNIC             = "%s-nic"
	fmtIP              = "%s-ip"
	fmtNSG             = "%s-firewall"
	fmtVM              = "%s"
	fmtOSDisk          = "%s-os-disk"
	fmtOSDiskContainer = "vhd-%s" // place vhds of VMs in separate containers for ease of cleanup
	fmtOSDiskBlob      = "%s-os-disk.vhd"
)

// ResourceNaming provides methods to construct Azure resource names for a given
// machine name.
type ResourceNaming string

// IP returns the Azure resource name for an IP address
func (r ResourceNaming) IP() string { return fmt.Sprintf(fmtIP, r) }

// NIC returns the Azure resource name for a network interface
func (r ResourceNaming) NIC() string { return fmt.Sprintf(fmtNIC, r) }

// NSG returns the Azure resource name for a network security group
func (r ResourceNaming) NSG() string { return fmt.Sprintf(fmtNSG, r) }

// VM returns the Azure resource name for a VM
func (r ResourceNaming) VM() string { return fmt.Sprintf(fmtVM, r) }

// OSDisk returns the Azure resource name for an OS disk
func (r ResourceNaming) OSDisk() string { return fmt.Sprintf(fmtOSDisk, r) }

// OSDiskContainer returns the Azure resource name for an OS disk container
func (r ResourceNaming) OSDiskContainer() string { return fmt.Sprintf(fmtOSDiskContainer, r) }

// OSDiskBlob returns the Azure resource name for an OS disk blob
func (r ResourceNaming) OSDiskBlob() string { return fmt.Sprintf(fmtOSDiskBlob, r) }
