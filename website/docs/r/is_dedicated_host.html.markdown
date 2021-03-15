---
layout: "ibm"
page_title: "IBM : is_dedicated_host"
sidebar_current: "docs-ibm-resource-is-dedicated-host"
description: |-
  Manages IBM DedicatedHost.
---

# ibm\_is_dedicated_host

Provides a resource for DedicatedHost. This allows DedicatedHost to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_is_dedicated_host_group" "dh_group01" {
  family = "memory"
  class = "beta"
  zone = "us-south-1"
}
data "ibm_is_dedicated_host_group" "dgroup" {
	name = ibm_is_dedicated_host_group.dh_group01.name
}
resource "ibm_is_dedicated_host" "is_dedicated_host" {
  profile = "dh2-56x464"
	host_group = "cdc21b72d4e647b195de988b175e3d82"
	name = "testdh02"
}
```

## Argument Reference

The following arguments are supported:


* `instance_placement_enabled` - (Optional, bool) If set to true, instances can be placed on this dedicated host.
* `name` - (Optional, string) The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.
* `profile` - (Required, string) The Globally unique name of the dedicated host profile to use for this dedicated host.
* `resource_group` - (Optional, string) The unique identifier of the resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
* `host_group` - (Optional, string) The unique identifier of the dedicated host group for this dedicated host.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHost.
* `available_memory` - The amount of memory in gibibytes that is currently available for instances.
* `available_vcpu` - The available VCPU for the dedicated host.
* `created_at` - The date and time that the dedicated host was created.
* `crn` - The CRN for this dedicated host.
* `host_group` - The unique identifier of the dedicated host group this dedicated host is in.
* `href` - The URL for this dedicated host.
* `instance_placement_enabled` - If set to true, instances can be placed on this dedicated host.
* `instances` - Array of instances that are allocated to this dedicated host.
* `lifecycle_state` - The lifecycle state of the dedicated host resource.
* `memory` - The total amount of memory in gibibytes for this host.
* `name` - The unique user-defined name for this dedicated host.
* `profile` - The profile this dedicated host uses.
* `provisionable` - Indicates whether this dedicated host is available for instance creation.
* `resource_group` - The unique identifier of the resource group for this dedicated host.
* `resource_type` - The type of resource referenced.
* `socket_count` - The total number of sockets for this host.
* `state` - The administrative state of the dedicated host.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.
* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on this dedicated host.
* `vcpu` - The total VCPU of the dedicated host.
* `zone` - The zone this dedicated host resides in.
