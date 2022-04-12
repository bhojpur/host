package templates

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

const VsphereCloudProviderTemplate = `
[Global]
user = "{{ .VsphereConfig.Global.User }}"
password = "{{ .VsphereConfig.Global.Password }}"
{{- if ne .VsphereConfig.Global.VCenterIP "" }}
server = "{{ .VsphereConfig.Global.VCenterIP }}"
{{- end }}
{{- if ne .VsphereConfig.Global.VCenterPort "" }}
port = "{{ .VsphereConfig.Global.VCenterPort }}"
{{- end }}
insecure-flag = "{{ .VsphereConfig.Global.InsecureFlag }}"
{{- if ne .VsphereConfig.Global.Datacenters "" }}
datacenters = "{{ .VsphereConfig.Global.Datacenters }}"
{{- end }}
{{- if ne .VsphereConfig.Global.Datacenter "" }}
datacenter = "{{ .VsphereConfig.Global.Datacenter }}"
{{- end }}
{{- if ne .VsphereConfig.Global.DefaultDatastore "" }}
datastore = "{{ .VsphereConfig.Global.DefaultDatastore }}"
{{- end }}
{{- if ne .VsphereConfig.Global.WorkingDir "" }}
working-dir = "{{ .VsphereConfig.Global.WorkingDir }}"
{{- end }}
soap-roundtrip-count = "{{ .VsphereConfig.Global.RoundTripperCount }}"
{{- if ne .VsphereConfig.Global.VMUUID "" }}
vm-uuid = "{{ .VsphereConfig.Global.VMUUID }}"
{{- end }}
{{- if ne .VsphereConfig.Global.VMName "" }}
vm-name = "{{ .VsphereConfig.Global.VMName }}"
{{- end }}

{{ range $k,$v := .VsphereConfig.VirtualCenter }}
[VirtualCenter "{{ $k }}"]
        user = "{{ $v.User }}"
        password = "{{ $v.Password }}"
        {{- if ne $v.VCenterPort "" }}
        port = "{{ $v.VCenterPort }}"
        {{- end }}
        {{- if ne $v.Datacenters "" }}
        datacenters = "{{ $v.Datacenters }}"
        {{- end }}
        soap-roundtrip-count = "{{ $v.RoundTripperCount }}"
{{- end }}

[Workspace]
        server = "{{ .VsphereConfig.Workspace.VCenterIP }}"
        datacenter = "{{ .VsphereConfig.Workspace.Datacenter }}"
        folder = "{{ .VsphereConfig.Workspace.Folder }}"
        default-datastore = "{{ .VsphereConfig.Workspace.DefaultDatastore }}"
        resourcepool-path = "{{ .VsphereConfig.Workspace.ResourcePoolPath }}"

[Disk]
        {{- if ne .VsphereConfig.Disk.SCSIControllerType "" }}
        scsicontrollertype = {{ .VsphereConfig.Disk.SCSIControllerType }}
        {{- end }}

[Network]
        {{- if ne .VsphereConfig.Network.PublicNetwork "" }}
        public-network = "{{ .VsphereConfig.Network.PublicNetwork }}"
        {{- end }}
`
