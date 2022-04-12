package v1

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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ETCDSnapshotS3 struct {
	Endpoint            string `json:"endpoint,omitempty"`
	EndpointCA          string `json:"endpointCA,omitempty"`
	SkipSSLVerify       bool   `json:"skipSSLVerify,omitempty"`
	Bucket              string `json:"bucket,omitempty"`
	Region              string `json:"region,omitempty"`
	CloudCredentialName string `json:"cloudCredentialName,omitempty"`
	Folder              string `json:"folder,omitempty"`
}

type ETCDSnapshotCreate struct {
	// Changing the Generation is the only thing required to initiate a snapshot creation.
	Generation int `json:"generation,omitempty"`
}

type ETCDSnapshotRestore struct {
	// Name refers to the name of the associated etcdsnapshot object
	Name string `json:"name,omitempty"`

	// Changing the Generation is the only thing required to initiate a snapshot restore.
	Generation int `json:"generation,omitempty"`
	// Set to either none (or empty string), all, or kubernetesVersion
	RestoreBKEConfig string `json:"restoreBKEConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ETCDSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	SnapshotFile      ETCDSnapshotFile   `json:"snapshotFile,omitempty"`
	Status            ETCDSnapshotStatus `json:"status"`
}

type ETCDSnapshotFile struct {
	Name      string          `json:"name,omitempty"`
	NodeName  string          `json:"nodeName,omitempty"`
	Location  string          `json:"location,omitempty"`
	Metadata  string          `json:"metadata,omitempty"`
	CreatedAt *metav1.Time    `json:"createdAt,omitempty"`
	Size      int64           `json:"size,omitempty"`
	S3        *ETCDSnapshotS3 `json:"s3,omitempty"`
	Status    string          `json:"status,omitempty"`
	Message   string          `json:"message,omitempty"`
}

type ETCDSnapshotStatus struct {
	Missing bool `json:"missing"`
}

type ETCD struct {
	DisableSnapshots     bool            `json:"disableSnapshots,omitempty"`
	SnapshotScheduleCron string          `json:"snapshotScheduleCron,omitempty"`
	SnapshotRetention    int             `json:"snapshotRetention,omitempty"`
	S3                   *ETCDSnapshotS3 `json:"s3,omitempty"`
}
