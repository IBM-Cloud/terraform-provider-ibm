---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_group"
description: |-
  Get information about DedicatedHostGroup
---

# ibm\_is_dedicated_host_group

Provides a read-only data source for DedicatedHostGroup. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_dedicated_host_group" "is_dedicated_host_group" {
	name = "my-host-group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The unique user-defined name for this dedicated host.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the DedicatedHostGroup.
* `class` - The dedicated host profile class for hosts in this group.
* `created_at` - The date and time that the dedicated host group was created.
* `crn` - The CRN for this dedicated host group.
* `dedicated_hosts` - The dedicated hosts that are in this dedicated host group. Nested `dedicated_hosts` blocks have the following structure:
	* `crn` - The CRN for this dedicated host.
	* `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		* `more_info` - Link to documentation about deleted resources.
	* `href` - The URL for this dedicated host.
	* `id` - The unique identifier for this dedicated host.
	* `name` - The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.
	* `resource_type` - The type of resource referenced.
* `family` - The dedicated host profile family for hosts in this group.
* `href` - The URL for this dedicated host group.
* `resource_group` - The unique identifier of the resource group for this dedicated host.
* `resource_type` - The type of resource referenced.
* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on this dedicated host group. Nested `supported_instance_profiles` blocks have the following structure:
	* `href` - The URL for this virtual server instance profile.
	* `name` - The globally unique name for this virtual server instance profile.
* `zone` - The zone this dedicated host group resides in.

