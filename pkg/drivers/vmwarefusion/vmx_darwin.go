package vmwarefusion

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

const vmx = `
 .encoding = "UTF-8"
 config.version = "8"
 displayName = "{{.MachineName}}"
 ethernet0.present = "TRUE"
 ethernet0.connectionType = "nat"
 ethernet0.virtualDev = "vmxnet3"
 ethernet0.wakeOnPcktRcv = "FALSE"
 ethernet0.addressType = "generated"
 ethernet0.linkStatePropagation.enable = "TRUE"
 pciBridge0.present = "TRUE"
 pciBridge4.present = "TRUE"
 pciBridge4.virtualDev = "pcieRootPort"
 pciBridge4.functions = "8"
 pciBridge5.present = "TRUE"
 pciBridge5.virtualDev = "pcieRootPort"
 pciBridge5.functions = "8"
 pciBridge6.present = "TRUE"
 pciBridge6.virtualDev = "pcieRootPort"
 pciBridge6.functions = "8"
 pciBridge7.present = "TRUE"
 pciBridge7.virtualDev = "pcieRootPort"
 pciBridge7.functions = "8"
 pciBridge0.pciSlotNumber = "17"
 pciBridge4.pciSlotNumber = "21"
 pciBridge5.pciSlotNumber = "22"
 pciBridge6.pciSlotNumber = "23"
 pciBridge7.pciSlotNumber = "24"
 scsi0.pciSlotNumber = "160"
 usb.pciSlotNumber = "32"
 ethernet0.pciSlotNumber = "192"
 sound.pciSlotNumber = "33"
 vmci0.pciSlotNumber = "35"
 sata0.pciSlotNumber = "36"
 floppy0.present = "FALSE"
 guestOS = "other3xlinux-64"
 hpet0.present = "TRUE"
 sata0.present = "TRUE"
 sata0:1.present = "TRUE"
 sata0:1.fileName = "{{.ISO}}"
 sata0:1.deviceType = "cdrom-image"
 {{ if .ConfigDriveURL }}
 sata0:2.present = "TRUE"
 sata0:2.fileName = "{{.ConfigDriveISO}}"
 sata0:2.deviceType = "cdrom-image"
 {{ end }}
 vmci0.present = "TRUE"
 mem.hotadd = "TRUE"
 memsize = "{{.Memory}}"
 powerType.powerOff = "soft"
 powerType.powerOn = "soft"
 powerType.reset = "soft"
 powerType.suspend = "soft"
 scsi0.present = "TRUE"
 scsi0.virtualDev = "pvscsi"
 scsi0:0.fileName = "{{.MachineName}}.vmdk"
 scsi0:0.present = "TRUE"
 tools.synctime = "TRUE"
 virtualHW.productCompatibility = "hosted"
 virtualHW.version = "10"
 msg.autoanswer = "TRUE"
 uuid.action = "create"
 numvcpus = "{{.CPU}}"
 hgfs.mapRootShare = "FALSE"
 hgfs.linkRootShare = "FALSE"
 `
