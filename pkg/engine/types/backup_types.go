package types

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
	"github.com/bhojpur/host/pkg/core/condition"
	v1 "k8s.io/api/core/v1"
)

const (
	BackupConditionCreated   condition.Cond = "Created"
	BackupConditionCompleted condition.Cond = "Completed"
)

type BackupConfig struct {
	// Enable or disable recurring backups in Bhojpur Host
	Enabled *bool `yaml:"enabled" json:"enabled,omitempty" bhojpur:"default=true"`
	// Backup interval in hours
	IntervalHours int `yaml:"interval_hours" json:"intervalHours,omitempty" bhojpur:"default=12"`
	// Number of backups to keep
	Retention int `yaml:"retention" json:"retention,omitempty" bhojpur:"default=6"`
	// s3 target
	S3BackupConfig *S3BackupConfig `yaml:",omitempty" json:"s3BackupConfig"`
	// replace special characters in snapshot names
	SafeTimestamp bool `yaml:"safe_timestamp" json:"safeTimestamp,omitempty"`
	// Backup execution timeout
	Timeout int `yaml:"timeout" json:"timeout,omitempty" bhojpur:"default=300"`
}

type S3BackupConfig struct {
	// Access key ID
	AccessKey string `yaml:"access_key" json:"accessKey,omitempty"`
	// Secret access key
	SecretKey string `yaml:"secret_key" json:"secretKey,omitempty" bhojpur:"type=password" `
	// name of the bucket to use for backup
	BucketName string `yaml:"bucket_name" json:"bucketName,omitempty"`
	// AWS Region, AWS spcific
	Region string `yaml:"region" json:"region,omitempty"`
	// Endpoint is used if this is not an AWS API
	Endpoint string `yaml:"endpoint" json:"endpoint"`
	// CustomCA is used to connect to custom s3 endpoints
	CustomCA string `yaml:"custom_ca" json:"customCa,omitempty"`
	// Folder to place the files
	Folder string `yaml:"folder" json:"folder,omitempty"`
}

type EtcdBackupSpec struct {
	// cluster ID
	ClusterID string `json:"clusterId,omitempty" bhojpur:"required,type=reference[cluster],noupdate"`
	// manual backup flag
	Manual bool `yaml:"manual" json:"manual,omitempty"`
	// actual file name on the target
	Filename string `yaml:"filename" json:"filename,omitempty" bhojpur:"noupdate"`
	// backupConfig
	BackupConfig BackupConfig `yaml:",omitempty" json:"backupConfig,omitempty" bhojpur:"noupdate"`
}

type EtcdBackupStatus struct {
	Conditions []EtcdBackupCondition `json:"conditions"`
	// version of k8s in the backup pulled from BKE config
	KubernetesVersion string `yaml:"kubernetesVersion" json:"kubernetesVersion,omitempty" bhojpur:"noupdate"`
	// json + gzipped + base64 backup of the cluster object when the backup was created
	ClusterObject string `yaml:"clusterObject" json:"clusterObject,omitempty" bhojpur:"type=password,noupdate"`
}

type EtcdBackupCondition struct {
	// Type of condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}
