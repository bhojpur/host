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
	"context"

	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewClient creates a gRPC client for a driver plugin
func NewClient(driverName string, addr string) (CloseableDriver, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewDriverClient(conn)
	return &grpcClient{
		client:     c,
		driverName: driverName,
		conn:       conn,
	}, nil
}

// grpcClient defines the gRPC client struct
type grpcClient struct {
	client     DriverClient
	driverName string
	conn       *grpc.ClientConn
}

// Create call grpc create
func (rpc *grpcClient) Create(ctx context.Context, opts *DriverOptions, clusterInfo *ClusterInfo) (*ClusterInfo, error) {
	o, err := rpc.client.Create(ctx, &CreateRequest{
		DriverOptions: opts,
		ClusterInfo:   clusterInfo,
	})
	err = handlErr(err)
	if err == nil && o.CreateError != "" {
		err = errors.New(o.CreateError)
	}
	return o, err
}

// Update call grpc update
func (rpc *grpcClient) Update(ctx context.Context, clusterInfo *ClusterInfo, opts *DriverOptions) (*ClusterInfo, error) {
	o, err := rpc.client.Update(ctx, &UpdateRequest{
		ClusterInfo:   clusterInfo,
		DriverOptions: opts,
	})
	return o, handlErr(err)
}

func (rpc *grpcClient) PostCheck(ctx context.Context, clusterInfo *ClusterInfo) (*ClusterInfo, error) {
	o, err := rpc.client.PostCheck(ctx, clusterInfo)
	return o, handlErr(err)
}

// Remove call grpc remove
func (rpc *grpcClient) Remove(ctx context.Context, clusterInfo *ClusterInfo) error {
	_, err := rpc.client.Remove(ctx, clusterInfo)
	return handlErr(err)
}

// GetDriverCreateOptions call grpc getDriverCreateOptions
func (rpc *grpcClient) GetDriverCreateOptions(ctx context.Context) (*DriverFlags, error) {
	o, err := rpc.client.GetDriverCreateOptions(ctx, &Empty{})
	return o, handlErr(err)
}

// GetDriverUpdateOptions call grpc getDriverUpdateOptions
func (rpc *grpcClient) GetDriverUpdateOptions(ctx context.Context) (*DriverFlags, error) {
	o, err := rpc.client.GetDriverUpdateOptions(ctx, &Empty{})
	return o, handlErr(err)
}

func (rpc *grpcClient) GetVersion(ctx context.Context, info *ClusterInfo) (*KubernetesVersion, error) {
	version, err := rpc.client.GetVersion(ctx, info)
	return version, handlErr(err)
}

func (rpc *grpcClient) SetVersion(ctx context.Context, info *ClusterInfo, version *KubernetesVersion) error {
	_, err := rpc.client.SetVersion(ctx, &SetVersionRequest{Info: info, Version: version})
	return handlErr(err)
}

func (rpc *grpcClient) GetClusterSize(ctx context.Context, info *ClusterInfo) (*NodeCount, error) {
	size, err := rpc.client.GetNodeCount(ctx, info)
	return size, handlErr(err)
}

func (rpc *grpcClient) SetClusterSize(ctx context.Context, info *ClusterInfo, count *NodeCount) error {
	_, err := rpc.client.SetNodeCount(ctx, &SetNodeCountRequest{Info: info, Count: count})
	return handlErr(err)
}

func (rpc *grpcClient) GetCapabilities(ctx context.Context) (*Capabilities, error) {
	return rpc.client.GetCapabilities(ctx, &Empty{})
}

func (rpc *grpcClient) GetK8SCapabilities(ctx context.Context, opts *DriverOptions) (*K8SCapabilities, error) {
	capabilities, err := rpc.client.GetK8SCapabilities(ctx, opts)
	return capabilities, handlErr(err)
}

func (rpc *grpcClient) Close() error {
	return rpc.conn.Close()
}

func (rpc *grpcClient) ETCDSave(ctx context.Context, clusterInfo *ClusterInfo, opts *DriverOptions, snapshotName string) error {
	_, err := rpc.client.ETCDSave(ctx, &SaveETCDSnapshotRequest{Info: clusterInfo, SnapshotName: snapshotName, DriverOptions: opts})
	return handlErr(err)
}

func (rpc *grpcClient) ETCDRestore(ctx context.Context, clusterInfo *ClusterInfo, opts *DriverOptions, snapshotName string) (*ClusterInfo, error) {
	o, err := rpc.client.ETCDRestore(ctx, &RestoreETCDSnapshotRequest{Info: clusterInfo, SnapshotName: snapshotName, DriverOptions: opts})
	return o, handlErr(err)
}

func (rpc *grpcClient) ETCDRemoveSnapshot(ctx context.Context, clusterInfo *ClusterInfo, opts *DriverOptions, snapshotName string) error {
	_, err := rpc.client.ETCDRemoveSnapshot(ctx, &RemoveETCDSnapshotRequest{Info: clusterInfo, SnapshotName: snapshotName, DriverOptions: opts})
	return handlErr(err)
}

func (rpc *grpcClient) RemoveLegacyServiceAccount(ctx context.Context, info *ClusterInfo) error {
	_, err := rpc.client.RemoveLegacyServiceAccount(ctx, info)
	return handlErr(err)
}

func handlErr(err error) error {
	if st, ok := status.FromError(err); ok {
		if st.Code() == codes.Unknown && st.Message() != "" {
			return errors.New(st.Message())
		}
	}
	return err
}
