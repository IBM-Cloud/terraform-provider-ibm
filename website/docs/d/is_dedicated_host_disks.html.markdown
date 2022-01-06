---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_disks"
description: |-
  Get information about dedicated host disk collection.
---

# ibm_is_dedicated_host_disks

Retrieve the dedicated host disk collection. For more information, about dedicated host disk collection, see [managing dedicated hosts and groups](https://cloud.ibm.com/docs/vpc?topic=vpc-manage-dedicated-hosts-groups).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_dedicated_host_disks" "example" {
  dedicated_host = ibm_is_dedicated_host.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `dedicated_host` - (Required, String) The dedicated host identifier.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `disks` - (List) The collection of the dedicated host's disks. 

  Nested scheme for `disks`:
  - `available` - The remaining space left for instance placement in GB (gigabytes).
  - `created_at` - The date and time that the disk was created.
  - `href` - The URL for this disk.
  - `id` - The unique identifier for this disk.
  - `instance_disks` - (List) Instance disks that are on this dedicated host disk. 

    Nested scheme for `instance_disks`:
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. 

      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this instance disk.
    - `id` - (String) The unique identifier for this instance disk.
    - `name` - (String) The user-defined name for this disk.
    - `resource_type` - (String) The resource type.
  - `interface_type` - (String) The disk interface used for attaching the diskThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
  - `lifecycle_state` - (String) The lifecycle state of this dedicated host disk.
  - `name` - (String) The user-defined or system-provided name for this disk.
  - `provisionable` - (String) Indicates whether this dedicated host disk is available for instance disk creation.
  - `resource_type` - (String) The type of resource referenced.
  - `size` - (String) The size of the disk in GB (gigabytes).
  - `supported_instance_interface_types` - (String) The instance disk interfaces supported for this dedicated host disk.
- `id` - (String) The unique identifier of the dedicated host disk collection.
