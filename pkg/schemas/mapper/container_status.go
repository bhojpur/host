package mapper

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

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"github.com/bhojpur/host/pkg/core/types/values"
)

type ContainerStatus struct {
}

type containerState struct {
	state         string
	message       string
	transitioning string
	exitCode      interface{}
	restartCount  int64
}

func message(m map[string]interface{}) string {
	if m["message"] == nil {
		return convert.ToString(m["reason"])
	}
	return fmt.Sprintf("%s: %s", m["reason"], m["message"])
}

func checkStatus(containerStates map[string]containerState, containerStatus []map[string]interface{}) {
	for _, status := range containerStatus {
		name := convert.ToString(status["name"])
		restartCount, _ := convert.ToNumber(status["restartCount"])
		s := containerState{
			state:         "pending",
			restartCount:  restartCount,
			transitioning: "no",
		}
		for k, v := range convert.ToMapInterface(status["state"]) {
			m := convert.ToMapInterface(v)
			switch k {
			case "terminated":
				s.state = "terminated"
				s.message = message(m)
				s.exitCode = m["exitCode"]
				if convert.ToString(s.exitCode) == "0" {
					s.transitioning = "no"
				} else {
					s.transitioning = "error"
				}
			case "running":
				ready := convert.ToBool(status["ready"])
				if ready {
					s.state = "running"
					s.transitioning = "no"
				} else {
					s.state = "notready"
					s.transitioning = "yes"
				}
			case "waiting":
				s.state = "waiting"
				s.transitioning = "yes"
				s.message = message(m)
			}
		}

		containerStates[name] = s
	}
}

func (n ContainerStatus) FromInternal(data map[string]interface{}) {
	containerStates := map[string]containerState{}
	containerStatus := convert.ToMapSlice(values.GetValueN(data, "status", "containerStatuses"))
	checkStatus(containerStates, containerStatus)

	containerStatus = convert.ToMapSlice(values.GetValueN(data, "status", "initContainerStatuses"))
	checkStatus(containerStates, containerStatus)

	containers := append(convert.ToMapSlice(values.GetValueN(data, "containers")),
		convert.ToMapSlice(values.GetValueN(data, "initContainers"))...)
	for _, container := range containers {
		if container == nil {
			continue
		}

		name := convert.ToString(container["name"])
		state, ok := containerStates[name]
		if ok {
			container["state"] = state.state
			container["transitioning"] = state.transitioning
			container["transitioningMessage"] = state.message
			container["restartCount"] = state.restartCount
			container["exitCode"] = state.exitCode
		} else {
			container["state"] = "unknown"
		}
	}
}

func (n ContainerStatus) ToInternal(data map[string]interface{}) error {
	return nil
}

func (n ContainerStatus) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
