package v1

import (
	"github.com/bhojpur/host/pkg/core/objectclient"
	"github.com/bhojpur/host/pkg/labni/client"
	"github.com/bhojpur/host/pkg/labni/controller"
	"k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

type Interface interface {
	ClusterRoleBindingsGetter
	ClusterRolesGetter
	RoleBindingsGetter
	RolesGetter
}

type Client struct {
	controllerFactory controller.SharedControllerFactory
	clientFactory     client.SharedClientFactory
}

func NewForConfig(cfg rest.Config) (Interface, error) {
	scheme := runtime.NewScheme()
	if err := v1.AddToScheme(scheme); err != nil {
		return nil, err
	}
	controllerFactory, err := controller.NewSharedControllerFactoryFromConfig(&cfg, scheme)
	if err != nil {
		return nil, err
	}
	return NewFromControllerFactory(controllerFactory)
}

func NewFromControllerFactory(factory controller.SharedControllerFactory) (Interface, error) {
	return &Client{
		controllerFactory: factory,
		clientFactory:     factory.SharedCacheFactory().SharedClientFactory(),
	}, nil
}

type ClusterRoleBindingsGetter interface {
	ClusterRoleBindings(namespace string) ClusterRoleBindingInterface
}

func (c *Client) ClusterRoleBindings(namespace string) ClusterRoleBindingInterface {
	sharedClient := c.clientFactory.ForResourceKind(ClusterRoleBindingGroupVersionResource, ClusterRoleBindingGroupVersionKind.Kind, false)
	objectClient := objectclient.NewObjectClient(namespace, sharedClient, &ClusterRoleBindingResource, ClusterRoleBindingGroupVersionKind, clusterRoleBindingFactory{})
	return &clusterRoleBindingClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type ClusterRolesGetter interface {
	ClusterRoles(namespace string) ClusterRoleInterface
}

func (c *Client) ClusterRoles(namespace string) ClusterRoleInterface {
	sharedClient := c.clientFactory.ForResourceKind(ClusterRoleGroupVersionResource, ClusterRoleGroupVersionKind.Kind, false)
	objectClient := objectclient.NewObjectClient(namespace, sharedClient, &ClusterRoleResource, ClusterRoleGroupVersionKind, clusterRoleFactory{})
	return &clusterRoleClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type RoleBindingsGetter interface {
	RoleBindings(namespace string) RoleBindingInterface
}

func (c *Client) RoleBindings(namespace string) RoleBindingInterface {
	sharedClient := c.clientFactory.ForResourceKind(RoleBindingGroupVersionResource, RoleBindingGroupVersionKind.Kind, true)
	objectClient := objectclient.NewObjectClient(namespace, sharedClient, &RoleBindingResource, RoleBindingGroupVersionKind, roleBindingFactory{})
	return &roleBindingClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}

type RolesGetter interface {
	Roles(namespace string) RoleInterface
}

func (c *Client) Roles(namespace string) RoleInterface {
	sharedClient := c.clientFactory.ForResourceKind(RoleGroupVersionResource, RoleGroupVersionKind.Kind, true)
	objectClient := objectclient.NewObjectClient(namespace, sharedClient, &RoleResource, RoleGroupVersionKind, roleFactory{})
	return &roleClient{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
