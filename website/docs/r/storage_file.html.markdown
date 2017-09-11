---
layout: "ibm"
page_title: "IBM: storage_file"
sidebar_current: "docs-ibm-resource-storage-file"
description: |-
  Manages IBM Storage File.
---

# ibm\_storage_file

Provides a file storage resource. This allows NFS-based [Endurance](https://knowledgelayer.softlayer.com/topic/endurance-storage) and [Performance](https://knowledgelayer.softlayer.com/topic/performance-storage) to be created, updated, and deleted.

File storage is mounted using the NFS protocol. For example, if the `hostname` of the file storage resource is `nfsdal0501a.service.softlayer.com` and the `volumename` is` IBM01SV278685_7`, the mount point would be `nfsdal0501a.service.softlayer.com:\IBM01SV278685_7`. 

See [accessing file store on Linux](https://knowledgelayer.softlayer.com/procedure/accessing-file-storage-linux) for NFS configuration of Linux systems. For additional details, please refer to the [file storage docs](https://knowledgelayer.softlayer.com/topic/file-storage) and the [file storage overview](http://www.softlayer.com/file-storage).

## Example Usage

In the following example, you can create 20G of Endurance file storage with a 10G snapshot capacity and 0.25 IOPS/GB.

```hcl
resource "ibm_storage_file" "fs_endurance" {
        type = "Endurance"
        datacenter = "dal06"
        capacity = 20
        iops = 0.25
        
        # Optional fields
        allowed_virtual_guest_ids = [ "28961689" ]
        allowed_subnets = [ "10.146.139.64/26" ]
        allowed_ip_addresses = [ "10.146.139.84" ]
        snapshot_capacity = 10  
        
        #Optional fields for snapshot
        
        snapshot = [
  		{
			scheduleType="WEEKLY",
			retentionCount= 20,
			minute= 2,
			hour= 13,
			dayOfWeek= "SUNDAY",
			enable= true
		},
		{
			scheduleType="HOURLY",
			retentionCount= 20,
			minute= 2,
			enable= true
		},
 		]		
}

```

In the following example, you can create 20G of Performance file storage with 100 IOPS.

```hcl
resource "ibm_storage_file" "fs_performance" {
        type = "Performance"
        datacenter = "dal06"
        capacity = 20
        iops = 100
        # Optional fields
        allowed_virtual_guest_ids = [ "28961689" ]
        allowed_subnets = [ "10.146.139.64/26" ]
        allowed_ip_addresses = [ "10.146.139.84" ]
}
```

## Argument Reference

The following arguments are supported:

* `type` - (Required, string) The type of the storage. Accepted values are `Endurance` and `Performance`.
* `datacenter` - (Required, string) The data center the file storage instance is to be provisioned in.
* `capacity` - (Required, integer) The amount of storage capacity to allocate, expressed in gigabytes.
* `iops` - (Required, float) The IOPS value for the storage instance. Available values for Endurance storage can be found in the [KnowledgeLayer docs](https://knowledgelayer.softlayer.com/learning/introduction-endurance-storage).
* `snapshot_capacity` - (Optional, integer) The amount of snapshot capacity to allocate, expressed in gigabytes. Only applies to `Endurance` storage.
* `allowed_virtual_guest_ids` - (Optional, array of integers) Specify allowed virtual guests. Virtual guests need to be in the same data center. You can also use this field to list the virtual guests which were provided access to this storage through the `file_storage_ids` argument in the `ibm_compute_vm_instance` resource. 
* `allowed_hardware_ids`- (Optional, array of integers) Specify allowed bare metal servers. Bare metal servers need to be in the same data center. You can use also this field to list the bare metals which were provided access to this storage through the `file_storage_ids` argument in the `ibm_compute_bare_metal` resource . 
* `allowed_subnets` - (Optional, array of integers) Specify allowed subnets. Subnets should be in the same data center.
* `allowed_ip_addresses` - (Optional, array of integers) Specify allowed IP addresses. IP addresses need to be in the same data center.
* `snapshot` - (Optional) Specifies the parameter required for a snapshot schedule. Only applies to Endurance storage.
* `scheduleType` - (String) Specifies the snapshot schedule type. Accepted values are `HOURLY`, `WEEKLY`, and `DAILY`.
* `retentionCount` - (Integer) Specifies the retention count. Required for all types of `scheduleType`.
* `minute` - (Integer) Specifies the minute. Required for all types of `scheduleType`.
* `hour` - (Integer) Specifies the hour. Required if `scheduleType` is set to `DAILY` and `WEEKLY`.
* `dayOfWeek` - (string) Specifies the day of the week. Required if the `scheduleType` is set to `WEEKLY`.
* `enable` - (boolean) Specifies whether to disable the already created snapshot.
* `notes` - (Optional,string) Specifies a note to associate with the file storage.
* `tags` - (Optional, array of strings) Set tags on the file storage instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - Identifier of the storage volume.
* `hostname` - The fully qualified domain name of the storage. 
* `volumename` - The name of the storage volume.
* `mountpoint` - The network mount address of the storage.
