---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_disk_management"
description: |-
  Manages IBM Dedicated host disk management.
---

# ibm_is_dedicated_host_disk_management

Create, update, delete and suspend the dedicated host disk management resource. For more information, about dedicated host disk management, see [migrating a dedicated host instance to another host](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-migrating-dedicated-host).

## Example usage

```terraform
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
  host_group = "1e09281b-f177-46fb-baf1-bc152b2e391a"
  name = "testdh02"
}
data "ibm_is_dedicated_host" "dhost" {
	host_group = "1e09281b-f177-46fb-baf1-bc152b2e391a"
	name = "my-dedicated-host"
}
data "ibm_is_dedicated_host_disks" "test1" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
}
data "ibm_is_dedicated_host_disk" "test1" {
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disk = ibm_is_dedicated_host_disk_management.disks.disks.0.id
}
resource "ibm_is_dedicated_host_disk_management" "disks"{
  dedicated_host = data.ibm_is_dedicated_host.dhost.id
  disks {
    name = "mydisk01"
    id = data.ibm_is_dedicated_host.dhost.disks.0.id
  }
}
```

## Argument reference
Review the argument reference that you can specify for your resource.

- `dedicated_host` - (Required, Force New Resource, String) The unique-identifier of the dedicated host.
- `disks` - (Required, List) Disks that needs to be updated. Nested `disks` blocks have the following structure:
  
  Nested scheme for `disks`:
  - `id` - (Required, String) The unique-identifier of the dedicated host disk.
  - `name` - (Required, String) The unique user defined name for the dedicated host disk.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique-identifier of the dedicated host disk management.

## Import

The `ibm_is_dedicated_host_disk_management` resource can be imported by using dedicated host disk management ID.

**Example**

```
$ terraform import ibm_is_dedicated_host_disk_management.example 0716-1c372bb2-decc-4555-b1a6-5d128c612316c
```
