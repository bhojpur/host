package client

import (
	"github.com/bhojpur/host/pkg/core/types"
)

const (
	BkeK8sServiceOptionType                 = "bkeK8sServiceOption"
	BkeK8sServiceOptionFieldAnnotations     = "annotations"
	BkeK8sServiceOptionFieldCreated         = "created"
	BkeK8sServiceOptionFieldCreatorID       = "creatorId"
	BkeK8sServiceOptionFieldLabels          = "labels"
	BkeK8sServiceOptionFieldName            = "name"
	BkeK8sServiceOptionFieldOwnerReferences = "ownerReferences"
	BkeK8sServiceOptionFieldRemoved         = "removed"
	BkeK8sServiceOptionFieldServiceOptions  = "serviceOptions"
	BkeK8sServiceOptionFieldUUID            = "uuid"
)

type BkeK8sServiceOption struct {
	types.Resource
	Annotations     map[string]string          `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string                     `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string                     `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string          `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string                     `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference           `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string                     `json:"removed,omitempty" yaml:"removed,omitempty"`
	ServiceOptions  *KubernetesServicesOptions `json:"serviceOptions,omitempty" yaml:"serviceOptions,omitempty"`
	UUID            string                     `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type BkeK8sServiceOptionCollection struct {
	types.Collection
	Data   []BkeK8sServiceOption `json:"data,omitempty"`
	client *BkeK8sServiceOptionClient
}

type BkeK8sServiceOptionClient struct {
	apiClient *Client
}

type BkeK8sServiceOptionOperations interface {
	List(opts *types.ListOpts) (*BkeK8sServiceOptionCollection, error)
	ListAll(opts *types.ListOpts) (*BkeK8sServiceOptionCollection, error)
	Create(opts *BkeK8sServiceOption) (*BkeK8sServiceOption, error)
	Update(existing *BkeK8sServiceOption, updates interface{}) (*BkeK8sServiceOption, error)
	Replace(existing *BkeK8sServiceOption) (*BkeK8sServiceOption, error)
	ByID(id string) (*BkeK8sServiceOption, error)
	Delete(container *BkeK8sServiceOption) error
}

func newBkeK8sServiceOptionClient(apiClient *Client) *BkeK8sServiceOptionClient {
	return &BkeK8sServiceOptionClient{
		apiClient: apiClient,
	}
}

func (c *BkeK8sServiceOptionClient) Create(container *BkeK8sServiceOption) (*BkeK8sServiceOption, error) {
	resp := &BkeK8sServiceOption{}
	err := c.apiClient.Ops.DoCreate(BkeK8sServiceOptionType, container, resp)
	return resp, err
}

func (c *BkeK8sServiceOptionClient) Update(existing *BkeK8sServiceOption, updates interface{}) (*BkeK8sServiceOption, error) {
	resp := &BkeK8sServiceOption{}
	err := c.apiClient.Ops.DoUpdate(BkeK8sServiceOptionType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BkeK8sServiceOptionClient) Replace(obj *BkeK8sServiceOption) (*BkeK8sServiceOption, error) {
	resp := &BkeK8sServiceOption{}
	err := c.apiClient.Ops.DoReplace(BkeK8sServiceOptionType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *BkeK8sServiceOptionClient) List(opts *types.ListOpts) (*BkeK8sServiceOptionCollection, error) {
	resp := &BkeK8sServiceOptionCollection{}
	err := c.apiClient.Ops.DoList(BkeK8sServiceOptionType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *BkeK8sServiceOptionClient) ListAll(opts *types.ListOpts) (*BkeK8sServiceOptionCollection, error) {
	resp := &BkeK8sServiceOptionCollection{}
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

func (cc *BkeK8sServiceOptionCollection) Next() (*BkeK8sServiceOptionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BkeK8sServiceOptionCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BkeK8sServiceOptionClient) ByID(id string) (*BkeK8sServiceOption, error) {
	resp := &BkeK8sServiceOption{}
	err := c.apiClient.Ops.DoByID(BkeK8sServiceOptionType, id, resp)
	return resp, err
}

func (c *BkeK8sServiceOptionClient) Delete(container *BkeK8sServiceOption) error {
	return c.apiClient.Ops.DoResourceDelete(BkeK8sServiceOptionType, &container.Resource)
}
