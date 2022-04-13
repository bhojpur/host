package service

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
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"reflect"
	"strconv"
	"strings"
	"syscall"
	"time"

	v3 "github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/container/cluster"
	"github.com/bhojpur/host/pkg/container/drivers/aks"
	"github.com/bhojpur/host/pkg/container/drivers/eks"
	"github.com/bhojpur/host/pkg/container/drivers/gke"
	kubeimport "github.com/bhojpur/host/pkg/container/drivers/import"
	"github.com/bhojpur/host/pkg/container/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	pluginAddress = map[string]string{}
	Drivers       = map[string]types.Driver{
		GoogleKubernetesEngineDriverName:        gke.NewDriver(),
		AzureKubernetesServiceDriverName:        aks.NewDriver(),
		AmazonElasticContainerServiceDriverName: eks.NewDriver(),
		ImportDriverName:                        kubeimport.NewDriver(),
	}
)

const (
	ListenAddress                           = "127.0.0.1:"
	GoogleKubernetesEngineDriverName        = "googlekubernetesengine"
	AzureKubernetesServiceDriverName        = "azurekubernetesservice"
	AmazonElasticContainerServiceDriverName = "amazonelasticcontainerservice"
	ImportDriverName                        = "import"
	BhojpurKubernetesEngineDriverName       = "bhojpurkubernetesengine"
)

type controllerConfigGetter struct {
	driverName  string
	clusterSpec v3.ClusterSpec
	clusterName string
}

func (c controllerConfigGetter) GetConfig() (types.DriverOptions, error) {
	driverOptions := types.DriverOptions{
		BoolOptions:        make(map[string]bool),
		StringOptions:      make(map[string]string),
		IntOptions:         make(map[string]int64),
		StringSliceOptions: make(map[string]*types.StringSlice),
	}
	data := map[string]interface{}{}
	switch c.driverName {
	case ImportDriverName:
		config, err := toMap(c.clusterSpec.ImportedConfig, "json")
		if err != nil {
			return driverOptions, err
		}
		data = config
		flatten(data, &driverOptions)
	default:
		config, err := toMap(c.clusterSpec.GenericEngineConfig, "json")
		if err != nil {
			return driverOptions, err
		}
		data = config
		flatten(data, &driverOptions)
	}

	driverOptions.StringOptions["name"] = c.clusterName
	displayName := c.clusterSpec.DisplayName
	if displayName == "" {
		displayName = c.clusterName
	}
	driverOptions.StringOptions["displayName"] = displayName

	return driverOptions, nil
}

// flatten take a map and flatten it and convert it into driverOptions
func flatten(data map[string]interface{}, driverOptions *types.DriverOptions) {
	for k, v := range data {
		switch v.(type) {
		case float64:
			driverOptions.IntOptions[k] = int64(v.(float64))
		case string:
			driverOptions.StringOptions[k] = v.(string)
		case bool:
			driverOptions.BoolOptions[k] = v.(bool)
		case []interface{}:
			// lists of strings come across as lists of interfaces, have to convert them manually
			var stringArray []string

			for _, stringInterface := range v.([]interface{}) {
				switch stringInterface.(type) {
				case string:
					stringArray = append(stringArray, stringInterface.(string))
				}
			}

			// if the length is 0 then it must not have been an array of strings
			if len(stringArray) != 0 {
				driverOptions.StringSliceOptions[k] = &types.StringSlice{Value: stringArray}
			}
		case []string:
			driverOptions.StringSliceOptions[k] = &types.StringSlice{Value: v.([]string)}
		case map[string]interface{}:
			// hack for labels
			if k == "labels" {
				r := []string{}
				for key1, value1 := range v.(map[string]interface{}) {
					r = append(r, fmt.Sprintf("%v=%v", key1, value1))
				}
				driverOptions.StringSliceOptions[k] = &types.StringSlice{Value: r}
			} else {
				flatten(v.(map[string]interface{}), driverOptions)
			}
		case nil:
			logrus.Debugf("could not convert %v because value is nil %v=%v", reflect.TypeOf(v), k, v)
		default:
			logrus.Warnf("could not convert %v %v=%v", reflect.TypeOf(v), k, v)
		}
	}
}

func toMap(obj interface{}, format string) (map[string]interface{}, error) {
	if format == "json" {
		data, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		return result, nil
	} else if format == "yaml" {
		data, err := yaml.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var result map[string]interface{}
		if err := yaml.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}

type EngineService struct {
	store cluster.PersistentStore
}

func NewEngineService(store cluster.PersistentStore) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (e *EngineService) convertCluster(name string, listenAddr string, spec v3.ClusterSpec) (*cluster.Cluster, error) {
	// todo: decide whether we need a driver field
	driverName := ""
	if spec.ImportedConfig != nil {
		driverName = ImportDriverName
	} else if spec.BhojpurKubernetesEngineConfig != nil {
		driverName = BhojpurKubernetesEngineDriverName
	} else if spec.GenericEngineConfig != nil {
		driverName = (*spec.GenericEngineConfig)["driverName"].(string)
		if driverName == "" {
			return nil, fmt.Errorf("no driver name supplied")
		}
	}
	if driverName == "" {
		return nil, fmt.Errorf("no driver config found")
	}

	configGetter := controllerConfigGetter{
		driverName:  driverName,
		clusterSpec: spec,
		clusterName: name,
	}
	clusterPlugin, err := cluster.NewCluster(driverName, name, listenAddr, configGetter, e.store)
	if err != nil {
		return nil, err
	}

	// verify driver is running
	failures := 0
	for {
		_, err = clusterPlugin.GetCapabilities(context.Background())
		if err == nil {
			break
		} else if failures > 5 {
			clusterPlugin.Driver.Close()
			return nil, fmt.Errorf("error checking driver is up: %v", err)
		}

		failures = failures + 1
		time.Sleep(time.Duration(failures*failures) * time.Second)
	}

	return clusterPlugin, nil
}

// Create creates the stub for cluster manager to call
func (e *EngineService) Create(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) (string, string, string, error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return "", "", "", err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return "", "", "", fmt.Errorf("error starting driver: %v", err)
	}

	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return "", "", "", err
	}

	defer cls.Driver.Close()

	if err := cls.Create(ctx); err != nil {
		return "", "", "", err
	}
	endpoint := cls.Endpoint
	if !strings.HasPrefix(endpoint, "https://") {
		endpoint = fmt.Sprintf("https://%s", cls.Endpoint)
	}
	return endpoint, cls.ServiceAccountToken, cls.RootCACert, nil
}

func (e *EngineService) getRunningDriver(hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) (*RunningDriver, error) {
	return &RunningDriver{
		Name:    hostfarmDriver.Name,
		Builtin: hostfarmDriver.Spec.BuiltIn,
		Path:    hostfarmDriver.Status.ExecutablePath,
	}, nil
}

// Update creates the stub for cluster manager to call
func (e *EngineService) Update(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) (string, string, string, error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return "", "", "", err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return "", "", "", fmt.Errorf("error starting driver: %v", err)
	}

	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return "", "", "", err
	}

	defer cls.Driver.Close()

	if err := cls.Update(ctx); err != nil {
		return "", "", "", err
	}
	endpoint := cls.Endpoint
	if !strings.HasPrefix(endpoint, "https://") {
		endpoint = fmt.Sprintf("https://%s", cls.Endpoint)
	}
	return endpoint, cls.ServiceAccountToken, cls.RootCACert, nil
}

// Remove removes stub for cluster manager to call
func (e *EngineService) Remove(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec, forceRemove bool) error {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return fmt.Errorf("error starting driver: %v", err)
	}

	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return err
	}

	defer cls.Driver.Close()

	return cls.Remove(ctx, forceRemove)
}

func (e *EngineService) GetDriverCreateOptions(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) (*types.DriverFlags,
	error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return nil, err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting driver: %v", err)
	}

	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return nil, err
	}

	defer cls.Driver.Close()

	return cls.GetDriverCreateOptions(ctx)
}

func (e *EngineService) GetDriverUpdateOptions(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) (*types.DriverFlags,
	error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return nil, err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting driver: %v", err)
	}

	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return nil, err
	}

	defer cls.Driver.Close()

	return cls.GetDriverUpdateOptions(ctx)
}

func (e *EngineService) GetK8sCapabilities(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver,
	clusterSpec v3.ClusterSpec) (*types.K8SCapabilities, error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return nil, err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting driver: %v", err)
	}

	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return nil, err
	}

	defer cls.Driver.Close()

	return cls.GetK8SCapabilities(ctx)
}

type RunningDriver struct {
	Name    string
	Path    string
	Builtin bool
	Server  *types.GrpcServer

	listenAddress string
	cancel        context.CancelFunc
}

func (r *RunningDriver) Start() (string, error) {
	ephemeralListenAddress := fmt.Sprintf("%s0", ListenAddress)
	p, err := net.Listen("tcp", ephemeralListenAddress) // passing this port will cause go to provide open ephemeral port
	if err != nil {
		return "", fmt.Errorf("failed retrieving port for driver: %v", err)
	}

	listenAddress := p.Addr().String()
	if err := p.Close(); err != nil {
		return "", fmt.Errorf("failed to close port before starting driver: %v", err)
	}

	port, err := portOnly(listenAddress)
	if err != nil {
		return "", err
	}

	if r.Builtin {
		driver := Drivers[r.Name]
		if driver == nil {
			return "", fmt.Errorf("no driver for name: %v", r.Name)
		}

		addr := make(chan string)
		errChan := make(chan error)
		r.Server = types.NewServer(driver, addr)
		go r.Server.Serve(listenAddress, errChan)

		// if the error hasn't appeared after 5 seconds assume it won't error
		var err error
		select {
		case err = <-errChan:
			// get error
		case <-time.After(5 * time.Second):
			// do nothing
		}
		if err != nil {
			return "", fmt.Errorf("error starting driver: %v", err)
		}

		r.listenAddress = <-addr
	} else {
		var processContext context.Context
		processContext, r.cancel = context.WithCancel(context.Background())

		cmd := exec.CommandContext(processContext, r.Path, port)

		if os.Getenv("BHOJPUR_DEV_MODE") == "" {
			cred, err := getUserCred()
			if err != nil {
				return "", errors.WithMessage(err, "get user cred error")
			}

			cmd.SysProcAttr = &syscall.SysProcAttr{}
			cmd.SysProcAttr.Credential = cred
			cmd.SysProcAttr.Chroot = "/opt/jail/driver-jail"
			cmd.Env = whitelistEnvvars([]string{"PATH=/usr/bin"})
		}

		// redirect output to console
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			return "", fmt.Errorf("error starting driver: %v", err)
		}

		time.Sleep(5 * time.Second)

		r.listenAddress = listenAddress
	}

	logrus.Infof("Bhojpur Host farm driver %v listening on address %v", r.Name, r.listenAddress)

	return r.listenAddress, nil
}

// portOnly attempts to return port fragment of address
func portOnly(address string) (string, error) {
	portParseErr := fmt.Errorf("failed to parse port from address [%s]", address)

	_, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", errors.Wrap(err, portParseErr.Error())
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", portParseErr
	}

	if portNum < 1 || portNum > 65535 {
		return "", errors.Wrap(fmt.Errorf(fmt.Sprintf("invalid port [%s], port range is between 1 and 65535", port)), portParseErr.Error())
	}

	return port, nil
}

func (r *RunningDriver) Stop() {
	if r.Builtin {
		r.Server.Stop()
	} else {
		r.cancel()
	}

	logrus.Infof("Bhojpur Host farm driver %v stopped", r.Name)
}

func (e *EngineService) ETCDSave(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec, snapshotName string) error {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return fmt.Errorf("error starting driver: %v", err)
	}
	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return err
	}
	defer cls.Driver.Close()

	return cls.ETCDSave(ctx, snapshotName)
}

func (e *EngineService) ETCDRestore(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec, backup string) (string, string, string, error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return "", "", "", err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return "", "", "", fmt.Errorf("error starting driver: %v", err)
	}
	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return "", "", "", err
	}
	defer cls.Driver.Close()

	if err = cls.ETCDRestore(ctx, backup); err != nil {
		return "", "", "", err
	}

	endpoint := cls.Endpoint
	if !strings.HasPrefix(endpoint, "https://") {
		endpoint = fmt.Sprintf("https://%s", cls.Endpoint)
	}
	return endpoint, cls.ServiceAccountToken, cls.RootCACert, nil

}

func (e *EngineService) ETCDRemoveSnapshot(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec, snapshotName string) error {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return fmt.Errorf("error starting driver: %v", err)
	}
	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return err
	}
	defer cls.Driver.Close()

	return cls.ETCDRemoveSnapshot(ctx, snapshotName)
}

func (e *EngineService) GenerateServiceAccount(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) (string, error) {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return "", err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return "", fmt.Errorf("error starting driver: %v", err)
	}
	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return "", err
	}
	defer cls.Driver.Close()

	err = cls.GenerateServiceAccount(ctx)
	if err != nil {
		return "", err
	}

	return cls.ServiceAccountToken, nil
}

func (e *EngineService) RemoveLegacyServiceAccount(ctx context.Context, name string, hostfarmDriver *v3.ContainerDriver, clusterSpec v3.ClusterSpec) error {
	runningDriver, err := e.getRunningDriver(hostfarmDriver, clusterSpec)
	if err != nil {
		return err
	}

	listenAddr, err := runningDriver.Start()
	if err != nil {
		return fmt.Errorf("error starting driver: %v", err)
	}
	defer runningDriver.Stop()

	cls, err := e.convertCluster(name, listenAddr, clusterSpec)
	if err != nil {
		return err
	}
	defer cls.Driver.Close()

	return cls.RemoveLegacyServiceAccount(ctx)
}

// getUserCred looks up the user and provides it in syscall.Credential
func getUserCred() (*syscall.Credential, error) {
	u, err := user.Current()
	if err != nil {
		uID := os.Getuid()
		u, err = user.LookupId(strconv.Itoa(uID))
		if err != nil {
			return nil, err
		}
	}

	i, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return nil, err
	}
	uid := uint32(i)

	i, err = strconv.ParseUint(u.Gid, 10, 32)
	if err != nil {
		return nil, err
	}
	gid := uint32(i)

	return &syscall.Credential{Uid: uid, Gid: gid}, nil
}

func whitelistEnvvars(envvars []string) []string {
	wl := os.Getenv("BHOJPUR_HOSTFARM_ENGINE_WHITELIST_ENVVARS")
	envWhiteList := strings.Split(wl, ",")

	if len(envWhiteList) == 0 {
		envWhiteList = []string{
			"HTTP_PROXY",
			"HTTPS_PROXY",
			"NO_PROXY",
		}
	}

	for _, wlVar := range envWhiteList {
		if val := os.Getenv(wlVar); val != "" {
			envvars = append(envvars, wlVar+"="+val)
		}
	}

	return envvars
}
