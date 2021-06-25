---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_disk"
description: |-
  Manages IBM Cloud Bare Metal Server Disk.
---

# ibm\_is_bare_metal_server_disk

Import the details of an existing IBM Cloud Bare Metal Server Disk as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```terraform

data "ibm_is_bare_metal_server_disk" "ds_bms_disk" {
  bare_metal_server         = "xxxx-xxxxx-xxxxx-xxxx"
  disk                      = "xxxx-xxxxx-xxxxx-xxxx"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `bare_metal_server` - (Required, String) The id for this bare metal server.
- `disk` - (Required, String) The id for this bare metal server disk.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `href` - (String) The URL for this bare metal server disk.
- `id` - (String) The unique identifier for this bare metal server disk.
- `interface_type` - (String) The disk interface used for attaching the disk. Supported values are [ **nvme**, **sata** ].
- `name` - (String) The user-defined name for this disk.
- `resource_type` - (String) The resource type.
- `size` - (String) The size of the disk in GB (gigabytes).

