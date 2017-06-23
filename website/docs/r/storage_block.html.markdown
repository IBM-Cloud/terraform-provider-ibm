---
layout: "ibm"
page_title: "IBM: storage_block"
sidebar_current: "docs-ibm-resource-storage-block"
description: |-
  Manages IBM Storage Block.
---
# ibm\_storage_block

Provides a resource to create, update, and delete iSCSI-based [Endurance](https://knowledgelayer.softlayer.com/topic/endurance-storage) and [Performance](https://knowledgelayer.softlayer.com/topic/performance-storage) block storage.

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
}
```

## Argument Reference

The following arguments are supported:

* `type` - (Required, string) The type of the storage. Accepted values are `Endurance` and `Performance`.
* `datacenter` - (Required, string) The data center the instance is to be provisioned in.
* `capacity` - (Required, integer) The amount of storage capacity to allocate, specified in gigabytes.
* `iops` - (Required, float) The IOPS value for the storage. You can find available values for Endurance storage in the [Bluemix Infrastructure (SoftLayer) docs](https://knowledgelayer.softlayer.com/learning/introduction-endurance-storage).
* `os_format_type` - (Required, string) Specifies which OS type to use when formatting the storage space. This should match the OS type that will be connecting to the LUN.
* `snapshot_capacity` - (Optional, integer) The amount of snapshot capacity to allocate, specified in gigabytes. Only applies to Endurance storage.
* `allowed_virtual_guest_ids` - (Optional, array of integers) Specifies allowed virtual guests. Virtual guests need to be in the same data center. You can also use this field to list the virtual guests which were provided access to this storage through the `block_storage_ids` argument in the `ibm_compute_vm_instance` resource. 
* `allowed_hardware_ids` - (Optional, array of integers) Specifies allowed bare metal servers. Bare metal servers need to be in the same data center. You can also use this field to list the bare metals which were provided access to this storage through the `block_storage_ids` argument in the `ibm_compute_bare_metal` resource. 
* `allowed_ip_addresses` - (Optional, array of string) Specifies allowed IP addresses. IP addresses need to be in the same data center.


## Attributes Reference

The following attributes are exported:

* `id` - Identifier of the storage.
* `hostname` - The fully qualified domain name of the storage.
* `volumename` - The name of the storage volume.
* `allowed_virtual_guest_info` - Contains user name, password and host IQN of the virtual guests with access to the storage.
* `allowed_hardware_info` - Contains user name, password and host IQN of the bare metal servers with access to the storage.
