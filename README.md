# Bhojpur Host - Infrastructure Framework

The `Bhojpur Host` is a high-performance, compute hosting resource provider applied within the
[Bhojpur.NET Platform](https://github.com/bhojpur/platform) ecosystem for delivery of distributed
`applications` or `services`. It lets you create `virtual` or `Container` hosts and/or `Unikernel`
hosts on a bare metal computer, on cloud providers (e.g., Amazon Web Services, Google Cloud Platform,
Microsoft Azure, DigitalOcean, IBM Bluemix, Oracle Cloud Infrastructre), and inside your own data
center (e.g. using OpenStack, VMware). It creates servers, installs the `Bhojpur Host` server-side
runtime engine on them, then configures the `Bhojpur Host` client-side CLI to talk to them.

The `Bhojpur Host` project is intended to be *embedded* and executed by full Bhojpur products. The
standalone Bhojpur Host `CLI` functionality will remain, but human use of it will not be our primary
focus as we would expect inputs provided by other things (e.g., Terraform or User Inferfaces)

## Key Features

- Computing Host instance management
- Host Farm (i.e. cluster) management
- Container Cluster Operations management

It integrates seamlessly with the [Bhojpur DCP](https://github.com/bhojpur/dcp) (i.e., distributed
cloud platform), which leverages [Bhojpur OS](https://github.com/bhojpur/os) as its operating system
and the [Bhojpur Kernel](https://github.com/bhojpur/kernel) framework for `Unikernel` development.

## Computing Host Instance Provision

The `Bhojpur Host` machine instance is managed using `hostutl` tool. It can provision hosting nodes
(i.e. machine instances) in a Virtualized Data Center. It could connect to a wide range of Cloud
infrastructure providers.

To get a list of hosting `machines` available in the resource pool, type the following command

```bash
$ hostutl ls
```

To create a new `machine` instance locally using `virtualbox` driver

```bash
$ hostutl create --driver virtualbox default
```

To remove a host `machine` from managed resource pool

```bash
$ hostutl rm default
```

## Host Farm Cluster Provision

The `Bhojpur Host` server farm is managed using the `hostfarm` tool. It creates Kubernetes
clusters on different Cloud infrastructure providers.

A list of commands to manage `Kubernetes Cluster` in a multi-cloud environment

```bash
$ hostfarm create --driver $driverName [OPTIONS] cluster-name
$ hostfarm inspect cluster-name
$ hostfarm ls
$ hostfarm update [OPTIONS] cluster-name
$ hostfarm rm cluster-name
```

To see what driver `create` options it has, run the following command

```bash
$ hostfarm `create --driver` $driverName --help
```

To see what `update` options for a cluster, run the following command

```bash
$ hostfarm update --help cluster-name
```

A `serviceAccountToken` that binds to the `clusterAdmin` is automatically created for you,
to see what it is, run the following command

```bash
$ hostfarm inspect clusterName
```

Before running `Google Kubernetes Engine` driver, make sure you have the credential. To get
the credential, you can run any of the steps below

`gcloud auth login` or,

`export GOOGLE_APPLICATION_CREDENTIALS=$HOME/gce-credentials.json` or,

```bash
$ hostfarm create --driver gke --gke-credential-path /path/to/credential cluster-name
```

## Container Cluster Operations Management

The `Bhojpur Kubernetes Engine` could host a wide range of `applications` or `services`. It
is managed using the `hostops` tool. The installation and configuration of hosting clusters
is the primary focus of `hostops` tool.

```bash
$ hostops
```

## Build Source Code

Firstly, you need a running instance of `Docker` instance so that the `dapper` tool can
build everything from the source code and use a local `Docker Registry`.

```bash
$ sudo make
```

Alternately, you can use the following commands

```bash
$ task build-tools
$ task build-cloud
```
