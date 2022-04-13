package client

import (
	"github.com/bhojpur/host/pkg/core/types"
)

const (
	BkeK8sSystemImageType                 = "bkeK8sSystemImage"
	BkeK8sSystemImageFieldAnnotations     = "annotations"
	BkeK8sSystemImageFieldCreated         = "created"
	BkeK8sSystemImageFieldCreatorID       = "creatorId"
	BkeK8sSystemImageFieldLabels          = "labels"
	BkeK8sSystemImageFieldName            = "name"
	BkeK8sSystemImageFieldOwnerReferences = "ownerReferences"
	BkeK8sSystemImageFieldRemoved         = "removed"
	BkeK8sSystemImageFieldSystemImages    = "systemImages"
	BkeK8sSystemImageFieldUUID            = "uuid"
)

type BkeK8sSystemImage struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	SystemImages    *BKESystemImages  `json:"systemImages,omitempty" yaml:"systemImages,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type BkeK8sSystemImageCollection struct {
	types.Collection
	Data   []BkeK8sSystemImage `json:"data,omitempty"`
	client *BkeK8sSystemImageClient
}

type BkeK8sSystemImageClient struct {
	apiClient *Client
}

type BkeK8sSystemImageOperations interface {
	List(opts *types.ListOpts) (*BkeK8sSystemImageCollection, error)
	ListAll(opts *types.ListOpts) (*BkeK8sSystemImageCollection, error)
	Create(opts *BkeK8sSystemImage) (*BkeK8sSystemImage, error)
	Update(existing *BkeK8sSystemImage, updates interface{}) (*BkeK8sSystemImage, error)
	Replace(existing *BkeK8sSystemImage) (*BkeK8sSystemImage, error)
	ByID(id string) (*BkeK8sSystemImage, error)
	Delete(container *BkeK8sSystemImage) error
}

func newBkeK8sSystemImageClient(apiClient *Client) *BkeK8sSystemImageClient {
	return &BkeK8sSystemImageClient{
		apiClient: apiClient,
	}
}

func (c *BkeK8sSystemImageClient) Create(container *BkeK8sSystemImage) (*BkeK8sSystemImage, error) {
	resp := &BkeK8sSystemImage{}
	err := c.apiClient.Ops.DoCreate(BkeK8sSystemImageType, container, resp)
	return resp, err
}

func (c *BkeK8sSystemImageClient) Update(existing *BkeK8sSystemImage, updates interface{}) (*BkeK8sSystemImage, error) {
	resp := &BkeK8sSystemImage{}
	err := c.apiClient.Ops.DoUpdate(BkeK8sSystemImageType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BkeK8sSystemImageClient) Replace(obj *BkeK8sSystemImage) (*BkeK8sSystemImage, error) {
	resp := &BkeK8sSystemImage{}
	err := c.apiClient.Ops.DoReplace(BkeK8sSystemImageType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *BkeK8sSystemImageClient) List(opts *types.ListOpts) (*BkeK8sSystemImageCollection, error) {
	resp := &BkeK8sSystemImageCollection{}
	err := c.apiClient.Ops.DoList(BkeK8sSystemImageType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *BkeK8sSystemImageClient) ListAll(opts *types.ListOpts) (*BkeK8sSystemImageCollection, error) {
	resp := &BkeK8sSystemImageCollection{}
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

func (cc *BkeK8sSystemImageCollection) Next() (*BkeK8sSystemImageCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BkeK8sSystemImageCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BkeK8sSystemImageClient) ByID(id string) (*BkeK8sSystemImage, error) {
	resp := &BkeK8sSystemImage{}
	err := c.apiClient.Ops.DoByID(BkeK8sSystemImageType, id, resp)
	return resp, err
}

func (c *BkeK8sSystemImageClient) Delete(container *BkeK8sSystemImage) error {
	return c.apiClient.Ops.DoResourceDelete(BkeK8sSystemImageType, &container.Resource)
}
