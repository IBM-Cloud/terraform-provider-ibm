---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_disks"
description: |-
  Get information about DedicatedHostDiskCollection
---

# ibm\_is_dedicated_host_disks

Provides a read-only data source for DedicatedHostDiskCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "is_dedicated_host_disks" "is_dedicated_host_disks" {
	dedicated_host = "dedicatedhost id"
}
```

## Argument Reference

The following arguments are supported:

* `dedicated_host` - (Required, string) The dedicated host identifier.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHostDiskCollection.
* `disks` - Collection of the dedicated host's disks. Nested `disks` blocks have the following structure:
	* `available` - The remaining space left for instance placement in GB (gigabytes).
	* `created_at` - The date and time that the disk was created.
	* `href` - The URL for this disk.
	* `id` - The unique identifier for this disk.
	* `instance_disks` - Instance disks that are on this dedicated host disk. Nested `instance_disks` blocks have the following structure:
		* `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			* `more_info` - Link to documentation about deleted resources.
		* `href` - The URL for this instance disk.
		* `id` - The unique identifier for this instance disk.
		* `name` - The user-defined name for this disk.
		* `resource_type` - The resource type.
	* `interface_type` - The disk interface used for attaching the diskThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	* `lifecycle_state` - The lifecycle state of this dedicated host disk.
	* `name` - The user-defined or system-provided name for this disk.
	* `provisionable` - Indicates whether this dedicated host disk is available for instance disk creation.
	* `resource_type` - The type of resource referenced.
	* `size` - The size of the disk in GB (gigabytes).
	* `supported_instance_interface_types` - The instance disk interfaces supported for this dedicated host disk.

