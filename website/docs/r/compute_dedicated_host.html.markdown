---
layout: "ibm"
page_title: "IBM : compute_dedicated_host"
sidebar_current: "docs-ibm-resource-compute-dedicated-host"
description: |-
  Manages IBM Dedicated Host.
---

# ibm\_compute_dedicated_host

Provides a Dedicated Host resource. This allows dedicated host to be created, updated, and cancelled.

For additional details please refer to the [SoftLayer API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_DedicatedHost).

## Example Usage

In the following example, you can create a dedicated host:

```hcl
resource "ibm_compute_dedicated_host" "dedicatedhost" {
  hostname        = "host"
  domain          = "example.com"
  router_hostname = "bcr01a.dal09"
  datacenter      = "dal09"
}

```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Required, string) The data center in which the dedicated host resides.
* `hostname` - (Required, string) The host name of dedicatated host.
* `domian` - (Required, string) The domain of dedicatated host..
* `router_hostname` - (Required, string) The hostname of the primary router associated with the dedicated host.
* `hourly_billing` - (Optional, boolean) The billing type for the host. When set to `true`, the dedicated host is billed on hourly usage. Otherwise, the dedicated host is billed on a monthly basis. The default value is `true`.
* `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the dedicated host to become available before declaring it as created. The default value is `90`.
* `tags` - (Optional, array of strings) Tags associated with the dedicated host.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the dedicated host.
* `cpu_count` - The capacity that the dedicated host's CPU allocation is restricted to.
* `disk_capacity` - The capacity that the dedicated host's disk allocation is restricted to.
* `memory_capacity` - The capacity that the dedicated host's memory allocation is restricted to.

## Import

`ibm_compute_dedicated_host` can be imported using the ID.

Example:

```
$ terraform import ibm_compute_dedicated_host.dedicatedhost 238756

```