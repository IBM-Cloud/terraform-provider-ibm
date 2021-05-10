---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_disk_management"
description: |-
  Manages IBM DedicatedHost Disk Management.
---

# ibm\_is_dedicated_host_disk_management

Provides a resource for DedicatedHost Disk Management. This allows DedicatedHost disk names to be updated

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

## Argument Reference

The following arguments are supported:


* `dedicated_host` - (Required, string, ForceNew) The unique-identifier of the dedicated host
* `disks` - (Required, string) Disks that needs to be updated. Nested `disks` blocks have the following structure:
	* `id` - (Required, string) The unique-identifier of the dedicated host disk.
	* `name` - (Required, string) The unique user defined name for the dedicated host disk

## Attribute Reference

The following attributes are exported:


* `id` - The unique-identifier of the Dedicated host disk management

## Import

ibm_is_dedicated_host_disk_management can be imported using Dedicated Host disk management ID, eg

```
$ terraform import ibm_is_dedicated_host_disk_management.example 0716-1c372bb2-decc-4555-b1a6-5d128c62806c
```