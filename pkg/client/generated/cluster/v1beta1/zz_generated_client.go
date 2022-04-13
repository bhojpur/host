package client

import (
	"github.com/bhojpur/host/pkg/core/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Machine MachineOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Machine = newMachineClient(client)

	return client, nil
}