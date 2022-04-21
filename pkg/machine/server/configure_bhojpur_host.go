// This file is safe to edit. Once it exists it will not be overwritten

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package server

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/bhojpur/host/pkg/machine/server/operations"
	"github.com/bhojpur/host/pkg/machine/server/operations/config"
	"github.com/bhojpur/host/pkg/machine/server/operations/container"
	"github.com/bhojpur/host/pkg/machine/server/operations/distribution"
	"github.com/bhojpur/host/pkg/machine/server/operations/exec"
	"github.com/bhojpur/host/pkg/machine/server/operations/image"
	"github.com/bhojpur/host/pkg/machine/server/operations/network"
	"github.com/bhojpur/host/pkg/machine/server/operations/node"
	"github.com/bhojpur/host/pkg/machine/server/operations/plugin"
	"github.com/bhojpur/host/pkg/machine/server/operations/secret"
	"github.com/bhojpur/host/pkg/machine/server/operations/service"
	"github.com/bhojpur/host/pkg/machine/server/operations/session"
	"github.com/bhojpur/host/pkg/machine/server/operations/swarm"
	"github.com/bhojpur/host/pkg/machine/server/operations/system"
	"github.com/bhojpur/host/pkg/machine/server/operations/task"
	"github.com/bhojpur/host/pkg/machine/server/operations/volume"
)

//go:generate swagger generate server --target ../../../../host --name BhojpurHost --spec ../../swagger/swagger.yml --model-package pkg/machine/models --server-package pkg/machine/server --principal interface{}

func configureFlags(api *operations.BhojpurHostAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BhojpurHostAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.BinConsumer = runtime.ByteStreamConsumer()
	api.JSONConsumer = runtime.JSONConsumer()
	api.TarConsumer = runtime.ConsumerFunc(func(r io.Reader, target interface{}) error {
		return errors.NotImplemented("tar consumer has not yet been implemented")
	})
	api.TxtConsumer = runtime.TextConsumer()

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.TarProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("tar producer has not yet been implemented")
	})
	api.TxtProducer = runtime.TextProducer()

	if api.ImageBuildPruneHandler == nil {
		api.ImageBuildPruneHandler = image.BuildPruneHandlerFunc(func(params image.BuildPruneParams) middleware.Responder {
			return middleware.NotImplemented("operation image.BuildPrune has not yet been implemented")
		})
	}
	if api.ConfigConfigCreateHandler == nil {
		api.ConfigConfigCreateHandler = config.ConfigCreateHandlerFunc(func(params config.ConfigCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation config.ConfigCreate has not yet been implemented")
		})
	}
	if api.ConfigConfigDeleteHandler == nil {
		api.ConfigConfigDeleteHandler = config.ConfigDeleteHandlerFunc(func(params config.ConfigDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation config.ConfigDelete has not yet been implemented")
		})
	}
	if api.ConfigConfigInspectHandler == nil {
		api.ConfigConfigInspectHandler = config.ConfigInspectHandlerFunc(func(params config.ConfigInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation config.ConfigInspect has not yet been implemented")
		})
	}
	if api.ConfigConfigListHandler == nil {
		api.ConfigConfigListHandler = config.ConfigListHandlerFunc(func(params config.ConfigListParams) middleware.Responder {
			return middleware.NotImplemented("operation config.ConfigList has not yet been implemented")
		})
	}
	if api.ConfigConfigUpdateHandler == nil {
		api.ConfigConfigUpdateHandler = config.ConfigUpdateHandlerFunc(func(params config.ConfigUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation config.ConfigUpdate has not yet been implemented")
		})
	}
	if api.ContainerContainerArchiveHandler == nil {
		api.ContainerContainerArchiveHandler = container.ContainerArchiveHandlerFunc(func(params container.ContainerArchiveParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerArchive has not yet been implemented")
		})
	}
	if api.ContainerContainerArchiveInfoHandler == nil {
		api.ContainerContainerArchiveInfoHandler = container.ContainerArchiveInfoHandlerFunc(func(params container.ContainerArchiveInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerArchiveInfo has not yet been implemented")
		})
	}
	if api.ContainerContainerAttachHandler == nil {
		api.ContainerContainerAttachHandler = container.ContainerAttachHandlerFunc(func(params container.ContainerAttachParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerAttach has not yet been implemented")
		})
	}
	if api.ContainerContainerAttachWebsocketHandler == nil {
		api.ContainerContainerAttachWebsocketHandler = container.ContainerAttachWebsocketHandlerFunc(func(params container.ContainerAttachWebsocketParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerAttachWebsocket has not yet been implemented")
		})
	}
	if api.ContainerContainerChangesHandler == nil {
		api.ContainerContainerChangesHandler = container.ContainerChangesHandlerFunc(func(params container.ContainerChangesParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerChanges has not yet been implemented")
		})
	}
	if api.ContainerContainerCreateHandler == nil {
		api.ContainerContainerCreateHandler = container.ContainerCreateHandlerFunc(func(params container.ContainerCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerCreate has not yet been implemented")
		})
	}
	if api.ContainerContainerDeleteHandler == nil {
		api.ContainerContainerDeleteHandler = container.ContainerDeleteHandlerFunc(func(params container.ContainerDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerDelete has not yet been implemented")
		})
	}
	if api.ExecContainerExecHandler == nil {
		api.ExecContainerExecHandler = exec.ContainerExecHandlerFunc(func(params exec.ContainerExecParams) middleware.Responder {
			return middleware.NotImplemented("operation exec.ContainerExec has not yet been implemented")
		})
	}
	if api.ContainerContainerExportHandler == nil {
		api.ContainerContainerExportHandler = container.ContainerExportHandlerFunc(func(params container.ContainerExportParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerExport has not yet been implemented")
		})
	}
	if api.ContainerContainerInspectHandler == nil {
		api.ContainerContainerInspectHandler = container.ContainerInspectHandlerFunc(func(params container.ContainerInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerInspect has not yet been implemented")
		})
	}
	if api.ContainerContainerKillHandler == nil {
		api.ContainerContainerKillHandler = container.ContainerKillHandlerFunc(func(params container.ContainerKillParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerKill has not yet been implemented")
		})
	}
	if api.ContainerContainerListHandler == nil {
		api.ContainerContainerListHandler = container.ContainerListHandlerFunc(func(params container.ContainerListParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerList has not yet been implemented")
		})
	}
	if api.ContainerContainerLogsHandler == nil {
		api.ContainerContainerLogsHandler = container.ContainerLogsHandlerFunc(func(params container.ContainerLogsParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerLogs has not yet been implemented")
		})
	}
	if api.ContainerContainerPauseHandler == nil {
		api.ContainerContainerPauseHandler = container.ContainerPauseHandlerFunc(func(params container.ContainerPauseParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerPause has not yet been implemented")
		})
	}
	if api.ContainerContainerPruneHandler == nil {
		api.ContainerContainerPruneHandler = container.ContainerPruneHandlerFunc(func(params container.ContainerPruneParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerPrune has not yet been implemented")
		})
	}
	if api.ContainerContainerRenameHandler == nil {
		api.ContainerContainerRenameHandler = container.ContainerRenameHandlerFunc(func(params container.ContainerRenameParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerRename has not yet been implemented")
		})
	}
	if api.ContainerContainerResizeHandler == nil {
		api.ContainerContainerResizeHandler = container.ContainerResizeHandlerFunc(func(params container.ContainerResizeParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerResize has not yet been implemented")
		})
	}
	if api.ContainerContainerRestartHandler == nil {
		api.ContainerContainerRestartHandler = container.ContainerRestartHandlerFunc(func(params container.ContainerRestartParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerRestart has not yet been implemented")
		})
	}
	if api.ContainerContainerStartHandler == nil {
		api.ContainerContainerStartHandler = container.ContainerStartHandlerFunc(func(params container.ContainerStartParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerStart has not yet been implemented")
		})
	}
	if api.ContainerContainerStatsHandler == nil {
		api.ContainerContainerStatsHandler = container.ContainerStatsHandlerFunc(func(params container.ContainerStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerStats has not yet been implemented")
		})
	}
	if api.ContainerContainerStopHandler == nil {
		api.ContainerContainerStopHandler = container.ContainerStopHandlerFunc(func(params container.ContainerStopParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerStop has not yet been implemented")
		})
	}
	if api.ContainerContainerTopHandler == nil {
		api.ContainerContainerTopHandler = container.ContainerTopHandlerFunc(func(params container.ContainerTopParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerTop has not yet been implemented")
		})
	}
	if api.ContainerContainerUnpauseHandler == nil {
		api.ContainerContainerUnpauseHandler = container.ContainerUnpauseHandlerFunc(func(params container.ContainerUnpauseParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerUnpause has not yet been implemented")
		})
	}
	if api.ContainerContainerUpdateHandler == nil {
		api.ContainerContainerUpdateHandler = container.ContainerUpdateHandlerFunc(func(params container.ContainerUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerUpdate has not yet been implemented")
		})
	}
	if api.ContainerContainerWaitHandler == nil {
		api.ContainerContainerWaitHandler = container.ContainerWaitHandlerFunc(func(params container.ContainerWaitParams) middleware.Responder {
			return middleware.NotImplemented("operation container.ContainerWait has not yet been implemented")
		})
	}
	if api.DistributionDistributionInspectHandler == nil {
		api.DistributionDistributionInspectHandler = distribution.DistributionInspectHandlerFunc(func(params distribution.DistributionInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation distribution.DistributionInspect has not yet been implemented")
		})
	}
	if api.ExecExecInspectHandler == nil {
		api.ExecExecInspectHandler = exec.ExecInspectHandlerFunc(func(params exec.ExecInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation exec.ExecInspect has not yet been implemented")
		})
	}
	if api.ExecExecResizeHandler == nil {
		api.ExecExecResizeHandler = exec.ExecResizeHandlerFunc(func(params exec.ExecResizeParams) middleware.Responder {
			return middleware.NotImplemented("operation exec.ExecResize has not yet been implemented")
		})
	}
	if api.ExecExecStartHandler == nil {
		api.ExecExecStartHandler = exec.ExecStartHandlerFunc(func(params exec.ExecStartParams) middleware.Responder {
			return middleware.NotImplemented("operation exec.ExecStart has not yet been implemented")
		})
	}
	if api.PluginGetPluginPrivilegesHandler == nil {
		api.PluginGetPluginPrivilegesHandler = plugin.GetPluginPrivilegesHandlerFunc(func(params plugin.GetPluginPrivilegesParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.GetPluginPrivileges has not yet been implemented")
		})
	}
	if api.ImageImageBuildHandler == nil {
		api.ImageImageBuildHandler = image.ImageBuildHandlerFunc(func(params image.ImageBuildParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageBuild has not yet been implemented")
		})
	}
	if api.ImageImageCommitHandler == nil {
		api.ImageImageCommitHandler = image.ImageCommitHandlerFunc(func(params image.ImageCommitParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageCommit has not yet been implemented")
		})
	}
	if api.ImageImageCreateHandler == nil {
		api.ImageImageCreateHandler = image.ImageCreateHandlerFunc(func(params image.ImageCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageCreate has not yet been implemented")
		})
	}
	if api.ImageImageDeleteHandler == nil {
		api.ImageImageDeleteHandler = image.ImageDeleteHandlerFunc(func(params image.ImageDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageDelete has not yet been implemented")
		})
	}
	if api.ImageImageGetHandler == nil {
		api.ImageImageGetHandler = image.ImageGetHandlerFunc(func(params image.ImageGetParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageGet has not yet been implemented")
		})
	}
	if api.ImageImageGetAllHandler == nil {
		api.ImageImageGetAllHandler = image.ImageGetAllHandlerFunc(func(params image.ImageGetAllParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageGetAll has not yet been implemented")
		})
	}
	if api.ImageImageHistoryHandler == nil {
		api.ImageImageHistoryHandler = image.ImageHistoryHandlerFunc(func(params image.ImageHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageHistory has not yet been implemented")
		})
	}
	if api.ImageImageInspectHandler == nil {
		api.ImageImageInspectHandler = image.ImageInspectHandlerFunc(func(params image.ImageInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageInspect has not yet been implemented")
		})
	}
	if api.ImageImageListHandler == nil {
		api.ImageImageListHandler = image.ImageListHandlerFunc(func(params image.ImageListParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageList has not yet been implemented")
		})
	}
	if api.ImageImageLoadHandler == nil {
		api.ImageImageLoadHandler = image.ImageLoadHandlerFunc(func(params image.ImageLoadParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageLoad has not yet been implemented")
		})
	}
	if api.ImageImagePruneHandler == nil {
		api.ImageImagePruneHandler = image.ImagePruneHandlerFunc(func(params image.ImagePruneParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImagePrune has not yet been implemented")
		})
	}
	if api.ImageImagePushHandler == nil {
		api.ImageImagePushHandler = image.ImagePushHandlerFunc(func(params image.ImagePushParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImagePush has not yet been implemented")
		})
	}
	if api.ImageImageSearchHandler == nil {
		api.ImageImageSearchHandler = image.ImageSearchHandlerFunc(func(params image.ImageSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageSearch has not yet been implemented")
		})
	}
	if api.ImageImageTagHandler == nil {
		api.ImageImageTagHandler = image.ImageTagHandlerFunc(func(params image.ImageTagParams) middleware.Responder {
			return middleware.NotImplemented("operation image.ImageTag has not yet been implemented")
		})
	}
	if api.NetworkNetworkConnectHandler == nil {
		api.NetworkNetworkConnectHandler = network.NetworkConnectHandlerFunc(func(params network.NetworkConnectParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkConnect has not yet been implemented")
		})
	}
	if api.NetworkNetworkCreateHandler == nil {
		api.NetworkNetworkCreateHandler = network.NetworkCreateHandlerFunc(func(params network.NetworkCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkCreate has not yet been implemented")
		})
	}
	if api.NetworkNetworkDeleteHandler == nil {
		api.NetworkNetworkDeleteHandler = network.NetworkDeleteHandlerFunc(func(params network.NetworkDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkDelete has not yet been implemented")
		})
	}
	if api.NetworkNetworkDisconnectHandler == nil {
		api.NetworkNetworkDisconnectHandler = network.NetworkDisconnectHandlerFunc(func(params network.NetworkDisconnectParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkDisconnect has not yet been implemented")
		})
	}
	if api.NetworkNetworkInspectHandler == nil {
		api.NetworkNetworkInspectHandler = network.NetworkInspectHandlerFunc(func(params network.NetworkInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkInspect has not yet been implemented")
		})
	}
	if api.NetworkNetworkListHandler == nil {
		api.NetworkNetworkListHandler = network.NetworkListHandlerFunc(func(params network.NetworkListParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkList has not yet been implemented")
		})
	}
	if api.NetworkNetworkPruneHandler == nil {
		api.NetworkNetworkPruneHandler = network.NetworkPruneHandlerFunc(func(params network.NetworkPruneParams) middleware.Responder {
			return middleware.NotImplemented("operation network.NetworkPrune has not yet been implemented")
		})
	}
	if api.NodeNodeDeleteHandler == nil {
		api.NodeNodeDeleteHandler = node.NodeDeleteHandlerFunc(func(params node.NodeDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation node.NodeDelete has not yet been implemented")
		})
	}
	if api.NodeNodeInspectHandler == nil {
		api.NodeNodeInspectHandler = node.NodeInspectHandlerFunc(func(params node.NodeInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation node.NodeInspect has not yet been implemented")
		})
	}
	if api.NodeNodeListHandler == nil {
		api.NodeNodeListHandler = node.NodeListHandlerFunc(func(params node.NodeListParams) middleware.Responder {
			return middleware.NotImplemented("operation node.NodeList has not yet been implemented")
		})
	}
	if api.NodeNodeUpdateHandler == nil {
		api.NodeNodeUpdateHandler = node.NodeUpdateHandlerFunc(func(params node.NodeUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation node.NodeUpdate has not yet been implemented")
		})
	}
	if api.PluginPluginCreateHandler == nil {
		api.PluginPluginCreateHandler = plugin.PluginCreateHandlerFunc(func(params plugin.PluginCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginCreate has not yet been implemented")
		})
	}
	if api.PluginPluginDeleteHandler == nil {
		api.PluginPluginDeleteHandler = plugin.PluginDeleteHandlerFunc(func(params plugin.PluginDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginDelete has not yet been implemented")
		})
	}
	if api.PluginPluginDisableHandler == nil {
		api.PluginPluginDisableHandler = plugin.PluginDisableHandlerFunc(func(params plugin.PluginDisableParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginDisable has not yet been implemented")
		})
	}
	if api.PluginPluginEnableHandler == nil {
		api.PluginPluginEnableHandler = plugin.PluginEnableHandlerFunc(func(params plugin.PluginEnableParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginEnable has not yet been implemented")
		})
	}
	if api.PluginPluginInspectHandler == nil {
		api.PluginPluginInspectHandler = plugin.PluginInspectHandlerFunc(func(params plugin.PluginInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginInspect has not yet been implemented")
		})
	}
	if api.PluginPluginListHandler == nil {
		api.PluginPluginListHandler = plugin.PluginListHandlerFunc(func(params plugin.PluginListParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginList has not yet been implemented")
		})
	}
	if api.PluginPluginPullHandler == nil {
		api.PluginPluginPullHandler = plugin.PluginPullHandlerFunc(func(params plugin.PluginPullParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginPull has not yet been implemented")
		})
	}
	if api.PluginPluginPushHandler == nil {
		api.PluginPluginPushHandler = plugin.PluginPushHandlerFunc(func(params plugin.PluginPushParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginPush has not yet been implemented")
		})
	}
	if api.PluginPluginSetHandler == nil {
		api.PluginPluginSetHandler = plugin.PluginSetHandlerFunc(func(params plugin.PluginSetParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginSet has not yet been implemented")
		})
	}
	if api.PluginPluginUpgradeHandler == nil {
		api.PluginPluginUpgradeHandler = plugin.PluginUpgradeHandlerFunc(func(params plugin.PluginUpgradeParams) middleware.Responder {
			return middleware.NotImplemented("operation plugin.PluginUpgrade has not yet been implemented")
		})
	}
	if api.ContainerPutContainerArchiveHandler == nil {
		api.ContainerPutContainerArchiveHandler = container.PutContainerArchiveHandlerFunc(func(params container.PutContainerArchiveParams) middleware.Responder {
			return middleware.NotImplemented("operation container.PutContainerArchive has not yet been implemented")
		})
	}
	if api.SecretSecretCreateHandler == nil {
		api.SecretSecretCreateHandler = secret.SecretCreateHandlerFunc(func(params secret.SecretCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation secret.SecretCreate has not yet been implemented")
		})
	}
	if api.SecretSecretDeleteHandler == nil {
		api.SecretSecretDeleteHandler = secret.SecretDeleteHandlerFunc(func(params secret.SecretDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation secret.SecretDelete has not yet been implemented")
		})
	}
	if api.SecretSecretInspectHandler == nil {
		api.SecretSecretInspectHandler = secret.SecretInspectHandlerFunc(func(params secret.SecretInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation secret.SecretInspect has not yet been implemented")
		})
	}
	if api.SecretSecretListHandler == nil {
		api.SecretSecretListHandler = secret.SecretListHandlerFunc(func(params secret.SecretListParams) middleware.Responder {
			return middleware.NotImplemented("operation secret.SecretList has not yet been implemented")
		})
	}
	if api.SecretSecretUpdateHandler == nil {
		api.SecretSecretUpdateHandler = secret.SecretUpdateHandlerFunc(func(params secret.SecretUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation secret.SecretUpdate has not yet been implemented")
		})
	}
	if api.ServiceServiceCreateHandler == nil {
		api.ServiceServiceCreateHandler = service.ServiceCreateHandlerFunc(func(params service.ServiceCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceCreate has not yet been implemented")
		})
	}
	if api.ServiceServiceDeleteHandler == nil {
		api.ServiceServiceDeleteHandler = service.ServiceDeleteHandlerFunc(func(params service.ServiceDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceDelete has not yet been implemented")
		})
	}
	if api.ServiceServiceInspectHandler == nil {
		api.ServiceServiceInspectHandler = service.ServiceInspectHandlerFunc(func(params service.ServiceInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceInspect has not yet been implemented")
		})
	}
	if api.ServiceServiceListHandler == nil {
		api.ServiceServiceListHandler = service.ServiceListHandlerFunc(func(params service.ServiceListParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceList has not yet been implemented")
		})
	}
	if api.ServiceServiceLogsHandler == nil {
		api.ServiceServiceLogsHandler = service.ServiceLogsHandlerFunc(func(params service.ServiceLogsParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceLogs has not yet been implemented")
		})
	}
	if api.ServiceServiceUpdateHandler == nil {
		api.ServiceServiceUpdateHandler = service.ServiceUpdateHandlerFunc(func(params service.ServiceUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceUpdate has not yet been implemented")
		})
	}
	if api.SessionSessionHandler == nil {
		api.SessionSessionHandler = session.SessionHandlerFunc(func(params session.SessionParams) middleware.Responder {
			return middleware.NotImplemented("operation session.Session has not yet been implemented")
		})
	}
	if api.SwarmSwarmInitHandler == nil {
		api.SwarmSwarmInitHandler = swarm.SwarmInitHandlerFunc(func(params swarm.SwarmInitParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmInit has not yet been implemented")
		})
	}
	if api.SwarmSwarmInspectHandler == nil {
		api.SwarmSwarmInspectHandler = swarm.SwarmInspectHandlerFunc(func(params swarm.SwarmInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmInspect has not yet been implemented")
		})
	}
	if api.SwarmSwarmJoinHandler == nil {
		api.SwarmSwarmJoinHandler = swarm.SwarmJoinHandlerFunc(func(params swarm.SwarmJoinParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmJoin has not yet been implemented")
		})
	}
	if api.SwarmSwarmLeaveHandler == nil {
		api.SwarmSwarmLeaveHandler = swarm.SwarmLeaveHandlerFunc(func(params swarm.SwarmLeaveParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmLeave has not yet been implemented")
		})
	}
	if api.SwarmSwarmUnlockHandler == nil {
		api.SwarmSwarmUnlockHandler = swarm.SwarmUnlockHandlerFunc(func(params swarm.SwarmUnlockParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmUnlock has not yet been implemented")
		})
	}
	if api.SwarmSwarmUnlockkeyHandler == nil {
		api.SwarmSwarmUnlockkeyHandler = swarm.SwarmUnlockkeyHandlerFunc(func(params swarm.SwarmUnlockkeyParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmUnlockkey has not yet been implemented")
		})
	}
	if api.SwarmSwarmUpdateHandler == nil {
		api.SwarmSwarmUpdateHandler = swarm.SwarmUpdateHandlerFunc(func(params swarm.SwarmUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation swarm.SwarmUpdate has not yet been implemented")
		})
	}
	if api.SystemSystemAuthHandler == nil {
		api.SystemSystemAuthHandler = system.SystemAuthHandlerFunc(func(params system.SystemAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemAuth has not yet been implemented")
		})
	}
	if api.SystemSystemDataUsageHandler == nil {
		api.SystemSystemDataUsageHandler = system.SystemDataUsageHandlerFunc(func(params system.SystemDataUsageParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemDataUsage has not yet been implemented")
		})
	}
	if api.SystemSystemEventsHandler == nil {
		api.SystemSystemEventsHandler = system.SystemEventsHandlerFunc(func(params system.SystemEventsParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemEvents has not yet been implemented")
		})
	}
	if api.SystemSystemInfoHandler == nil {
		api.SystemSystemInfoHandler = system.SystemInfoHandlerFunc(func(params system.SystemInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemInfo has not yet been implemented")
		})
	}
	if api.SystemSystemPingHandler == nil {
		api.SystemSystemPingHandler = system.SystemPingHandlerFunc(func(params system.SystemPingParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemPing has not yet been implemented")
		})
	}
	if api.SystemSystemPingHeadHandler == nil {
		api.SystemSystemPingHeadHandler = system.SystemPingHeadHandlerFunc(func(params system.SystemPingHeadParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemPingHead has not yet been implemented")
		})
	}
	if api.SystemSystemVersionHandler == nil {
		api.SystemSystemVersionHandler = system.SystemVersionHandlerFunc(func(params system.SystemVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation system.SystemVersion has not yet been implemented")
		})
	}
	if api.TaskTaskInspectHandler == nil {
		api.TaskTaskInspectHandler = task.TaskInspectHandlerFunc(func(params task.TaskInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation task.TaskInspect has not yet been implemented")
		})
	}
	if api.TaskTaskListHandler == nil {
		api.TaskTaskListHandler = task.TaskListHandlerFunc(func(params task.TaskListParams) middleware.Responder {
			return middleware.NotImplemented("operation task.TaskList has not yet been implemented")
		})
	}
	if api.TaskTaskLogsHandler == nil {
		api.TaskTaskLogsHandler = task.TaskLogsHandlerFunc(func(params task.TaskLogsParams) middleware.Responder {
			return middleware.NotImplemented("operation task.TaskLogs has not yet been implemented")
		})
	}
	if api.VolumeVolumeCreateHandler == nil {
		api.VolumeVolumeCreateHandler = volume.VolumeCreateHandlerFunc(func(params volume.VolumeCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation volume.VolumeCreate has not yet been implemented")
		})
	}
	if api.VolumeVolumeDeleteHandler == nil {
		api.VolumeVolumeDeleteHandler = volume.VolumeDeleteHandlerFunc(func(params volume.VolumeDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation volume.VolumeDelete has not yet been implemented")
		})
	}
	if api.VolumeVolumeInspectHandler == nil {
		api.VolumeVolumeInspectHandler = volume.VolumeInspectHandlerFunc(func(params volume.VolumeInspectParams) middleware.Responder {
			return middleware.NotImplemented("operation volume.VolumeInspect has not yet been implemented")
		})
	}
	if api.VolumeVolumeListHandler == nil {
		api.VolumeVolumeListHandler = volume.VolumeListHandlerFunc(func(params volume.VolumeListParams) middleware.Responder {
			return middleware.NotImplemented("operation volume.VolumeList has not yet been implemented")
		})
	}
	if api.VolumeVolumePruneHandler == nil {
		api.VolumeVolumePruneHandler = volume.VolumePruneHandlerFunc(func(params volume.VolumePruneParams) middleware.Responder {
			return middleware.NotImplemented("operation volume.VolumePrune has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
