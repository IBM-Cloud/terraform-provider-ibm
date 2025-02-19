---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_volumes"
description: |-
  Manages Instance volumes in the Power Virtual Server cloud.
---

# ibm_pi_instance_volumes

Retrieves information about the persistent storage volumes that are mounted to a Power Systems Virtual Server instance. For more information, about power instance volume, see [snapshotting, cloning, and restoring](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-volume-snapshot-clone).

## Example Usage

The following example retrieves information about the volumes attached to the `terraform-test-instance` instance.

```terraform
data "ibm_pi_instance_volumes" "ds_volumes" {
  pi_instance_name     = "terraform-test-instance"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`
  
Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_name` - (Required, String) The unique identifier or name of the instance.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `boot_volume_id` - (String) The unique identifier of the boot volume.
- `instance_volumes` - (List) List of volumes attached to instance.

  Nested scheme for `instance_volumes`:
  - `bootable`- (Boolean) Indicates if the volume is boot capable.
  - `creation_date` - (String) Date of volume creation.
  - `crn` - (String) The CRN of this resource.
  - `freeze_time` - (String) Time of remote copy relationship.
  - `href` - (String) The hyper link of the volume.
  - `id` - (String) The unique identifier of the volume.
  - `last_update_date` - (String) The date when the volume last updated.
  - `name` - (String) The name of the volume.
  - `pool` - (String) Volume pool, name of storage pool where the volume is located.
  - `replication_enabled` - (Boolean) Indicates whether replication is enabled on the volume.
  - `replication_sites` - (List) List of replication sites for volume replication.
  - `shareable` - (Boolean) Indicates if the volume is shareable between VMs.
  - `size` - (Integer) The size of this volume in GB.
  - `state` - (String) The state of the volume.
  - `type` - (String) The disk type that is used for this volume.
  - `user_tags` - (List) List of user tags attached to the resource.
