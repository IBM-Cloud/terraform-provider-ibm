---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host"
description: |-
  Get information about DedicatedHost
---

# ibm\_is_dedicated_host

Provides a read-only data source for DedicatedHost. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_dedicated_host" "is_dedicated_host" {
	host_group = "1e09281b-f177-46fb-baf1-bc152b2e391a"
	name = "my-dedicated-host"
}
```

## Argument Reference

The following arguments are supported:

* `host_group` - (Required, string) The unique identifier of the Dedicated host group.
* `name` - (Required, string) The unique name of this dedicated host.
* `resource_group` - (Optional, string) The unique identifier of the resource group

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHost.
* `available_memory` - The amount of memory in gibibytes that is currently available for instances.

* `available_vcpu` - The available VCPU for the dedicated host. Nested `available_vcpu` blocks have the following structure:
	* `architecture` - The VCPU architecture.
	* `count` - The number of VCPUs assigned.

* `created_at` - The date and time that the dedicated host was created.

* `crn` - The CRN for this dedicated host.

* `host_group` - The unique identifier of the dedicated host group this dedicated host is in.
* `href` - The URL for this dedicated host.

* `instance_placement_enabled` - If set to true, instances can be placed on this dedicated host.

* `instances` - Array of instances that are allocated to this dedicated host. Nested `instances` blocks have the following structure:
	* `crn` - The CRN for this virtual server instance.
	* `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		* `more_info` - Link to documentation about deleted resources.
	* `href` - The URL for this virtual server instance.
	* `id` - The unique identifier for this virtual server instance.
	* `name` - The user-defined name for this virtual server instance (and default system hostname).

* `lifecycle_state` - The lifecycle state of the dedicated host resource.

* `memory` - The total amount of memory in gibibytes for this host.

* `name` - The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.

* `profile` - The profile this dedicated host uses. Nested `profile` blocks have the following structure:
	* `href` - The URL for this dedicated host.
	* `name` - The globally unique name for this dedicated host profile.

* `provisionable` - Indicates whether this dedicated host is available for instance creation.

* `resource_group` - The unique identifier of the resource group.
* `resource_type` - The type of resource referenced.

* `socket_count` - The total number of sockets for this host.

* `state` - The administrative state of the dedicated host.

* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on this dedicated host. Nested `supported_instance_profiles` blocks have the following structure:
	* `href` - The URL for this virtual server instance profile.
	* `name` - The globally unique name for this virtual server instance profile.

* `vcpu` - The total VCPU of the dedicated host. Nested `vcpu` blocks have the following structure:
	* `architecture` - The VCPU architecture.
	* `count` - The number of VCPUs assigned.

* `zone` - The globally unique name of the zone this dedicated host resides in.
