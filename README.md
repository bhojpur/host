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

- Computing Host Framework
- Clustering Framework

## Simple Usage

```bash
$ hostutl
```

## Build Source Code

Firstly, you need a running instance of `Docker` so that the `dapper` tools can build everything
from the source code and use a local `Docker Registry`.

```bash
$ make
```
