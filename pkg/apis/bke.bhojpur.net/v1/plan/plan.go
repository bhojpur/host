package plan

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

import capi "sigs.k8s.io/cluster-api/api/v1beta1"

type Plan struct {
	Nodes    map[string]*Node         `json:"nodes,omitempty"`
	Machines map[string]*capi.Machine `json:"machines,omitempty"`
	Metadata map[string]*Metadata     `json:"metadata,omitempty"`
	Cluster  *capi.Cluster            `json:"cluster,omitempty"`
}

type Metadata struct {
	Labels      map[string]string `json:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ProbeStatus struct {
	Healthy      bool `json:"healthy,omitempty"`
	SuccessCount int  `json:"successCount,omitempty"`
	FailureCount int  `json:"failureCount,omitempty"`
}

type Node struct {
	Plan           NodePlan                             `json:"plan,omitempty"`
	AppliedPlan    *NodePlan                            `json:"appliedPlan,omitempty"`
	Output         map[string][]byte                    `json:"-"`
	PeriodicOutput map[string]PeriodicInstructionOutput `json:"-"`
	Failed         bool                                 `json:"failed,omitempty"`
	InSync         bool                                 `json:"inSync,omitempty"`
	Healthy        bool                                 `json:"healthy,omitempty"`
	ProbeStatus    map[string]ProbeStatus               `json:"probeStatus,omitempty"`
}

type PeriodicInstructionOutput struct {
	Name                  string `json:"name"`
	Stdout                []byte `json:"stdout"`                // Stdout is a byte array of the gzip+base64 stdout output
	Stderr                []byte `json:"stderr"`                // Stderr is a byte array of the gzip+base64 stderr output
	ExitCode              int    `json:"exitCode"`              // ExitCode is an int representing the exit code of the last run instruction
	LastSuccessfulRunTime string `json:"lastSuccessfulRunTime"` // LastSuccessfulRunTime is a time.UnixDate formatted string of the last time the instruction was run
}

type Secret struct {
	ServerToken string `json:"serverToken,omitempty"`
	AgentToken  string `json:"agentToken,omitempty"`
}

type OneTimeInstruction struct {
	Name       string   `json:"name,omitempty"`
	Image      string   `json:"image,omitempty"`
	Env        []string `json:"env,omitempty"`
	Args       []string `json:"args,omitempty"`
	Command    string   `json:"command,omitempty"`
	SaveOutput bool     `json:"saveOutput,omitempty"`
}

type PeriodicInstruction struct {
	Name          string   `json:"name,omitempty"`
	Image         string   `json:"image,omitempty"`
	Env           []string `json:"env,omitempty"`
	Args          []string `json:"args,omitempty"`
	Command       string   `json:"command,omitempty"`
	PeriodSeconds int      `json:"periodSeconds,omitempty"` // default 600, i.e. 10 minutes
}

type File struct {
	Content string `json:"content,omitempty"`
	Path    string `json:"path,omitempty"`
	Dynamic bool   `json:"dynamic,omitempty"`
}

type NodePlan struct {
	Files                []File                `json:"files,omitempty"`
	Instructions         []OneTimeInstruction  `json:"instructions,omitempty"`
	PeriodicInstructions []PeriodicInstruction `json:"periodicInstructions,omitempty"`
	Error                string                `json:"error,omitempty"`
	Probes               map[string]Probe      `json:"probes,omitempty"`
}
