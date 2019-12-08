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

* `datacenter` - (Required, Forces new resource, string) The data center in which the dedicated host resides.
* `hostname` - (Required, string) The host name of dedicatated host.
* `domian` - (Required, Forces new resource, string) The domain of dedicatated host..
* `router_hostname` - (Required, Forces new resource, string) The hostname of the primary router associated with the dedicated host.
* `hourly_billing` - (Optional, Forces new resource, boolean) The billing type for the host. When set to `true`, the dedicated host is billed on hourly usage. Otherwise, the dedicated host is billed on a monthly basis. The default value is `true`.
* `flavor` - (Optional, Forces new resource, string) The flavor of dedicated host. Default value `56_CORES_X_242_RAM_X_1_4_TB`. [Log in to the IBM-Cloud Infrastructure (SoftLayer) API to see available flavor types](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/813/getItems.json). Use your API as the password to log in. Log in and find the key called `keyName`.
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