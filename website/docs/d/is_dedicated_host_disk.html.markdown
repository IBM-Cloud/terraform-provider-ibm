---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_disk"
description: |-
  Get information about dedicated host disk.
---

# ibm_is_dedicated_host_disk

Retrieve the dedicated host disk. For more information, about dedicated host disk, see [migrating a dedicated host instance to another host](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-migrating-dedicated-host).

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
data "ibm_is_dedicated_host_disk" "example" {
  dedicated_host = ibm_is_dedicated_host.example.id
  disk           = data.ibm_is_dedicated_host_disks.example.disks.0.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `dedicated_host` - (Required, String) The dedicated host identifier.
- `disk` - (Required, String) The dedicated host disk identifier.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `available` - (String) The remaining space left for instance placement in GB (gigabytes).
- `created_at` - (Timestamp) The date and time that the disk was created.
- `href` - The URL for this disk.
- `id` - (String) The unique identifier of the dedicated host disk.
- `instance_disks` - (List) Instance disks that are on the dedicated host disk. 

  Nested scheme for `instance_disks`:
  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides the supplementary information. 

      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
  - `href` - (String) The URL for this instance disk.
  - `id` - (String) The unique identifier for this instance disk.
  - `name` - (String) The user defined name for this disk.
  - `resource_type` - (String) The resource type.
- `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halts processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
- `lifecycle_state` - (String) The lifecycle state of this dedicated host disk.
- `name` - (String) The user defined or system provided name for this disk.
- `provisionable` - (String) Indicates whether this dedicated host disk is available for instance disk creation.
- `resource_type` - (String) The type of resource referenced.
- `size` - (String) The size of the disk in GB (gigabytes).
- `supported_instance_interface_types` - (String) The instance disk interfaces supported for this dedicated host disk.

