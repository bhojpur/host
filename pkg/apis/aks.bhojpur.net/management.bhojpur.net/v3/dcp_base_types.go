package v3

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

//DcpConfig provides desired configuration for DCP clusters
type DcpConfig struct {
	Version                string `yaml:"kubernetes_version" json:"kubernetesVersion,omitempty"`
	ClusterUpgradeStrategy `yaml:"dcp_upgrade_strategy,omitempty" json:"dcpupgradeStrategy,omitempty"`
}
type UkeConfig struct {
	Version                string `yaml:"kubernetes_version" json:"kubernetesVersion,omitempty"`
	ClusterUpgradeStrategy `yaml:"uke_upgrade_strategy,omitempty" json:"ukeupgradeStrategy,omitempty"`
}

//ClusterUpgradeStrategy provides configuration to the downstream system-upgrade-controller
type ClusterUpgradeStrategy struct {
	// How many controlplane nodes should be upgrade at time, defaults to 1
	ServerConcurrency int `yaml:"server_concurrency" json:"serverConcurrency,omitempty" bhojpur:"min=1"`
	// How many workers should be upgraded at a time
	WorkerConcurrency int `yaml:"worker_concurrency" json:"workerConcurrency,omitempty" bhojpur:"min=1"`
	// Whether controlplane nodes should be drained
	DrainServerNodes bool `yaml:"drain_server_nodes" json:"drainServerNodes,omitempty"`
	// Whether worker nodes should be drained
	DrainWorkerNodes bool `yaml:"drain_worker_nodes" json:"drainWorkerNodes,omitempty"`
}

func (r *UkeConfig) SetStrategy(serverConcurrency, workerConcurrency int) {
	r.ClusterUpgradeStrategy.ServerConcurrency = serverConcurrency
	r.ClusterUpgradeStrategy.WorkerConcurrency = workerConcurrency
}
func (k *DcpConfig) SetStrategy(serverConcurrency, workerConcurrency int) {
	k.ClusterUpgradeStrategy.ServerConcurrency = serverConcurrency
	k.ClusterUpgradeStrategy.WorkerConcurrency = workerConcurrency
}
