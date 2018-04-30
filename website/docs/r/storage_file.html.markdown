---
layout: "ibm"
page_title: "IBM: storage_file"
sidebar_current: "docs-ibm-resource-storage-file"
description: |-
  Manages IBM Storage File.
---

# ibm\_storage_file

Provides a file storage resource. This allows NFS-based [Endurance](https://knowledgelayer.softlayer.com/topic/endurance-storage) and [Performance](https://knowledgelayer.softlayer.com/topic/performance-storage) file storage and [NAS\/FTP](https://console.bluemix.net/docs/infrastructure/network-attached-storage/index.html#getting-started-with-nas) file storage to be created, updated, and deleted.

File storage is mounted using the NFS protocol. For example, a file storage resource with the `hostname` argument set to `nfsdal0501a.service.softlayer.com` and the `volumename` argument set to ` IBM01SV278685_7` has the mount point `nfsdal0501a.service.softlayer.com:\IBM01SV278685_7`.

See [accessing file store on Linux](https://knowledgelayer.softlayer.com/procedure/accessing-file-storage-linux) for NFS configuration of Linux systems. For additional details, please refer to the [file storage docs](https://knowledgelayer.softlayer.com/topic/file-storage) and the [file storage overview](http://www.softlayer.com/file-storage).

## Example Usage

In the following example, you can create 20G of Endurance file storage with a 10G snapshot capacity and 0.25 IOPS/GB.

```hcl
resource "ibm_storage_file" "fs_endurance" {
  type       = "Endurance"
  datacenter = "dal06"
  capacity   = 20
  iops       = 0.25

  # Optional fields
  allowed_virtual_guest_ids = ["28961689"]
  allowed_subnets           = ["10.146.139.64/26"]
  allowed_ip_addresses      = ["10.146.139.84"]
  snapshot_capacity         = 10
  hourly_billing            = true

  # Optional fields for snapshot
  snapshot = [
    {
      schedule_type   = "WEEKLY"
      retention_count = 20
      minute          = 2
      hour            = 13
      day_of_week     = "SUNDAY"
      enable          = true
    },
    {
      schedule_type   = "HOURLY"
      retention_count = 20
      minute          = 2
      enable          = true
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
        hourly_billing = true
}
```

In the following example, you can create 20G of NAS/FTP file storage.

```hcl
resource "ibm_storage_file" "nas/ftp" {
        type = "NAS/FTP"
        datacenter = "dal06"
        capacity = 20
}
```


## Argument Reference

The following arguments are supported:

* `type` - (Required, string) The type of the storage. Accepted values are `Endurance`, `Performance` and `NAS/FTP`
* `datacenter` - (Required, string) The data center where you want to provision the file storage instance.
* `capacity` - (Required, integer) The amount of storage capacity you want to allocate, expressed in gigabytes.
* `iops` - (Required, float) The IOPS value for the storage instance. You can find available values for Endurance storage in the [KnowledgeLayer docs](https://knowledgelayer.softlayer.com/learning/introduction-endurance-storage). This attribute shouldn't be set when `type` is `NAS/FTP`.
* `snapshot_capacity` - (Optional, integer) Applies to Endurance storage only. The amount of snapshot capacity you want to allocate, expressed in gigabytes. This attribute shouldn't be set when `type` is `NAS/FTP`.
* `allowed_virtual_guest_ids` - (Optional, array of integers) The virtual guests that you want to give access to this instance. Virtual guests must be in the same data center as the block storage. You can also use this field to import the list of virtual guests that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_vm_instance` resource. This attribute shouldn't be set when `type` is `NAS/FTP`.
* `allowed_hardware_ids` - (Optional, array of integers) The bare metal servers that you want to give access to this instance. Bare metal servers must be in the same data center as the block storage. You can also use this field to import the list of bare metal servers that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_bare_metal` resource. This attribute shouldn't be set when `type` is `NAS/FTP`.
* `allowed_subnets` - (Optional, array of integers) The subnets that you want to give access to this instance. Subnets must be in the same data center as the block storage. This attribute shouldn't be set when `type` is `NAS/FTP`.
* `allowed_ip_addresses` - (Optional, array of string) The IP addresses that you want to allow. IP addresses must be in the same data center as the block storage. This attribute shouldn't be set when `type` is `NAS/FTP`.
* `snapshot_schedule` - (Optional, array) Applies only to Endurance storage. Specifies the parameters required for a snapshot schedule. This attribute shouldn't be set when `type` is `NAS/FTP`.
    * `schedule_type` - (String) The snapshot schedule type. Accepted values are `HOURLY`, `WEEKLY`, and `DAILY`.
    * `retention_count` - (Integer) The retention count for a snapshot schedule. Required for all types of `schedule_type`.
    * `minute` - (Integer) The minute for a snapshot schedule. Required for all types of `schedule_type`.
    * `hour` - (Integer) The hour for a snapshot schedule. Required if `schedule_type` is set to `DAILY` or `WEEKLY`.
    * `day_of_week` - (String) The day of the week for a snapshot schedule. Required if the `schedule_type` is set to `WEEKLY`.
    * `enable` - (Boolean) Whether to disable an existing snapshot schedule.

* `notes` - (Optional, string) Descriptive text to associate with the file storage.  

* `tags` - (Optional, array of strings) Tags associated with the file storage instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.  
* `hourly_billing` - (Optional,Boolean) Set true to enable hourly billing. Default is false. When `type` is `NAS/FTP` its value must be false.
**NOTE**: `Hourly billing` is only available in updated datacenters with improved capabilities.Plesae refer the link to get the updated list of datacenter. http://knowledgelayer.softlayer.com/articles/new-ibm-block-and-file-storage-location-and-features


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the storage volume.
* `hostname` - The fully qualified domain name of the storage.
* `volumename` - The name of the storage volume.
* `mountpoint` - The network mount address of the storage.