---
layout: "ibm"
page_title: "IBM : is_dedicated_host_groups"
sidebar_current: "docs-ibm-datasource-is-dedicated-host-groups"
description: |-
  Get information about DedicatedHostGroupCollection
---

# ibm\_is_dedicated_host_groups

Provides a read-only data source for DedicatedHostGroupCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "is_dedicated_host_groups" "is_dedicated_host_groups" {
	name = "my-host-group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHostGroupCollection.
* `first` - A link to the first page of resources. Nested `first` blocks have the following structure:
	* `href` - The URL for a page of resources.

* `groups` - Collection of dedicated host groups. Nested `groups` blocks have the following structure:
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
	* `id` - The unique identifier for this dedicated host group.
	* `name` - The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.
	* `resource_group` - The resource group for this dedicated host group. Nested `resource_group` blocks have the following structure:
		* `href` - The URL for this resource group.
		* `id` - The unique identifier for this resource group.
		* `name` - The user-defined name for this resource group.
	* `resource_type` - The type of resource referenced.
	* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on this dedicated host group. Nested `supported_instance_profiles` blocks have the following structure:
		* `href` - The URL for this virtual server instance profile.
		* `name` - The globally unique name for this virtual server instance profile.
	* `zone` - The zone this dedicated host group resides in. Nested `zone` blocks have the following structure:
		* `href` - The URL for this zone.
		* `name` - The globally unique name for this zone.

* `limit` - The maximum number of resources that can be returned by the request.

* `next` - A link to the next page of resources. This property is present for all pagesexcept the last page. Nested `next` blocks have the following structure:
	* `href` - The URL for a page of resources.

* `total_count` - The total number of resources across all pages.

