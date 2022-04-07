# Bhojpur Host - Infrastructure Framework

The `Bhojpur Host` is a high-performance, computing resource provider applied within the
[Bhojpur.NET Platform](https://github.com/bhojpur/platform) ecosystem for delivery of distributed
`applications` or `services`. It lets you create `Container` or `Unikernel` hosts on a bare metal
computer, on cloud providers (e.g., Amazon Web Services, Google Cloud Platform, Microsoft Azure,
DigitalOcean, IBM Bluemix), and inside your own data center (e.g. using OpenStack, VMware). It
creates servers, installs the `Bhojpur Host` server-side runtime engine on them, then configures
the `Bhojpur Host` client-side CLI to talk to them.

The `Bhojpur Host` project is intended to be *embedded* and executed by full Bhojpur products. The
standalone Bhojpur CLI functionality will remain, but human use of it will not be our primary focus
as we would expect inputs provided by other things (e.g., Terraform or User Inferfaces)

## Key Features

- Computing Machine instance management
- Container Cluster management framework

## Computing Machine Instance

List of hosting `machines` available in the resource pool

```bash
$ hostutl ls
```

Creating a new `machine` instance locally using `virtualbox` driver

```bash
$ hostutl create --driver virtualbox default
```

Removing a host `machine` from managed resource pool

```bash
$ hostutl rm default
```

## Container Cluster Management

A list of commands to manage `Kubernetes Cluster` in a multi-cloud environment

```bash
$ hostfarm create --driver $driverName [OPTIONS] cluster-name
$ hostfarm inspect cluster-name
$ hostfarm ls
$ hostfarm update [OPTIONS] cluster-name
$ hostfarm rm cluster-name
```

To see what driver `create` options it has, run

```bash
$ hostfarm `create --driver` $driverName --help
```

To see what `update` options for a cluster, run

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

## Build Source Code

Firstly, you need a running instance of `Docker` so that the `dapper` tools can build everything
from the source code and use a local `Docker Registry`.

```bash
$ make
```
