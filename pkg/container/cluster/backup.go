package cluster

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
	"context"
	"encoding/json"

	"github.com/bhojpur/host/pkg/container/types"
)

func (c *Cluster) ETCDSave(ctx context.Context, snapshotName string) error {
	driverOpts, err := c.getDriverOps()
	if err != nil {
		return err
	}
	return c.Driver.ETCDSave(ctx, toInfo(c), driverOpts, snapshotName)
}

func (c *Cluster) ETCDRestore(ctx context.Context, snapshotName string) error {
	driverOpts, err := c.getDriverOps()
	if err != nil {
		return err
	}
	if err := c.PersistStore.PersistStatus(*c, Updating); err != nil {
		return err
	}
	info, err := c.Driver.ETCDRestore(ctx, toInfo(c), driverOpts, snapshotName)
	if err != nil {
		return err
	}

	transformClusterInfo(c, info)

	return c.PostCheck(ctx)
}

func (c *Cluster) getDriverOps() (*types.DriverOptions, error) {
	if err := c.restore(); err != nil {
		return nil, err
	}
	driverOpts, err := c.ConfigGetter.GetConfig()
	if err != nil {
		return nil, err
	}

	driverOpts.StringOptions["name"] = c.Name

	for k, v := range c.Metadata {
		if k == "state" {
			state := make(map[string]interface{})
			if err := json.Unmarshal([]byte(v), &state); err == nil {
				flattenIfNotExist(state, &driverOpts)
			}

			continue
		}

		driverOpts.StringOptions[k] = v
	}

	return &driverOpts, nil
}

func (c *Cluster) ETCDRemoveSnapshot(ctx context.Context, snapshotName string) error {
	driverOpts, err := c.getDriverOps()
	if err != nil {
		return err
	}
	return c.Driver.ETCDRemoveSnapshot(ctx, toInfo(c), driverOpts, snapshotName)
}
