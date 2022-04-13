package client

import (
	"github.com/bhojpur/host/pkg/core/types"
)

const (
	ContainerDriverType                      = "containerDriver"
	ContainerDriverFieldActive               = "active"
	ContainerDriverFieldActualURL            = "actualUrl"
	ContainerDriverFieldAnnotations          = "annotations"
	ContainerDriverFieldBuiltIn              = "builtIn"
	ContainerDriverFieldChecksum             = "checksum"
	ContainerDriverFieldConditions           = "conditions"
	ContainerDriverFieldCreated              = "created"
	ContainerDriverFieldCreatorID            = "creatorId"
	ContainerDriverFieldExecutablePath       = "executablePath"
	ContainerDriverFieldLabels               = "labels"
	ContainerDriverFieldName                 = "name"
	ContainerDriverFieldOwnerReferences      = "ownerReferences"
	ContainerDriverFieldRemoved              = "removed"
	ContainerDriverFieldState                = "state"
	ContainerDriverFieldTransitioning        = "transitioning"
	ContainerDriverFieldTransitioningMessage = "transitioningMessage"
	ContainerDriverFieldUIURL                = "uiUrl"
	ContainerDriverFieldURL                  = "url"
	ContainerDriverFieldUUID                 = "uuid"
	ContainerDriverFieldWhitelistDomains     = "whitelistDomains"
)

type ContainerDriver struct {
	types.Resource
	Active               bool              `json:"active,omitempty" yaml:"active,omitempty"`
	ActualURL            string            `json:"actualUrl,omitempty" yaml:"actualUrl,omitempty"`
	Annotations          map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	BuiltIn              bool              `json:"builtIn,omitempty" yaml:"builtIn,omitempty"`
	Checksum             string            `json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Conditions           []Condition       `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created              string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	ExecutablePath       string            `json:"executablePath,omitempty" yaml:"executablePath,omitempty"`
	Labels               map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences      []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed              string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string            `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning        string            `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string            `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UIURL                string            `json:"uiUrl,omitempty" yaml:"uiUrl,omitempty"`
	URL                  string            `json:"url,omitempty" yaml:"url,omitempty"`
	UUID                 string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	WhitelistDomains     []string          `json:"whitelistDomains,omitempty" yaml:"whitelistDomains,omitempty"`
}

type ContainerDriverCollection struct {
	types.Collection
	Data   []ContainerDriver `json:"data,omitempty"`
	client *ContainerDriverClient
}

type ContainerDriverClient struct {
	apiClient *Client
}

type ContainerDriverOperations interface {
	List(opts *types.ListOpts) (*ContainerDriverCollection, error)
	ListAll(opts *types.ListOpts) (*ContainerDriverCollection, error)
	Create(opts *ContainerDriver) (*ContainerDriver, error)
	Update(existing *ContainerDriver, updates interface{}) (*ContainerDriver, error)
	Replace(existing *ContainerDriver) (*ContainerDriver, error)
	ByID(id string) (*ContainerDriver, error)
	Delete(container *ContainerDriver) error

	ActionActivate(resource *ContainerDriver) error

	ActionDeactivate(resource *ContainerDriver) error

	CollectionActionRefresh(resource *ContainerDriverCollection) error
}

func newContainerDriverClient(apiClient *Client) *ContainerDriverClient {
	return &ContainerDriverClient{
		apiClient: apiClient,
	}
}

func (c *ContainerDriverClient) Create(container *ContainerDriver) (*ContainerDriver, error) {
	resp := &ContainerDriver{}
	err := c.apiClient.Ops.DoCreate(ContainerDriverType, container, resp)
	return resp, err
}

func (c *ContainerDriverClient) Update(existing *ContainerDriver, updates interface{}) (*ContainerDriver, error) {
	resp := &ContainerDriver{}
	err := c.apiClient.Ops.DoUpdate(ContainerDriverType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ContainerDriverClient) Replace(obj *ContainerDriver) (*ContainerDriver, error) {
	resp := &ContainerDriver{}
	err := c.apiClient.Ops.DoReplace(ContainerDriverType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ContainerDriverClient) List(opts *types.ListOpts) (*ContainerDriverCollection, error) {
	resp := &ContainerDriverCollection{}
	err := c.apiClient.Ops.DoList(ContainerDriverType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *ContainerDriverClient) ListAll(opts *types.ListOpts) (*ContainerDriverCollection, error) {
	resp := &ContainerDriverCollection{}
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

func (cc *ContainerDriverCollection) Next() (*ContainerDriverCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ContainerDriverCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ContainerDriverClient) ByID(id string) (*ContainerDriver, error) {
	resp := &ContainerDriver{}
	err := c.apiClient.Ops.DoByID(ContainerDriverType, id, resp)
	return resp, err
}

func (c *ContainerDriverClient) Delete(container *ContainerDriver) error {
	return c.apiClient.Ops.DoResourceDelete(ContainerDriverType, &container.Resource)
}

func (c *ContainerDriverClient) ActionActivate(resource *ContainerDriver) error {
	err := c.apiClient.Ops.DoAction(ContainerDriverType, "activate", &resource.Resource, nil, nil)
	return err
}

func (c *ContainerDriverClient) ActionDeactivate(resource *ContainerDriver) error {
	err := c.apiClient.Ops.DoAction(ContainerDriverType, "deactivate", &resource.Resource, nil, nil)
	return err
}

func (c *ContainerDriverClient) CollectionActionRefresh(resource *ContainerDriverCollection) error {
	err := c.apiClient.Ops.DoCollectionAction(ContainerDriverType, "refresh", &resource.Collection, nil, nil)
	return err
}
