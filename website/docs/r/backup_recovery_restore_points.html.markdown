---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_restore_points"
description: |-
  Manages backup_recovery_restore_points.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_restore_points

Create backup_recovery_restore_pointss with this resource.

**Note**
ibm_backup_recovery_restore_points resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery_restore_points" "backup_recovery_restore_points_instance" {
  end_time_usecs = 12
  environment = "kVMware"
  start_time_usecs = 14
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `end_time_usecs` - (Required, Forces new resource, Integer) Specifies the end time specified as a Unix epoch Timestamp in microseconds.
* `environment` - (Required, Forces new resource, String) Specifies the protection source environment type.
  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kAzure`, `kGCP`, `kKVM`, `kAcropolis`, `kAWS`, `kPhysical`, `kGPFS`, `kElastifile`, `kNetapp`, `kGenericNas`, `kIsilon`, `kFlashBlade`, `kPure`, `kIbmFlashSystem`, `kSQL`, `kExchange`, `kAD`, `kOracle`, `kView`, `kRemoteAdapter`, `kO365`, `kKubernetes`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHive`, `kSAPHANA`, `kHBase`, `kUDA`, `kSfdc`.
* `protection_group_ids` - (Required, Forces new resource, List) Specifies the jobs for which to get the full snapshot information.
* `source_id` - (Optional, Forces new resource, Integer) Specifies the id of the Protection Source which is to be restored.
* `start_time_usecs` - (Required, Forces new resource, Integer) Specifies the start time specified as a Unix epoch Timestamp in microseconds.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_restore_points.


## Import
Not Supported
