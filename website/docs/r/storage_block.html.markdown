---
layout: "ibm"
page_title: "IBM: storage_block"
sidebar_current: "docs-ibm-resource-storage-block"
description: |-
  Manages IBM Storage Block.
---
# ibm\_storage_block

Provides a block storage resource. This allows iSCSI-based [Endurance](https://knowledgelayer.softlayer.com/topic/endurance-storage) and [Performance](https://knowledgelayer.softlayer.com/topic/performance-storage) block storage to be created, updated, and deleted.

Block storage can be accessed and mounted through a Multipath I/O (MPIO) Internet Small Computer System Interface (iSCSI) connection.

To access block storage, see the KnowledgeLayer docs [for Linux](https://knowledgelayer.softlayer.com/procedure/block-storage-linux) or [for Windows](https://knowledgelayer.softlayer.com/procedure/accessing-block-storage-microsoft-windows).

## Example Usage

In the following example, you can create 20G of Endurance block storage with 10G snapshot capacity and 0.25 IOPS/GB.

```hcl
resource "ibm_storage_block" "test1" {
        type = "Endurance"
        datacenter = "dal05"
        capacity = 20
        iops = 0.25
        os_format_type = "Linux"

        # Optional fields
        allowed_virtual_guest_ids = [ 27699397 ]
        allowed_ip_addresses = ["10.40.98.193", "10.40.98.200"]
        snapshot_capacity = 10
        hourly_billing = true
}
```

In the following example, you can create 20G of Performance block storage and 100 IOPS.

```hcl
resource "ibm_storage_block" "test2" {
        type = "Performance"
        datacenter = "dal05"
        capacity = 20
        iops = 100
        os_format_type = "Linux"

        # Optional fields
        allowed_virtual_guest_ids = [ 27699397 ]
        allowed_ip_addresses = ["10.40.98.193", "10.40.98.200"]
        hourly_billing = true
}
```

## Argument Reference

The following arguments are supported:

* `type` - (Required, string) The type of the storage. Accepted values are `Endurance` and `Performance`.
* `datacenter` - (Required, string) The data center where you want to provision the block storage instance.
* `capacity` - (Required, integer) The amount of storage capacity you want to allocate, specified in gigabytes.
* `iops` - (Required, float) The IOPS value for the storage. You can find available values for Endurance storage in the [IBM Cloud Infrastructure (SoftLayer) docs](https://knowledgelayer.softlayer.com/learning/introduction-endurance-storage).
* `os_format_type` - (Required, string) The OS type used to format the storage space. This OS type must match the OS type that connects to the LUN.
* `snapshot_capacity` - (Optional, integer) Applies to Endurance storage only. The amount of snapshot capacity to allocate, specified in gigabytes.
* `allowed_virtual_guest_ids` - (Optional, array of integers) The virtual guests that you want to give access to this instance. Virtual guests must be in the same data center as the block storage. You can also use this field to import the list of virtual guests that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_vm_instance` resource.
* `allowed_hardware_ids` - (Optional, array of integers) The bare metal servers that you want to give access to this instance. Bare metal servers must be in the same data center as the block storage. You can also use this field to import the list of bare metal servers that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_bare_metal` resource.
* `allowed_ip_addresses` - (Optional, array of string) The IP addresses that you want to give access to this instance. IP addresses must be in the same data center as the block storage.
* `notes` - (Optional, string) A descriptive note that you want to associate with the block storage.
* `tags` - (Optional, array of strings) Tags associated with the storage block instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.
* `hourly_billing` - (Optional,Boolean) Set true to enable hourly billing.Default is false
**NOTE**: `Hourly billing` is only available in updated datacenters with improved capabilities.Plesae refer the link to get the updated list of datacenter. http://knowledgelayer.softlayer.com/articles/new-ibm-block-and-file-storage-location-and-features



## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the storage.
* `hostname` - The fully qualified domain name of the storage.
* `volumename` - The name of the storage volume.
* `allowed_virtual_guest_info` - Deprecated please use `allowed_host_info` instead.
* `allowed_hardware_info` - Deprecated please use `allowed_host_info` instead.
* `allowed_host_info` - The user name, password, and host IQN of the hosts with access to the storage.
