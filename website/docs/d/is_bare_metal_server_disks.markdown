---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_disks"
description: |-
  Manages IBM Cloud Bare Metal Server Disks.
---

# ibm\_is_bare_metal_server_disks

Import the details of an existing IBM Cloud vBare Metal Server Disk collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```terraform

data "ibm_is_bare_metal_server_disks" "ds_bmserver_disks" {
  bare_metal_server = "xxxx-xxxxx-xxxxx-xxxx"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `bare_metal_server` - (Required, String) The id for this bare metal server.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `disks` - (List of objects) A list of bare metal server disks. Disk is a block device that is locally attached to the physical server. By default, the listed disks are sorted by their created_at property values, with the newest disk first.

  Nested scheme for `disks`:
  - `href` - (String) The URL for this bare metal server disk.
  - `id` - (String) The unique identifier for this bare metal server disk.
  - `interface_type` - (String) The disk interface used for attaching the disk. Supported values are [ **nvme**, **sata** ].
  - `name` - (String) The user-defined name for this disk.
  - `resource_type` - (String) The resource type.
  - `size` - (String) The size of the disk in GB (gigabytes).