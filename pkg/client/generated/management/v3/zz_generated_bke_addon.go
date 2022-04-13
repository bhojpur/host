package client

import (
	"github.com/bhojpur/host/pkg/core/types"
)

const (
	BkeAddonType                 = "bkeAddon"
	BkeAddonFieldAnnotations     = "annotations"
	BkeAddonFieldCreated         = "created"
	BkeAddonFieldCreatorID       = "creatorId"
	BkeAddonFieldLabels          = "labels"
	BkeAddonFieldName            = "name"
	BkeAddonFieldOwnerReferences = "ownerReferences"
	BkeAddonFieldRemoved         = "removed"
	BkeAddonFieldTemplate        = "template"
	BkeAddonFieldUUID            = "uuid"
)

type BkeAddon struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Template        string            `json:"template,omitempty" yaml:"template,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type BkeAddonCollection struct {
	types.Collection
	Data   []BkeAddon `json:"data,omitempty"`
	client *BkeAddonClient
}

type BkeAddonClient struct {
	apiClient *Client
}

type BkeAddonOperations interface {
	List(opts *types.ListOpts) (*BkeAddonCollection, error)
	ListAll(opts *types.ListOpts) (*BkeAddonCollection, error)
	Create(opts *BkeAddon) (*BkeAddon, error)
	Update(existing *BkeAddon, updates interface{}) (*BkeAddon, error)
	Replace(existing *BkeAddon) (*BkeAddon, error)
	ByID(id string) (*BkeAddon, error)
	Delete(container *BkeAddon) error
}

func newBkeAddonClient(apiClient *Client) *BkeAddonClient {
	return &BkeAddonClient{
		apiClient: apiClient,
	}
}

func (c *BkeAddonClient) Create(container *BkeAddon) (*BkeAddon, error) {
	resp := &BkeAddon{}
	err := c.apiClient.Ops.DoCreate(BkeAddonType, container, resp)
	return resp, err
}

func (c *BkeAddonClient) Update(existing *BkeAddon, updates interface{}) (*BkeAddon, error) {
	resp := &BkeAddon{}
	err := c.apiClient.Ops.DoUpdate(BkeAddonType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BkeAddonClient) Replace(obj *BkeAddon) (*BkeAddon, error) {
	resp := &BkeAddon{}
	err := c.apiClient.Ops.DoReplace(BkeAddonType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *BkeAddonClient) List(opts *types.ListOpts) (*BkeAddonCollection, error) {
	resp := &BkeAddonCollection{}
	err := c.apiClient.Ops.DoList(BkeAddonType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *BkeAddonClient) ListAll(opts *types.ListOpts) (*BkeAddonCollection, error) {
	resp := &BkeAddonCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *BkeAddonCollection) Next() (*BkeAddonCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BkeAddonCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BkeAddonClient) ByID(id string) (*BkeAddon, error) {
	resp := &BkeAddon{}
	err := c.apiClient.Ops.DoByID(BkeAddonType, id, resp)
	return resp, err
}

func (c *BkeAddonClient) Delete(container *BkeAddon) error {
	return c.apiClient.Ops.DoResourceDelete(BkeAddonType, &container.Resource)
}
