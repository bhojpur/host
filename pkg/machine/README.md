# Bhojpur Host - Machine Provision Engine

A hosting management tool to provision Machine (i.e. nodes) for different cloud providers. It lets
you create Docker hosts on your computer, on cloud providers, and inside your own data center. It
creates servers, installs Docker on them, then configures the Docker client to talk to them.

## Installation

The package is intended to be embedded and executed by full [Bhojpur Host](https://github.com/bhojpur/host)
product and stand alone Bhojpur CLI functionality will remain, but the human use of it will not be the
primary focus as we will expect inputs provided by other things like Terraform or UIs.

Bhojpur CLI binaries can be found in our [Releases Page](https://github.com/bhojpur/host/releases)

## Driver Plugins

In addition to the core driver plugins bundled alongside Bhojpur Host machine, users can make and distribute
their own plugin for any virtualization technology or cloud provider.
