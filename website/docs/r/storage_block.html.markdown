---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: storage_block"
description: |-
  Manages IBM Storage block.
---

# ibm_storage_block
Create, delete, or update a block storage resource. For more information, about Block storage, see [getting startecwith block storage](https://cloud.ibm.com/docs/BlockStorage?topic=BlockStorage-getting-started). 

Block storage can be accessed and mounted through a Multipath Input/Output Internet Small Computer System Interface (iSCSI) connection.

## Example usage
In the following example, you can create 20G of Endurance block storage with 10G snapshot capacity and 0.25 IOPS/GB.

```terraform
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

```terraform
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

## Timeouts

The `ibm_storage_block` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 45 minutes) Used for creating instance.
- **delete** - (Default 45 minutes) Used for deleting instance.
- **update** - (Default 45 minutes) Used for updating instance.


## Argument reference 
Review the argument references that you can specify for your resource.

- `allowed_virtual_guest_ids`- (Optional, Array of Integers) The virtual guests that you want to give access to this instance. Virtual guests must be in the same data center as the block storage. You can also use this field to import the list of virtual guests that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_vm_instance` resource.
- `allowed_hardware_ids`- (Optional, Array of Integers) The Bare Metal servers that you want to give access to this instance. Bare Metal servers must be in the same data center as the block storage. You can also use this field to import the list of Bare Metal servers that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_bare_metal` resource.
- `allowed_ip_addresses`- (Optional, Array of string) The IP addresses that you want to give access to this instance. IP addresses must be in the same data center as the block storage.
- `capacity` - (Required, Integer) The amount of storage capacity that you want to allocate, specified in gigabytes.
- `datacenter`- (Required, Forces new resource, String) The data center where you want to provision the block storage instance.
- `hourly_billing` -  (Optional, Bool) Set true to enable hourly billing. Default value is **false**   **Note** `Hourly billing` is only available in updated data centers with improved capabilities. Refer to the link to get the updated list of data centers. See [file storage locations](https://cloud.ibm.com/docs/FileStorage?topic=FileStorage-selectDC).
- `iops`- (Required, Float) The IOPS value for the storage. For supported values for endurance storage, see [IBM Cloud Classic Infrastructure (SoftLayer)](https://cloud.ibm.com/docs/FileStorage?topic=FileStorage-orderingFileStorage).
- `os_format_type` - (Required, Forces new resource, String) The OS type used to format the storage space. This OS type must match the OS type that connects to the LUN. [Log in to the IBM Cloud Classic Infrastructure API to see available OS format types](https://api.softlayer.com/rest/v3/SoftLayer_Network_Storage_Iscsi_OS_Type/getAllObjects/). Use your API as the password to log in. Log in and find the key called `name`.
- `notes` -  (Optional, String) A descriptive note that you want to associate with the block storage.
- `snapshot_capacity` - (Optional, Forces new resource, Integer) The amount of snapshot capacity to allocate, specified in gigabytes.
- `type` - (Required, Forces new resource, String)The type of the storage. Accepted values are **Endurance** and **Performance**.
- `tags` - (Optional, Array of string) Tags associated with the storage block instance.     **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `allowed_virtual_guest_info` - (String) Deprecated please use `allowed_host_info` instead.
- `allowed_hardware_info` - (String) Deprecated please use `allowed_host_info` instead.
- `allowed_host_info` - (String) The user name, password, and host IQN of the hosts with access to the storage.
- `hostname` - (String) The fully qualified domain name of the storage.
- `id`- (String) The unique identifier of the storage.
- `lunid` -  (String) The `LUN` ID of the storage device.
- `volumename` - (String) The name of the storage volume.
