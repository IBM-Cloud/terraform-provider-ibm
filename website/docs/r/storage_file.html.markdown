---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: storage_file"
description: |-
  Manages IBM Cloud Storage file.
---

# ibm_storage_file
Create, delete, and update a file storage resource. This allows NFS-based Endurance and Performance [file storage](https://cloud.ibm.com/docs/infrastructure/FileStorage/index.html).

File storage is mounted by using the NFS protocol. For example, a file storage resource with the `hostname` argument set to `nfsdal0501a.service.softlayer.com` and the `volumename` argument set to ` IBM01SV278685_7` has the mount point `nfsdal0501a.service.softlayer.com:-IBM01SV278685_7`.

For more information, see [getting started with File Storage](https://cloud.ibm.com/docs/FileStorage/accessing-file-storage-linux.html) for NFS configuration.

## Example usage
In the following example, you can create 20G of Endurance file storage with a 10G snapshot capacity and 0.25 IOPS/GB.

```terraform
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
  snapshot_schedule {
    schedule_type   = "WEEKLY"
    retention_count = 20
    minute          = 2
    hour            = 13
    day_of_week     = "SUNDAY"
    enable          = true
  }
  snapshot_schedule {
    schedule_type   = "HOURLY"
    retention_count = 20
    minute          = 2
    enable          = true
  }

}

```

### In the following example, you can create 20G of Performance file storage with 100 IOPS.

```terraform
resource "ibm_storage_file" "fs_performance" {
  type       = "Performance"
  datacenter = "dal06"
  capacity   = 20
  iops       = 100

  # Optional fields
  allowed_virtual_guest_ids = ["28961689"]
  allowed_subnets           = ["10.146.139.64/26"]
  allowed_ip_addresses      = ["10.146.139.84"]
  hourly_billing            = true
}
```

## Timeouts

The `ibm_storage_file` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 45 minutes) Used for creating instance.
- **delete** - (Default 45 minutes) Used for deleting instance.
- **update** - (Default 45 minutes) Used for updating instance.

## Argument reference
Review the argument references that you can specify for your resource.

- `allowed_virtual_guest_ids` - (Optional, Array of Integers) The virtual guests that you want to give access to this instance. Virtual guests must be in the same data center as the block storage. You can also use this field to import the list of virtual guests that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_vm_instance` resource.
- `allowed_hardware_ids` - (Optional, Array of Integers) The Bare Metal servers that you want to give access to this instance. Bare Metal servers must be in the same data center as the block storage. You can also use this field to import the list of Bare Metal servers that have access to this storage from the `block_storage_ids` argument in the `ibm_compute_bare_metal` resource.
- `allowed_subnets` - (Optional, Array of Integers)The subnets that you want to give access to this instance. Subnets must be in the same data center as the block storage.
- `allowed_ip_addresses` - (Optional, Array of string) The IP addresses that you want to allow. IP addresses must be in the same data center as the block storage.
- `capacity` - (Required, Integer) The amount of storage capacity that you want to allocate, expressed in gigabytes.
- `datacenter` - (Required, Forces new resource, String) The data center where you want to provision the file storage instance.
- `hourly_billing` -  (Optional, Forces new resource,Bool) Set **true** to enable hourly billing. Default is **false**. **Note** `Hourly billing` is only available in updated data centers with improved capabilities. Refer to the [file storage locations](https://cloud.ibm.com/docs/FileStorage?topic=FileStorage-selectDC) to get the updated list of data centers.
- `iops`- (Required, Float) The IOPS value for the storage instance. For supported values, see [provisioning considerations](https://cloud.ibm.com/docs/FileStorage?topic=FileStorage-getting-started#provconsiderations).
- `notes` - (Optional, String)  Descriptive text to associate with the file storage.
- `snapshot_capacity` - (Optional, Forces new resource, Integer) The amount of snapshot capacity that you want to allocate, expressed in gigabytes.
- `snapshot_schedule` - (Optional, Array) Applies only to Endurance storage. Specifies the parameters required for a snapshot schedule.
- `snapshot_schedule.schedule_type` - (Optional, String) The snapshot schedule type. Accepted values are `HOURLY`, `WEEKLY`, and `DAILY`.
- `snapshot_schedule.retention_count` - (Optional, Integer) The retention count for a snapshot schedule. Required for all types of `schedule_type`.
- `snapshot_schedule.minute` - (Optional, Integer)The minute for a snapshot schedule. Required for all types of `schedule_type`.
- `snapshot_schedule.hour` - (Optional, Integer)The hour for a snapshot schedule. Required if `schedule_type` is set to `DAILY` or `WEEKLY`.
- `snapshot_schedule.day_of_week` - (Optional, String) The day of the week for a snapshot schedule. Required if the `schedule_type` is set to `WEEKLY`.
- `snapshot_schedule.enable` -  (Optional, Bool) Whether to disable an existing snapshot schedule.
- `tags` - (Optional, Arrays of Strings) Tags associated with the file storage instance.  **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `type` - (Required, Forces new resource, String) The type of the storage. Accepted values are `Endurance` and `Performance`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id`- (String) The unique identifier of the storage volume.
- `hostname`- (String) The fully qualified domain name of the storage.
- `mountpoint`- (String) The network mount address of the storage.
- `volumename`- (String) The name of the storage volume.
